package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/farming/types"
)

// NewPlan sets the next plan number to a given plan interface
func (k Keeper) NewPlan(ctx sdk.Context, plan types.PlanI) types.PlanI {
	if err := plan.SetId(k.GetNextPlanIdWithUpdate(ctx)); err != nil {
		panic(err)
	}

	return plan
}

// GetPlan implements PlanI.
func (k Keeper) GetPlan(ctx sdk.Context, id uint64) (plan types.PlanI, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPlanKey(id))
	if bz == nil {
		return plan, false
	}

	return k.decodePlan(bz), true
}

// GetAllPlans returns all plans in the Keeper.
func (k Keeper) GetAllPlans(ctx sdk.Context) (plans []types.PlanI) {
	k.IterateAllPlans(ctx, func(plan types.PlanI) (stop bool) {
		plans = append(plans, plan)
		return false
	})

	return plans
}

// SetPlan implements PlanI.
func (k Keeper) SetPlan(ctx sdk.Context, plan types.PlanI) {
	id := plan.GetId()
	store := ctx.KVStore(k.storeKey)

	bz, err := k.MarshalPlan(plan)
	if err != nil {
		panic(err)
	}

	store.Set(types.GetPlanKey(id), bz)
}

// RemovePlan removes an plan for the plan mapper store.
// NOTE: this will cause supply invariant violation if called
func (k Keeper) RemovePlan(ctx sdk.Context, plan types.PlanI) {
	id := plan.GetId()
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetPlanKey(id))
}

// IterateAllPlans iterates over all the stored plans and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateAllPlans(ctx sdk.Context, cb func(plan types.PlanI) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PlanKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		plan := k.decodePlan(iterator.Value())

		if cb(plan) {
			break
		}
	}
}

// GetPlanByFarmerAddrIndex reads from kvstore and return a specific Plan indexed by given farmer address
func (k Keeper) GetPlansByFarmerAddrIndex(ctx sdk.Context, farmerAcc sdk.AccAddress) (plans []types.PlanI) {
	k.IteratePlansByFarmerAddr(ctx, farmerAcc, func(plan types.PlanI) bool {
		plans = append(plans, plan)
		return false
	})

	return plans
}

// IterateAllPlans iterates over all the stored plans and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IteratePlansByFarmerAddr(ctx sdk.Context, farmerAcc sdk.AccAddress, cb func(plan types.PlanI) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetPlansByFarmerAddrIndexKey(farmerAcc))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		plan := k.decodePlan(iterator.Value())

		if cb(plan) {
			break
		}
	}
}

// SetPlanByFarmerAddrIndex sets Index by FarmerAddr
func (k Keeper) SetPlanByFarmerAddrIndex(ctx sdk.Context, plan types.PlanI, farmerAcc sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: plan.GetId()})
	store.Set(types.GetPlanByFarmerAddrIndexKey(farmerAcc, plan.GetId()), b)
}
