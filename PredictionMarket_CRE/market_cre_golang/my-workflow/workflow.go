package main

import (
	"cre-custom-data-feed-go/contracts/evm/src/generated/confidential_market"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	"log/slog"
	"math/big"
	"net/url"
	"strings"
)

// Config 工作流配置
type Config struct {
	ContractAddress string    `json:"contractAddress"` // 合约地址
	ChainSelector   uint64    `json:"chainSelector"`   // 链选择器，如 16015286601757825753 (Sepolia)
	APIConfig       APIConfig `json:"apiConfig"`       // 外部 API 配置
	GasLimit        uint64    `json:"gasLimit"`        // 交易 gas 限制
}

type APIConfig struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
}

// InitWorkflow 创建并返回工作流
func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	// 创建 EVM 客户端
	evmClient := &evm.Client{
		ChainSelector: config.ChainSelector,
	}

	// 创建合约实例
	contractAddress := common.HexToAddress(config.ContractAddress)
	contract, err := confidential_market.NewConfidentialMarket(evmClient, contractAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract binding: %w", err)
	}

	// 创建事件触发器，监听 SettlementRequested 事件
	// 我们不过滤特定市场，监听所有
	trigger, err := contract.LogTriggerSettlementRequestedLog(
		config.ChainSelector,
		evm.ConfidenceLevel_CONFIDENCE_LEVEL_LATEST,
		nil, // 无特定 topics 过滤
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create log trigger: %w", err)
	}

	// 使用闭包将 contract 传入处理器
	handler := cre.Handler(trigger, func(config *Config, runtime cre.Runtime, payload *bindings.DecodedLog[confidential_market.SettlementRequestedDecoded]) (string, error) {
		return onSettlementRequested(config, runtime, payload, contract)
	})

	return cre.Workflow[*Config]{handler}, nil
}

// 用于解析 Geocoding API 响应的临时结构
type geoResult struct {
	Lat float64 `json:"latitude" consensus_aggregation:"median"`
	Lon float64 `json:"longitude" consensus_aggregation:"median"`
}

type outcomeWrapper struct {
	Value uint8 `consensus_aggregation:"median"`
}

