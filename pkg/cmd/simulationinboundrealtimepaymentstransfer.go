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

var simulationsInboundRealTimePaymentsTransfersCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates an\n[Inbound Real-Time Payments Transfer](#inbound-real-time-payments-transfers) to\nyour account. Real-Time Payments are a beta feature.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-number-id",
			Usage:    "The identifier of the Account Number the inbound Real-Time Payments Transfer is for.",
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
			Name:     "request-for-payment-id",
			Usage:    "The identifier of a pending Request for Payment that this transfer will fulfill.",
			BodyPath: "request_for_payment_id",
		},
		&requestflag.Flag[string]{
			Name:     "unstructured-remittance-information",
			Usage:    "Additional information included with the transfer.",
			BodyPath: "unstructured_remittance_information",
		},
	},
	Action:          handleSimulationsInboundRealTimePaymentsTransfersCreate,
	HideHelpCommand: true,
}

func handleSimulationsInboundRealTimePaymentsTransfersCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationInboundRealTimePaymentsTransferNewParams{}

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
	_, err = client.Simulations.InboundRealTimePaymentsTransfers.New(ctx, params, options...)
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
		Title:          "simulations:inbound-real-time-payments-transfers create",
		Transform:      transform,
	})
}
