// Code generated — DO NOT EDIT.

package confidential_market

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb2 "github.com/smartcontractkit/chainlink-protos/cre/go/sdk"
	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Sprintf
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = emptypb.Empty{}
	_ = pb.NewBigIntFromInt
	_ = pb2.AggregationType_AGGREGATION_TYPE_COMMON_PREFIX
	_ = bindings.FilterOptions{}
	_ = evm.FilterLogTriggerRequest{}
	_ = cre.ResponseBufferTooSmall
	_ = rpc.API{}
	_ = json.Unmarshal
	_ = reflect.Bool
)

var ConfidentialMarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encryptedBet\",\"type\":\"bytes\"}],\"name\":\"BetPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"closeTime\",\"type\":\"uint48\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"outcome\",\"type\":\"uint8\"}],\"name\":\"MarketResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"SettlementRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WinningsClaimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OrderMarketid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betIndex\",\"type\":\"uint256\"}],\"name\":\"claimWinnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint48\",\"name\":\"closeTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"resolveWindow\",\"type\":\"uint48\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketTotalPools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"enumMarketData.MarketStatus\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint48\",\"name\":\"createTimeStamp\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"closeTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"resolveWindow\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"resolveTimeStamp\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"outcome\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"resolved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encryptedBet\",\"type\":\"bytes\"}],\"name\":\"placeBet\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"requestSettlement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"betIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"options\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"revealBets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setTrustedCRE\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"outcome\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"submitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedCREAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Structs

// Contract Method Inputs
type AdminsInput struct {
	Arg0 common.Address
}

type ClaimWinningsInput struct {
	MarketId *big.Int
	BetIndex *big.Int
}

type CreateMarketInput struct {
	Description   string
	CloseTime     *big.Int
	ResolveWindow *big.Int
}

type MarketTotalPoolsInput struct {
	Arg0 *big.Int
}

type MarketsInput struct {
	Arg0 *big.Int
}

type PlaceBetInput struct {
	MarketId     *big.Int
	EncryptedBet []byte
}

type RequestSettlementInput struct {
	MarketId *big.Int
}

type RevealBetsInput struct {
	MarketId   *big.Int
	BetIndices []*big.Int
	Options    []uint8
	Amounts    []*big.Int
}

type SetAdminInput struct {
	Admin  common.Address
	Status bool
}

type SetTrustedCREInput struct {
	Addr common.Address
}

type SubmitResultInput struct {
	MarketId  *big.Int
	Outcome   uint8
	Signature []byte
}

type TransferOwnershipInput struct {
	NewOwner common.Address
}

// Contract Method Outputs
type MarketsOutput struct {
	Id               *big.Int
	Description      string
	State            uint8
	CreateTimeStamp  *big.Int
	CloseTime        *big.Int
	ResolveWindow    *big.Int
	ResolveTimeStamp *big.Int
	Outcome          uint8
	Resolved         bool
}

// Errors
type OwnableInvalidOwner struct {
	Owner common.Address
}

type OwnableUnauthorizedAccount struct {
	Account common.Address
}

// Events
// The <Event>Topics struct should be used as a filter (for log triggers).
// Note: It is only possible to filter on indexed fields.
// Indexed (string and bytes) fields will be of type common.Hash.
// They need to he (crypto.Keccak256) hashed and passed in.
// Indexed (tuple/slice/array) fields can be passed in as is, the Encode<Event>Topics function will handle the hashing.
//
// The <Event>Decoded struct will be the result of calling decode (Adapt) on the log trigger result.
// Indexed dynamic type fields will be of type common.Hash.

type BetPlacedTopics struct {
	MarketId *big.Int
	User     common.Address
}

type BetPlacedDecoded struct {
	MarketId     *big.Int
	User         common.Address
	Amount       *big.Int
	EncryptedBet []byte
}

type MarketCreatedTopics struct {
	MarketId *big.Int
}

type MarketCreatedDecoded struct {
	MarketId    *big.Int
	Description string
	CloseTime   *big.Int
}

type MarketResolvedTopics struct {
	MarketId *big.Int
}

