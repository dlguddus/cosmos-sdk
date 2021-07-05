package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// farming module sentinel errors
var (
	ErrPlanNotExists             = sdkerrors.Register(ModuleName, 2, "plan not exists")
	ErrPlanTypeNotExists         = sdkerrors.Register(ModuleName, 3, "plan type not exists")
	ErrInvalidFarmingPoolAddress = sdkerrors.Register(ModuleName, 4, "invalid farming pool address")
	ErrInvalidFarmerAddress      = sdkerrors.Register(ModuleName, 5, "invalid farmer address")

	// TODO: TBD more err types for farming module
)