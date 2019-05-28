package mint

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/tendermint/tendermint/crypto"
)

var minterKey = []byte{0x00} // the one key to use for the keeper store

const (
	// default paramspace for params keeper
	DefaultParamspace = ModuleName

	// StoreKey is the default store key for mint
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the minting store.
	QuerierRoute = StoreKey
)

// ModuleAddress distribution module account address
var ModuleAddress = sdk.AccAddress(crypto.AddressHash([]byte(ModuleName)))

// keeper of the mint store
type Keeper struct {
	storeKey     sdk.StoreKey
	cdc          *codec.Codec
	paramSpace   params.Subspace
	sk           StakingKeeper
	supplyKeeper SupplyKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramSpace params.Subspace,
	sk StakingKeeper, supplyKeeper SupplyKeeper) Keeper {

	keeper := Keeper{
		storeKey:     key,
		cdc:          cdc,
		paramSpace:   paramSpace.WithKeyTable(ParamKeyTable()),
		sk:           sk,
		supplyKeeper: supplyKeeper,
	}
	return keeper
}

//______________________________________________________________________

// get the minter
func (k Keeper) GetMinter(ctx sdk.Context) (minter Minter) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(minterKey)
	if b == nil {
		panic("stored minter should not have been nil")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &minter)
	return
}

// set the minter
func (k Keeper) SetMinter(ctx sdk.Context, minter Minter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(minter)
	store.Set(minterKey, b)
}

//______________________________________________________________________

// GetParams returns the total set of minting parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of minting parameters.
func (k Keeper) SetParams(ctx sdk.Context, params Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

//______________________________________________________________________

// GetMinterAccount returns the mint ModuleAccount
func (k Keeper) GetMinterAccount(ctx sdk.Context) supply.ModuleAccount {
	return k.supplyKeeper.GetModuleAccountByName(ctx, ModuleName)
}

// SetMinterAccount stores the minter account
func (k Keeper) SetMinterAccount(ctx sdk.Context, macc supply.ModuleAccount) {
	if macc.Name() != ModuleName {
		panic(fmt.Sprintf("invalid name for minter module account (%s ≠ %s)", macc.Name(), ModuleName))
	}

	k.supplyKeeper.SetModuleAccount(ctx, macc)
}