type MarketResolvedDecoded struct {
	MarketId *big.Int
	Outcome  uint8
}

type OwnershipTransferredTopics struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

type OwnershipTransferredDecoded struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

type SettlementRequestedTopics struct {
	MarketId *big.Int
}

type SettlementRequestedDecoded struct {
	MarketId *big.Int
}

type WinningsClaimedTopics struct {
	MarketId *big.Int
	User     common.Address
}

type WinningsClaimedDecoded struct {
	MarketId *big.Int
	User     common.Address
	Amount   *big.Int
}

// Main Binding Type for ConfidentialMarket
type ConfidentialMarket struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   ConfidentialMarketCodec
}

type ConfidentialMarketCodec interface {
	EncodeOrderMarketidMethodCall() ([]byte, error)
	DecodeOrderMarketidMethodOutput(data []byte) (*big.Int, error)
	EncodeAdminsMethodCall(in AdminsInput) ([]byte, error)
	DecodeAdminsMethodOutput(data []byte) (bool, error)
	EncodeClaimWinningsMethodCall(in ClaimWinningsInput) ([]byte, error)
	EncodeCreateMarketMethodCall(in CreateMarketInput) ([]byte, error)
	DecodeCreateMarketMethodOutput(data []byte) (*big.Int, error)
	EncodeMarketTotalPoolsMethodCall(in MarketTotalPoolsInput) ([]byte, error)
	DecodeMarketTotalPoolsMethodOutput(data []byte) (*big.Int, error)
	EncodeMarketsMethodCall(in MarketsInput) ([]byte, error)
	DecodeMarketsMethodOutput(data []byte) (MarketsOutput, error)
	EncodeOwnerMethodCall() ([]byte, error)
	DecodeOwnerMethodOutput(data []byte) (common.Address, error)
	EncodePlaceBetMethodCall(in PlaceBetInput) ([]byte, error)
	EncodeRenounceOwnershipMethodCall() ([]byte, error)
	EncodeRequestSettlementMethodCall(in RequestSettlementInput) ([]byte, error)
	EncodeRevealBetsMethodCall(in RevealBetsInput) ([]byte, error)
	EncodeSetAdminMethodCall(in SetAdminInput) ([]byte, error)
	EncodeSetTrustedCREMethodCall(in SetTrustedCREInput) ([]byte, error)
	EncodeSubmitResultMethodCall(in SubmitResultInput) ([]byte, error)
	EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error)
	EncodeTrustedCREAddressMethodCall() ([]byte, error)
	DecodeTrustedCREAddressMethodOutput(data []byte) (common.Address, error)
	BetPlacedLogHash() []byte
	EncodeBetPlacedTopics(evt abi.Event, values []BetPlacedTopics) ([]*evm.TopicValues, error)
	DecodeBetPlaced(log *evm.Log) (*BetPlacedDecoded, error)
	MarketCreatedLogHash() []byte
	EncodeMarketCreatedTopics(evt abi.Event, values []MarketCreatedTopics) ([]*evm.TopicValues, error)
	DecodeMarketCreated(log *evm.Log) (*MarketCreatedDecoded, error)
	MarketResolvedLogHash() []byte
	EncodeMarketResolvedTopics(evt abi.Event, values []MarketResolvedTopics) ([]*evm.TopicValues, error)
	DecodeMarketResolved(log *evm.Log) (*MarketResolvedDecoded, error)
	OwnershipTransferredLogHash() []byte
	EncodeOwnershipTransferredTopics(evt abi.Event, values []OwnershipTransferredTopics) ([]*evm.TopicValues, error)
	DecodeOwnershipTransferred(log *evm.Log) (*OwnershipTransferredDecoded, error)
	SettlementRequestedLogHash() []byte
	EncodeSettlementRequestedTopics(evt abi.Event, values []SettlementRequestedTopics) ([]*evm.TopicValues, error)
	DecodeSettlementRequested(log *evm.Log) (*SettlementRequestedDecoded, error)
	WinningsClaimedLogHash() []byte
	EncodeWinningsClaimedTopics(evt abi.Event, values []WinningsClaimedTopics) ([]*evm.TopicValues, error)
	DecodeWinningsClaimed(log *evm.Log) (*WinningsClaimedDecoded, error)
}