// onSettlementRequested 是事件触发后的处理函数
func onSettlementRequested(
	config *Config,
	runtime cre.Runtime,
	payload *bindings.DecodedLog[confidential_market.SettlementRequestedDecoded],
	contract *confidential_market.ConfidentialMarket, // 传入合约绑定
) (string, error) {
	logger := runtime.Logger()

	// 1. 从事件中提取 marketId (*big.Int)
	marketId := payload.Data.MarketId
	logger.Info("Received SettlementRequested", "marketId", marketId)

	// 2. 从链上读取市场描述
	args := confidential_market.MarketsInput{
		Arg0: marketId,
	}

	// 官方推荐：使用 finalized 区块 (-3)
	blockNumber := big.NewInt(-3)
	marketsPromise := contract.Markets(runtime, args, blockNumber)

	marketsOutput, err := marketsPromise.Await()
	if err != nil {
		// 增加更详细的错误日志
		logger.Error("Failed to get market data from chain", "error", err, "marketId", marketId)
		return "", fmt.Errorf("failed to get market data from chain (market %s): %w", marketId.String(), err)
	}

	description := marketsOutput.Description

	logger.Info("Market description from chain", "description", description)

	//description := "2025-03-08:London" // 硬编码，便于模拟
	logger.Info("Market description", "description", description)

	// 3. 从 description 中提取城市名（假设格式为 "date:city"）
	parts := strings.Split(description, ":")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid description format: %s", description)
	}
	cityName := parts[1] // 取冒号后的城市名
	logger.Info("Using city name", "city", cityName)

	// 4. 调用 Geocoding API 获取经纬度
	geocodingURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1", url.QueryEscape(cityName))
	logger.Info("Calling Geocoding API", "url", geocodingURL)

	client := &http.Client{}
	geoResp, err := http.SendRequest(config, runtime, client, func(config *Config, logger *slog.Logger, sendRequester *http.SendRequester) (geoResult, error) {
		req := &http.Request{
			Method: "GET",
			Url:    geocodingURL,
		}
		resp, err := sendRequester.SendRequest(req).Await()
		if err != nil {
			return geoResult{}, err
		}
		if len(resp.Body) == 0 {
			return geoResult{}, fmt.Errorf("empty response body")
		}

		var geoData struct {
			Results []struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"results"`
		}
		if err := json.Unmarshal(resp.Body, &geoData); err != nil {
			return geoResult{}, fmt.Errorf("failed to parse geocoding response: %w", err)
		}
		if len(geoData.Results) == 0 {
			return geoResult{}, fmt.Errorf("city not found: %s", cityName)
		}
		return geoResult{
			Lat: geoData.Results[0].Latitude,
			Lon: geoData.Results[0].Longitude,
		}, nil
	}, cre.ConsensusAggregationFromTags[geoResult]()).Await()
	if err != nil {
		return "", fmt.Errorf("geocoding failed: %w", err)
	}
	logger.Info("Geocoding result", "lat", geoResp.Lat, "lon", geoResp.Lon)

	// 5. 调用 Weather API 获取当前天气（使用配置中的 URL）
	weatherURL := fmt.Sprintf("%s?latitude=%f&longitude=%f&current_weather=true",
		config.APIConfig.URL, geoResp.Lat, geoResp.Lon)
	logger.Info("Calling Weather API", "url", weatherURL)

	outcomeWrapperVal, err := http.SendRequest(config, runtime, client, func(config *Config, logger *slog.Logger, sendRequester *http.SendRequester) (outcomeWrapper, error) {
		req := &http.Request{
			Method:  config.APIConfig.Method,
			Url:     weatherURL,
			Headers: config.APIConfig.Headers,
		}
		resp, err := sendRequester.SendRequest(req).Await()
		if err != nil {
			return outcomeWrapper{}, err
		}
		if len(resp.Body) == 0 {
			return outcomeWrapper{}, fmt.Errorf("empty response body")
		}

		var weatherData struct {
			CurrentWeather struct {
				Temperature float64 `json:"temperature"`
			} `json:"current_weather"`
		}
		if err := json.Unmarshal(resp.Body, &weatherData); err != nil {
			return outcomeWrapper{}, fmt.Errorf("failed to parse weather response: %w", err)
		}

		// 根据温度决定 outcome（示例：温度 > 5℃ 为 1，否则为 0）
		if weatherData.CurrentWeather.Temperature > 5 {
			return outcomeWrapper{Value: 1}, nil
		}
		return outcomeWrapper{Value: 0}, nil
	}, cre.ConsensusAggregationFromTags[outcomeWrapper]()).Await()
	if err != nil {
		return "", fmt.Errorf("weather API failed: %w", err)
	}
	outcome := outcomeWrapperVal.Value
	logger.Info("Weather result", "outcome", outcome)

	// 6. 编码 submitResult 的 calldata（合约需要 signature 参数，但逻辑已忽略，传空切片）
	codec, err := confidential_market.NewCodec()
	if err != nil {
		return "", fmt.Errorf("failed to create codec: %w", err)
	}
	callData, err := codec.EncodeSubmitResultMethodCall(confidential_market.SubmitResultInput{
		MarketId:  marketId,
		Outcome:   outcome,
		Signature: []byte{}, // 合约不验证签名，传空即可
	})
	if err != nil {
		return "", fmt.Errorf("failed to encode calldata: %w", err)
	}

	// 7. 通过 EVM 客户端直接发送交易
	// 从 runtime 获取 EVM 客户端
	reportReq := &cre.ReportRequest{
		EncodedPayload: callData,
		EncoderName:    "evm",
		SigningAlgo:    "evm",
		HashingAlgo:    "keccak256",
	}

	reportPromise := runtime.GenerateReport(reportReq)
	report, err := reportPromise.Await()
	if err != nil {
		logger.Error("GenerateReport failed", "error", err)
		return "", fmt.Errorf("GenerateReport failed: %w", err)
	}

	// 设置 gas 配置
	gasConfig := &evm.GasConfig{
		GasLimit: config.GasLimit,
	}

	// 发送交易
	txResponse, err := contract.WriteReport(runtime, report, gasConfig).Await()
	if err != nil {
		logger.Error("WriteReport failed", "error", err)
		return "", fmt.Errorf("WriteReport failed: %w", err)
	}

	txHash := common.BytesToHash(txResponse.TxHash).Hex()
	logger.Info("submitResult succeeded", "txHash", txHash)

	return fmt.Sprintf("outcome: %d", outcome), nil
}
