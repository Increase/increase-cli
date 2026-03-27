// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/Increase/increase-cli/internal/apiquery"
	"github.com/Increase/increase-cli/internal/requestflag"
	"github.com/Increase/increase-go"
	"github.com/Increase/increase-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var simulationsInboundWireTransfersCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates an [Inbound Wire Transfer](#inbound-wire-transfers) to your account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-number-id",
			Usage:    "The identifier of the Account Number the inbound Wire Transfer is for.",
			Required: true,
			BodyPath: "account_number_id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The transfer amount in cents. Must be positive.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-address-line1",
			Usage:    "The sending bank will set creditor_address_line1 in production. You can simulate any value here.",
			BodyPath: "creditor_address_line1",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-address-line2",
			Usage:    "The sending bank will set creditor_address_line2 in production. You can simulate any value here.",
			BodyPath: "creditor_address_line2",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-address-line3",
			Usage:    "The sending bank will set creditor_address_line3 in production. You can simulate any value here.",
			BodyPath: "creditor_address_line3",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-name",
			Usage:    "The sending bank will set creditor_name in production. You can simulate any value here.",
			BodyPath: "creditor_name",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-address-line1",
			Usage:    "The sending bank will set debtor_address_line1 in production. You can simulate any value here.",
			BodyPath: "debtor_address_line1",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-address-line2",
			Usage:    "The sending bank will set debtor_address_line2 in production. You can simulate any value here.",
			BodyPath: "debtor_address_line2",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-address-line3",
			Usage:    "The sending bank will set debtor_address_line3 in production. You can simulate any value here.",
			BodyPath: "debtor_address_line3",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-name",
			Usage:    "The sending bank will set debtor_name in production. You can simulate any value here.",
			BodyPath: "debtor_name",
		},
		&requestflag.Flag[string]{
			Name:     "end-to-end-identification",
			Usage:    "The sending bank will set end_to_end_identification in production. You can simulate any value here.",
			BodyPath: "end_to_end_identification",
		},
		&requestflag.Flag[string]{
			Name:     "instructing-agent-routing-number",
			Usage:    "The sending bank will set instructing_agent_routing_number in production. You can simulate any value here.",
			BodyPath: "instructing_agent_routing_number",
		},
		&requestflag.Flag[string]{
			Name:     "instruction-identification",
			Usage:    "The sending bank will set instruction_identification in production. You can simulate any value here.",
			BodyPath: "instruction_identification",
		},
		&requestflag.Flag[string]{
			Name:     "unique-end-to-end-transaction-reference",
			Usage:    "The sending bank will set unique_end_to_end_transaction_reference in production. You can simulate any value here.",
			BodyPath: "unique_end_to_end_transaction_reference",
		},
		&requestflag.Flag[string]{
			Name:     "unstructured-remittance-information",
			Usage:    "The sending bank will set unstructured_remittance_information in production. You can simulate any value here.",
			BodyPath: "unstructured_remittance_information",
		},
		&requestflag.Flag[string]{
			Name:     "wire-drawdown-request-id",
			Usage:    "The identifier of a Wire Drawdown Request the inbound Wire Transfer is fulfilling.",
			BodyPath: "wire_drawdown_request_id",
		},
	},
	Action:          handleSimulationsInboundWireTransfersCreate,
	HideHelpCommand: true,
}

func handleSimulationsInboundWireTransfersCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationInboundWireTransferNewParams{}

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
	_, err = client.Simulations.InboundWireTransfers.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "simulations:inbound-wire-transfers create", obj, format, transform)
}