func NewConfidentialMarket(
	client *evm.Client,
	address common.Address,
	options *bindings.ContractInitOptions,
) (*ConfidentialMarket, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfidentialMarketMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &ConfidentialMarket{
		Address: address,
		Options: options,
		ABI:     &parsed,
		client:  client,
		Codec:   codec,
	}, nil
}

type Codec struct {
	abi *abi.ABI
}

func NewCodec() (ConfidentialMarketCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfidentialMarketMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
}

func (c *Codec) EncodeOrderMarketidMethodCall() ([]byte, error) {
	return c.abi.Pack("OrderMarketid")
}

func (c *Codec) DecodeOrderMarketidMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["OrderMarketid"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeAdminsMethodCall(in AdminsInput) ([]byte, error) {
	return c.abi.Pack("admins", in.Arg0)
}

func (c *Codec) DecodeAdminsMethodOutput(data []byte) (bool, error) {
	vals, err := c.abi.Methods["admins"].Outputs.Unpack(data)
	if err != nil {
		return *new(bool), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(bool), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result bool
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(bool), fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeClaimWinningsMethodCall(in ClaimWinningsInput) ([]byte, error) {
	return c.abi.Pack("claimWinnings", in.MarketId, in.BetIndex)
}

func (c *Codec) EncodeCreateMarketMethodCall(in CreateMarketInput) ([]byte, error) {
	return c.abi.Pack("createMarket", in.Description, in.CloseTime, in.ResolveWindow)
}

func (c *Codec) DecodeCreateMarketMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["createMarket"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeMarketTotalPoolsMethodCall(in MarketTotalPoolsInput) ([]byte, error) {
	return c.abi.Pack("marketTotalPools", in.Arg0)
}

func (c *Codec) DecodeMarketTotalPoolsMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["marketTotalPools"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeMarketsMethodCall(in MarketsInput) ([]byte, error) {
	return c.abi.Pack("markets", in.Arg0)
}

func (c *Codec) DecodeMarketsMethodOutput(data []byte) (MarketsOutput, error) {
	vals, err := c.abi.Methods["markets"].Outputs.Unpack(data)
	if err != nil {
		return MarketsOutput{}, err
	}
	if len(vals) != 9 {
		return MarketsOutput{}, fmt.Errorf("expected 9 values, got %d", len(vals))
	}
	jsonData0, err := json.Marshal(vals[0])
	if err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to marshal ABI result 0: %w", err)
	}

	var result0 *big.Int
	if err := json.Unmarshal(jsonData0, &result0); err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData1, err := json.Marshal(vals[1])
	if err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to marshal ABI result 1: %w", err)
	}

	var result1 string
	if err := json.Unmarshal(jsonData1, &result1); err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to unmarshal to string: %w", err)
	}
	jsonData2, err := json.Marshal(vals[2])
	if err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to marshal ABI result 2: %w", err)
	}

	var result2 uint8
	if err := json.Unmarshal(jsonData2, &result2); err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to unmarshal to uint8: %w", err)
	}
	jsonData3, err := json.Marshal(vals[3])
	if err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to marshal ABI result 3: %w", err)
	}

	var result3 *big.Int
	if err := json.Unmarshal(jsonData3, &result3); err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData4, err := json.Marshal(vals[4])
	if err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to marshal ABI result 4: %w", err)
	}

	var result4 *big.Int
	if err := json.Unmarshal(jsonData4, &result4); err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData5, err := json.Marshal(vals[5])
	if err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to marshal ABI result 5: %w", err)
	}

	var result5 *big.Int
	if err := json.Unmarshal(jsonData5, &result5); err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData6, err := json.Marshal(vals[6])
	if err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to marshal ABI result 6: %w", err)
	}

	var result6 *big.Int
	if err := json.Unmarshal(jsonData6, &result6); err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData7, err := json.Marshal(vals[7])
	if err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to marshal ABI result 7: %w", err)
	}

	var result7 uint8
	if err := json.Unmarshal(jsonData7, &result7); err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to unmarshal to uint8: %w", err)
	}
	jsonData8, err := json.Marshal(vals[8])
	if err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to marshal ABI result 8: %w", err)
	}

	var result8 bool
	if err := json.Unmarshal(jsonData8, &result8); err != nil {
		return MarketsOutput{}, fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return MarketsOutput{
		Id:               result0,
		Description:      result1,
		State:            result2,
		CreateTimeStamp:  result3,
		CloseTime:        result4,
		ResolveWindow:    result5,
		ResolveTimeStamp: result6,
		Outcome:          result7,
		Resolved:         result8,
	}, nil
}

