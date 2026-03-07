// Code generated — DO NOT EDIT.

//go:build !wasip1

package confidential_market

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	evmmock "github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/mock"
)

var (
	_ = errors.New
	_ = fmt.Errorf
	_ = big.NewInt
	_ = common.Big1
)

// ConfidentialMarketMock is a mock implementation of ConfidentialMarket for testing.
type ConfidentialMarketMock struct {
	OrderMarketid    func() (*big.Int, error)
	Admins           func(AdminsInput) (bool, error)
	MarketTotalPools func(MarketTotalPoolsInput) (*big.Int, error)
	Markets          func(MarketsInput) (MarketsOutput, error)
	Owner            func() (common.Address, error)
	TrustedSigner    func() (common.Address, error)
}

// NewConfidentialMarketMock creates a new ConfidentialMarketMock for testing.
func NewConfidentialMarketMock(address common.Address, clientMock *evmmock.ClientCapability) *ConfidentialMarketMock {
	mock := &ConfidentialMarketMock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi
	_ = abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["OrderMarketid"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.OrderMarketid == nil {
				return nil, errors.New("OrderMarketid method not mocked")
			}
			result, err := mock.OrderMarketid()
			if err != nil {
				return nil, err
			}
			return abi.Methods["OrderMarketid"].Outputs.Pack(result)
		},
		string(abi.Methods["admins"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Admins == nil {
				return nil, errors.New("admins method not mocked")
			}
			inputs := abi.Methods["admins"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := AdminsInput{
				Arg0: values[0].(common.Address),
			}

			result, err := mock.Admins(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["admins"].Outputs.Pack(result)
		},
		string(abi.Methods["marketTotalPools"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.MarketTotalPools == nil {
				return nil, errors.New("marketTotalPools method not mocked")
			}
			inputs := abi.Methods["marketTotalPools"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := MarketTotalPoolsInput{
				Arg0: values[0].(*big.Int),
			}

			result, err := mock.MarketTotalPools(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["marketTotalPools"].Outputs.Pack(result)
		},
		string(abi.Methods["markets"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Markets == nil {
				return nil, errors.New("markets method not mocked")
			}
			inputs := abi.Methods["markets"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := MarketsInput{
				Arg0: values[0].(*big.Int),
			}

			result, err := mock.Markets(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["markets"].Outputs.Pack(
				result.Id,
				result.Description,
				result.State,
				result.CreateTimeStamp,
				result.CloseTime,
				result.ResolveWindow,
				result.ResolveTimeStamp,
				result.Outcome,
				result.Resolved,
			)
		},
		string(abi.Methods["owner"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Owner == nil {
				return nil, errors.New("owner method not mocked")
			}
			result, err := mock.Owner()
			if err != nil {
				return nil, err
			}
			return abi.Methods["owner"].Outputs.Pack(result)
		},
		string(abi.Methods["trustedSigner"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.TrustedSigner == nil {
				return nil, errors.New("trustedSigner method not mocked")
			}
			result, err := mock.TrustedSigner()
			if err != nil {
				return nil, err
			}
			return abi.Methods["trustedSigner"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
