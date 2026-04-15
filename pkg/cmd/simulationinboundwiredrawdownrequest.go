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

var simulationsInboundWireDrawdownRequestsCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates receiving an\n[Inbound Wire Drawdown Request](#inbound-wire-drawdown-requests).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The amount being requested in cents.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-account-number",
			Usage:    "The creditor's account number.",
			Required: true,
			BodyPath: "creditor_account_number",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-routing-number",
			Usage:    "The creditor's routing number.",
			Required: true,
			BodyPath: "creditor_routing_number",
		},
		&requestflag.Flag[string]{
			Name:     "currency",
			Usage:    `The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being requested. Will always be "USD".`,
			Required: true,
			BodyPath: "currency",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-account-number-id",
			Usage:    "The Account Number to which the recipient of this request is being requested to send funds from.",
			Required: true,
			BodyPath: "recipient_account_number_id",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-address-line1",
			Usage:    "A free-form address field set by the sender representing the first line of the creditor's address.",
			BodyPath: "creditor_address_line1",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-address-line2",
			Usage:    "A free-form address field set by the sender representing the second line of the creditor's address.",
			BodyPath: "creditor_address_line2",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-address-line3",
			Usage:    "A free-form address field set by the sender representing the third line of the creditor's address.",
			BodyPath: "creditor_address_line3",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-name",
			Usage:    "A free-form name field set by the sender representing the creditor's name.",
			BodyPath: "creditor_name",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-account-number",
			Usage:    "The debtor's account number.",
			BodyPath: "debtor_account_number",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-address-line1",
			Usage:    "A free-form address field set by the sender representing the first line of the debtor's address.",
			BodyPath: "debtor_address_line1",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-address-line2",
			Usage:    "A free-form address field set by the sender representing the second line of the debtor's address.",
			BodyPath: "debtor_address_line2",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-address-line3",
			Usage:    "A free-form address field set by the sender.",
			BodyPath: "debtor_address_line3",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-name",
			Usage:    "A free-form name field set by the sender representing the debtor's name.",
			BodyPath: "debtor_name",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-routing-number",
			Usage:    "The debtor's routing number.",
			BodyPath: "debtor_routing_number",
		},
		&requestflag.Flag[string]{
			Name:     "end-to-end-identification",
			Usage:    "A free-form reference string set by the sender, to help identify the transfer.",
			BodyPath: "end_to_end_identification",
		},
		&requestflag.Flag[string]{
			Name:     "instruction-identification",
			Usage:    "The sending bank's identifier for the wire transfer.",
			BodyPath: "instruction_identification",
		},
		&requestflag.Flag[string]{
			Name:     "unique-end-to-end-transaction-reference",
			Usage:    "The Unique End-to-end Transaction Reference ([UETR](https://www.swift.com/payments/what-unique-end-end-transaction-reference-uetr)) of the transfer.",
			BodyPath: "unique_end_to_end_transaction_reference",
		},
		&requestflag.Flag[string]{
			Name:     "unstructured-remittance-information",
			Usage:    "A free-form message set by the sender.",
			BodyPath: "unstructured_remittance_information",
		},
	},
	Action:          handleSimulationsInboundWireDrawdownRequestsCreate,
	HideHelpCommand: true,
}

func handleSimulationsInboundWireDrawdownRequestsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationInboundWireDrawdownRequestNewParams{}

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
	_, err = client.Simulations.InboundWireDrawdownRequests.New(ctx, params, options...)
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
		Title:          "simulations:inbound-wire-drawdown-requests create",
		Transform:      transform,
	})
}