func (c *Codec) EncodeOwnerMethodCall() ([]byte, error) {
	return c.abi.Pack("owner")
}

func (c *Codec) DecodeOwnerMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["owner"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodePlaceBetMethodCall(in PlaceBetInput) ([]byte, error) {
	return c.abi.Pack("placeBet", in.MarketId, in.EncryptedBet)
}

func (c *Codec) EncodeRenounceOwnershipMethodCall() ([]byte, error) {
	return c.abi.Pack("renounceOwnership")
}

func (c *Codec) EncodeRequestSettlementMethodCall(in RequestSettlementInput) ([]byte, error) {
	return c.abi.Pack("requestSettlement", in.MarketId)
}

func (c *Codec) EncodeRevealBetsMethodCall(in RevealBetsInput) ([]byte, error) {
	return c.abi.Pack("revealBets", in.MarketId, in.BetIndices, in.Options, in.Amounts)
}

func (c *Codec) EncodeSetAdminMethodCall(in SetAdminInput) ([]byte, error) {
	return c.abi.Pack("setAdmin", in.Admin, in.Status)
}

func (c *Codec) EncodeSetTrustedCREMethodCall(in SetTrustedCREInput) ([]byte, error) {
	return c.abi.Pack("setTrustedCRE", in.Addr)
}

func (c *Codec) EncodeSubmitResultMethodCall(in SubmitResultInput) ([]byte, error) {
	return c.abi.Pack("submitResult", in.MarketId, in.Outcome, in.Signature)
}

func (c *Codec) EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error) {
	return c.abi.Pack("transferOwnership", in.NewOwner)
}

func (c *Codec) EncodeTrustedCREAddressMethodCall() ([]byte, error) {
	return c.abi.Pack("trustedCREAddress")
}

func (c *Codec) DecodeTrustedCREAddressMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["trustedCREAddress"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) BetPlacedLogHash() []byte {
	return c.abi.Events["BetPlaced"].ID.Bytes()
}

