// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/Increase/increase-cli/internal/apiquery"
	"github.com/Increase/increase-cli/internal/requestflag"
	"github.com/Increase/increase-go"
	"github.com/Increase/increase-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var simulationsInboundFednowTransfersCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates an [Inbound FedNow Transfer](#inbound-fednow-transfers) to your\naccount.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-number-id",
			Usage:    "The identifier of the Account Number the inbound FedNow Transfer is for.",
			Required: true,
			BodyPath: "account_number_id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The transfer amount in USD cents. Must be positive.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-account-number",
			Usage:    "The account number of the account that sent the transfer.",
			BodyPath: "debtor_account_number",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-name",
			Usage:    "The name provided by the sender of the transfer.",
			BodyPath: "debtor_name",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-routing-number",
			Usage:    "The routing number of the account that sent the transfer.",
			BodyPath: "debtor_routing_number",
		},
		&requestflag.Flag[string]{
			Name:     "unstructured-remittance-information",
			Usage:    "Additional information included with the transfer.",
			BodyPath: "unstructured_remittance_information",
		},
	},
	Action:          handleSimulationsInboundFednowTransfersCreate,
	HideHelpCommand: true,
}

func handleSimulationsInboundFednowTransfersCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationInboundFednowTransferNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatDots,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simulations.InboundFednowTransfers.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "simulations:inbound-fednow-transfers create",
		Transform:      transform,
	})
}
