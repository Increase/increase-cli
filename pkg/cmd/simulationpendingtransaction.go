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

var simulationsPendingTransactionsReleaseInboundFundsHold = cli.Command{
	Name:    "release-inbound-funds-hold",
	Usage:   "This endpoint simulates immediately releasing an Inbound Funds Hold, which might\nbe created as a result of, for example, an ACH debit.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "pending-transaction-id",
			Usage:    "The pending transaction to release. The pending transaction must have a `inbound_funds_hold` source.",
			Required: true,
		},
	},
	Action:          handleSimulationsPendingTransactionsReleaseInboundFundsHold,
	HideHelpCommand: true,
}

func handleSimulationsPendingTransactionsReleaseInboundFundsHold(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pending-transaction-id") && len(unusedArgs) > 0 {
		cmd.Set("pending-transaction-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatDots,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simulations.PendingTransactions.ReleaseInboundFundsHold(ctx, cmd.Value("pending-transaction-id").(string), options...)
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
		Title:          "simulations:pending-transactions release-inbound-funds-hold",
		Transform:      transform,
	})
}
