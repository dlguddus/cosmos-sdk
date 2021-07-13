package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/x/tax/types"
)

// GetQueryCmd returns a root CLI command handler for all x/tax query commands.
func GetQueryCmd() *cobra.Command {
	taxQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the tax module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	taxQueryCmd.AddCommand(
	// TODO: add query commands
	// GetCmdQueryTax(),
	)

	return taxQueryCmd
}
