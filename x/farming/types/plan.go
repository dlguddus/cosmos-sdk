package types

import (
	"time"

	"github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// TODO: Unimplemented
	//_ PlanI                           = (*BasePlan)(nil)
	_ codectypes.UnpackInterfacesMessage = (*BasePlan)(nil)
)

// NewBasePlan creates a new BasePlan object
//nolint:interfacer
func NewBasePlan(address sdk.AccAddress, pubKey cryptotypes.PubKey, PlanNumber, sequence uint64) *BasePlan {
	// TODO: Unimplemented
	return nil
}

// ProtoBasePlan - a prototype function for BasePlan
func ProtoBasePlan() PlanI {
	// TODO: Unimplemented
	return nil
	//return &BasePlan{}
}

// Validate checks for errors on the Plan fields
func (acc BasePlan) Validate() error {
	// TODO: Unimplemented
	return nil
}

func (acc BasePlan) String() string {
	out, _ := acc.MarshalYAML()
	return out.(string)
}

// MarshalYAML returns the YAML representation of an Plan.
func (acc BasePlan) MarshalYAML() (interface{}, error) {
	bz, err := codec.MarshalYAML(codec.NewProtoCodec(codectypes.NewInterfaceRegistry()), &acc)
	if err != nil {
		return nil, err
	}
	return string(bz), err
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (acc BasePlan) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	// TODO: Unimplemented
	return nil
}

type PlanI interface {
	proto.Message

	GetId() uint64
	SetId(uint64) error

	GetType() int32
	SetType(int32) error

	GetFarmingPoolAddress() sdk.AccAddress
	SetFarmingPoolAddress(sdk.AccAddress) error

	GetDistributionPoolAddress() sdk.AccAddress
	SetDistributionPoolAddress(sdk.AccAddress) error

	GetTerminationAddress() sdk.AccAddress
	SetTerminationAddress(sdk.AccAddress) error

	GetStakingReserveAddress() sdk.AccAddress
	SetStakingReserveAddress(sdk.AccAddress) error

	GetStakingCoinsWeight() sdk.DecCoins
	SetStakingCoinsWeight(sdk.DecCoins) error

	GetStartTime() time.Time
	SetStartTime(time.Time) error

	GetEndTime() time.Time
	SetEndTime(time.Time) error

	GetEpochDays() uint32
	SetEpochDays(uint32) error

	GetDistributionMethod() string
	GetDistributionThisEpoch() sdk.Coins
	IsTermitated() bool

	String() string
}