func (c *Codec) EncodeBetPlacedTopics(
	evt abi.Event,
	values []BetPlacedTopics,
) ([]*evm.TopicValues, error) {
	var marketIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.MarketId).IsZero() {
			marketIdRule = append(marketIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.MarketId)
		if err != nil {
			return nil, err
		}
		marketIdRule = append(marketIdRule, fieldVal)
	}
	var userRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.User).IsZero() {
			userRule = append(userRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.User)
		if err != nil {
			return nil, err
		}
		userRule = append(userRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		marketIdRule,
		userRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeBetPlaced decodes a log into a BetPlaced struct.
func (c *Codec) DecodeBetPlaced(log *evm.Log) (*BetPlacedDecoded, error) {
	event := new(BetPlacedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "BetPlaced", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["BetPlaced"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) MarketCreatedLogHash() []byte {
	return c.abi.Events["MarketCreated"].ID.Bytes()
}

func (c *Codec) EncodeMarketCreatedTopics(
	evt abi.Event,
	values []MarketCreatedTopics,
) ([]*evm.TopicValues, error) {
	var marketIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.MarketId).IsZero() {
			marketIdRule = append(marketIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.MarketId)
		if err != nil {
			return nil, err
		}
		marketIdRule = append(marketIdRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		marketIdRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeMarketCreated decodes a log into a MarketCreated struct.
func (c *Codec) DecodeMarketCreated(log *evm.Log) (*MarketCreatedDecoded, error) {
	event := new(MarketCreatedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "MarketCreated", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["MarketCreated"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) MarketResolvedLogHash() []byte {
	return c.abi.Events["MarketResolved"].ID.Bytes()
}

func (c *Codec) EncodeMarketResolvedTopics(
	evt abi.Event,
	values []MarketResolvedTopics,
) ([]*evm.TopicValues, error) {
	var marketIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.MarketId).IsZero() {
			marketIdRule = append(marketIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.MarketId)
		if err != nil {
			return nil, err
		}
		marketIdRule = append(marketIdRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		marketIdRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeMarketResolved decodes a log into a MarketResolved struct.
func (c *Codec) DecodeMarketResolved(log *evm.Log) (*MarketResolvedDecoded, error) {
	event := new(MarketResolvedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "MarketResolved", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["MarketResolved"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) OwnershipTransferredLogHash() []byte {
	return c.abi.Events["OwnershipTransferred"].ID.Bytes()
}

func (c *Codec) EncodeOwnershipTransferredTopics(
	evt abi.Event,
	values []OwnershipTransferredTopics,
) ([]*evm.TopicValues, error) {
	var previousOwnerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.PreviousOwner).IsZero() {
			previousOwnerRule = append(previousOwnerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.PreviousOwner)
		if err != nil {
			return nil, err
		}
		previousOwnerRule = append(previousOwnerRule, fieldVal)
	}
	var newOwnerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.NewOwner).IsZero() {
			newOwnerRule = append(newOwnerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.NewOwner)
		if err != nil {
			return nil, err
		}
		newOwnerRule = append(newOwnerRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		previousOwnerRule,
		newOwnerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeOwnershipTransferred decodes a log into a OwnershipTransferred struct.
func (c *Codec) DecodeOwnershipTransferred(log *evm.Log) (*OwnershipTransferredDecoded, error) {
	event := new(OwnershipTransferredDecoded)
	if err := c.abi.UnpackIntoInterface(event, "OwnershipTransferred", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["OwnershipTransferred"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) SettlementRequestedLogHash() []byte {
	return c.abi.Events["SettlementRequested"].ID.Bytes()
}

func (c *Codec) EncodeSettlementRequestedTopics(
	evt abi.Event,
	values []SettlementRequestedTopics,
) ([]*evm.TopicValues, error) {
	var marketIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.MarketId).IsZero() {
			marketIdRule = append(marketIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.MarketId)
		if err != nil {
			return nil, err
		}
		marketIdRule = append(marketIdRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		marketIdRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeSettlementRequested decodes a log into a SettlementRequested struct.
func (c *Codec) DecodeSettlementRequested(log *evm.Log) (*SettlementRequestedDecoded, error) {
	event := new(SettlementRequestedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "SettlementRequested", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["SettlementRequested"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) WinningsClaimedLogHash() []byte {
	return c.abi.Events["WinningsClaimed"].ID.Bytes()
}

func (c *Codec) EncodeWinningsClaimedTopics(
	evt abi.Event,
	values []WinningsClaimedTopics,
) ([]*evm.TopicValues, error) {
	var marketIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.MarketId).IsZero() {
			marketIdRule = append(marketIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.MarketId)
		if err != nil {
			return nil, err
		}
		marketIdRule = append(marketIdRule, fieldVal)
	}
	var userRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.User).IsZero() {
			userRule = append(userRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.User)
		if err != nil {
			return nil, err
		}
		userRule = append(userRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		marketIdRule,
		userRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeWinningsClaimed decodes a log into a WinningsClaimed struct.
func (c *Codec) DecodeWinningsClaimed(log *evm.Log) (*WinningsClaimedDecoded, error) {
	event := new(WinningsClaimedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "WinningsClaimed", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["WinningsClaimed"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c ConfidentialMarket) OrderMarketid(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeOrderMarketidMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*big.Int](*new(*big.Int), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (*big.Int, error) {
		return c.Codec.DecodeOrderMarketidMethodOutput(response.Data)
	})

}

func (c ConfidentialMarket) Admins(
	runtime cre.Runtime,
	args AdminsInput,
	blockNumber *big.Int,
) cre.Promise[bool] {
	calldata, err := c.Codec.EncodeAdminsMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[bool](*new(bool), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (bool, error) {
		return c.Codec.DecodeAdminsMethodOutput(response.Data)
	})

}

func (c ConfidentialMarket) MarketTotalPools(
	runtime cre.Runtime,
	args MarketTotalPoolsInput,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeMarketTotalPoolsMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[*big.Int](*new(*big.Int), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (*big.Int, error) {
		return c.Codec.DecodeMarketTotalPoolsMethodOutput(response.Data)
	})

}

func (c ConfidentialMarket) Markets(
	runtime cre.Runtime,
	args MarketsInput,
	blockNumber *big.Int,
) cre.Promise[MarketsOutput] {
	calldata, err := c.Codec.EncodeMarketsMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[MarketsOutput](MarketsOutput{}, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (MarketsOutput, error) {
		return c.Codec.DecodeMarketsMethodOutput(response.Data)
	})

}

func (c ConfidentialMarket) Owner(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeOwnerMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeOwnerMethodOutput(response.Data)
	})

}

func (c ConfidentialMarket) TrustedCREAddress(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeTrustedCREAddressMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeTrustedCREAddressMethodOutput(response.Data)
	})

}

func (c ConfidentialMarket) WriteReport(
	runtime cre.Runtime,
	report *cre.Report,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver:  c.Address.Bytes(),
		Report:    report,
		GasConfig: gasConfig,
	})
}

// DecodeOwnableInvalidOwnerError decodes a OwnableInvalidOwner error from revert data.
func (c *ConfidentialMarket) DecodeOwnableInvalidOwnerError(data []byte) (*OwnableInvalidOwner, error) {
	args := c.ABI.Errors["OwnableInvalidOwner"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	owner, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for owner in OwnableInvalidOwner error")
	}

	return &OwnableInvalidOwner{
		Owner: owner,
	}, nil
}

// Error implements the error interface for OwnableInvalidOwner.
func (e *OwnableInvalidOwner) Error() string {
	return fmt.Sprintf("OwnableInvalidOwner error: owner=%v;", e.Owner)
}

// DecodeOwnableUnauthorizedAccountError decodes a OwnableUnauthorizedAccount error from revert data.
func (c *ConfidentialMarket) DecodeOwnableUnauthorizedAccountError(data []byte) (*OwnableUnauthorizedAccount, error) {
	args := c.ABI.Errors["OwnableUnauthorizedAccount"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	account, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for account in OwnableUnauthorizedAccount error")
	}

	return &OwnableUnauthorizedAccount{
		Account: account,
	}, nil
}

// Error implements the error interface for OwnableUnauthorizedAccount.
func (e *OwnableUnauthorizedAccount) Error() string {
	return fmt.Sprintf("OwnableUnauthorizedAccount error: account=%v;", e.Account)
}

func (c *ConfidentialMarket) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	case common.Bytes2Hex(c.ABI.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]):
		return c.DecodeOwnableInvalidOwnerError(data)
	case common.Bytes2Hex(c.ABI.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]):
		return c.DecodeOwnableUnauthorizedAccountError(data)
	default:
		return nil, errors.New("unknown error selector")
	}
}

// BetPlacedTrigger wraps the raw log trigger and provides decoded BetPlacedDecoded data
type BetPlacedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                     // Embed the raw trigger
	contract                        *ConfidentialMarket // Keep reference for decoding
}

// Adapt method that decodes the log into BetPlaced data
func (t *BetPlacedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[BetPlacedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeBetPlaced(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode BetPlaced log: %w", err)
	}

	return &bindings.DecodedLog[BetPlacedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ConfidentialMarket) LogTriggerBetPlacedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []BetPlacedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[BetPlacedDecoded]], error) {
	event := c.ABI.Events["BetPlaced"]
	topics, err := c.Codec.EncodeBetPlacedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for BetPlaced: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &BetPlacedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ConfidentialMarket) FilterLogsBetPlaced(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.BetPlacedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// MarketCreatedTrigger wraps the raw log trigger and provides decoded MarketCreatedDecoded data
type MarketCreatedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                     // Embed the raw trigger
	contract                        *ConfidentialMarket // Keep reference for decoding
}

// Adapt method that decodes the log into MarketCreated data
func (t *MarketCreatedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[MarketCreatedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeMarketCreated(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode MarketCreated log: %w", err)
	}

	return &bindings.DecodedLog[MarketCreatedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ConfidentialMarket) LogTriggerMarketCreatedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []MarketCreatedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[MarketCreatedDecoded]], error) {
	event := c.ABI.Events["MarketCreated"]
	topics, err := c.Codec.EncodeMarketCreatedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for MarketCreated: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &MarketCreatedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ConfidentialMarket) FilterLogsMarketCreated(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.MarketCreatedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// MarketResolvedTrigger wraps the raw log trigger and provides decoded MarketResolvedDecoded data
type MarketResolvedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                     // Embed the raw trigger
	contract                        *ConfidentialMarket // Keep reference for decoding
}

// Adapt method that decodes the log into MarketResolved data
func (t *MarketResolvedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[MarketResolvedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeMarketResolved(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode MarketResolved log: %w", err)
	}

	return &bindings.DecodedLog[MarketResolvedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ConfidentialMarket) LogTriggerMarketResolvedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []MarketResolvedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[MarketResolvedDecoded]], error) {
	event := c.ABI.Events["MarketResolved"]
	topics, err := c.Codec.EncodeMarketResolvedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for MarketResolved: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &MarketResolvedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ConfidentialMarket) FilterLogsMarketResolved(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.MarketResolvedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// OwnershipTransferredTrigger wraps the raw log trigger and provides decoded OwnershipTransferredDecoded data
type OwnershipTransferredTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                     // Embed the raw trigger
	contract                        *ConfidentialMarket // Keep reference for decoding
}

// Adapt method that decodes the log into OwnershipTransferred data
func (t *OwnershipTransferredTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[OwnershipTransferredDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeOwnershipTransferred(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode OwnershipTransferred log: %w", err)
	}

	return &bindings.DecodedLog[OwnershipTransferredDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ConfidentialMarket) LogTriggerOwnershipTransferredLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OwnershipTransferredTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OwnershipTransferredDecoded]], error) {
	event := c.ABI.Events["OwnershipTransferred"]
	topics, err := c.Codec.EncodeOwnershipTransferredTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for OwnershipTransferred: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &OwnershipTransferredTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ConfidentialMarket) FilterLogsOwnershipTransferred(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.OwnershipTransferredLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// SettlementRequestedTrigger wraps the raw log trigger and provides decoded SettlementRequestedDecoded data
type SettlementRequestedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                     // Embed the raw trigger
	contract                        *ConfidentialMarket // Keep reference for decoding
}

// Adapt method that decodes the log into SettlementRequested data
func (t *SettlementRequestedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[SettlementRequestedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeSettlementRequested(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode SettlementRequested log: %w", err)
	}

	return &bindings.DecodedLog[SettlementRequestedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ConfidentialMarket) LogTriggerSettlementRequestedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []SettlementRequestedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[SettlementRequestedDecoded]], error) {
	event := c.ABI.Events["SettlementRequested"]
	topics, err := c.Codec.EncodeSettlementRequestedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for SettlementRequested: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &SettlementRequestedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ConfidentialMarket) FilterLogsSettlementRequested(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.SettlementRequestedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// WinningsClaimedTrigger wraps the raw log trigger and provides decoded WinningsClaimedDecoded data
type WinningsClaimedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                     // Embed the raw trigger
	contract                        *ConfidentialMarket // Keep reference for decoding
}

// Adapt method that decodes the log into WinningsClaimed data
func (t *WinningsClaimedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[WinningsClaimedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeWinningsClaimed(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode WinningsClaimed log: %w", err)
	}

	return &bindings.DecodedLog[WinningsClaimedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ConfidentialMarket) LogTriggerWinningsClaimedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []WinningsClaimedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[WinningsClaimedDecoded]], error) {
	event := c.ABI.Events["WinningsClaimed"]
	topics, err := c.Codec.EncodeWinningsClaimedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for WinningsClaimed: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &WinningsClaimedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ConfidentialMarket) FilterLogsWinningsClaimed(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.WinningsClaimedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}
