package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/dymensionxyz/dymension/v3/utils"
	"github.com/dymensionxyz/dymension/v3/x/denommetadata/types"

	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// NewCmdSubmitCreateDenomMetadataProposal broadcasts a CreateMetadataProposal message.
func NewCmdSubmitCreateDenomMetadataProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-denom-metadata-proposal denommetadata.json [flags]",
		Short: "proposal to create new denom metadata for a specific token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposal, deposit, err := utils.ParseProposal(cmd)
			if err != nil {
				return err
			}

			path := args[0]

			var metadatas []banktypes.Metadata
			metadatas, err = utils.ParseJsonFromFile[banktypes.Metadata](path)
			if err != nil {
				return err
			}

			for _, metadata := range metadatas {
				err = metadata.Validate()
				if err != nil {
					return err
				}
			}

			content := types.NewCreateMetadataProposal(proposal.Title, proposal.Description, metadatas)
			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)
			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	cmd.Flags().String(govcli.FlagTitle, "", "The proposal title")
	cmd.Flags().String(govcli.FlagDescription, "", "The proposal description")
	cmd.Flags().String(govcli.FlagDeposit, "", "The proposal deposit")

	return cmd
}
