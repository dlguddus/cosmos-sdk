package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Pool - tracking bonded and not-bonded token supply of the bond denomination
type Pool struct {
	NotBondedTokens sdk.Coins `json:"not_bonded_tokens"` // tokens which are not bonded to a validator (unbonded or unbonding)
	BondedTokens    sdk.Coins `json:"bonded_tokens"`     // tokens which are currently bonded to a validator
}

// NewPool creates a new Pool instance
func NewPool(unbonded, bonded sdk.Coins) Pool {
	return Pool{
		NotBondedTokens: unbonded,
		BondedTokens:    bonded,
	}
}

// String returns a human readable string representation of a pool.
func (p Pool) String() string {
	return fmt.Sprintf(`Pool:	
  Not Bonded Tokens:  %s	
  Bonded Tokens:      %s`, p.NotBondedTokens,
		p.BondedTokens)
}
