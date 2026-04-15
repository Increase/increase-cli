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

var simulationsInboundACHTransfersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Simulates an inbound ACH transfer to your account. This imitates initiating a\ntransfer to an Increase account from a different financial institution. The\ntransfer may be either a credit or a debit depending on if the `amount` is\npositive or negative. The result of calling this API will contain the created\ntransfer. You can pass a `resolve_at` parameter to allow for a window to\n[action on the Inbound ACH Transfer](https://increase.com/documentation/receiving-ach-transfers).\nAlternatively, if you don't pass the `resolve_at` parameter the result will\ncontain either a [Transaction](#transactions) or a\n[Declined Transaction](#declined-transactions) depending on whether or not the\ntransfer is allowed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-number-id",
			Usage:    "The identifier of the Account Number the inbound ACH Transfer is for.",
			Required: true,
			BodyPath: "account_number_id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The transfer amount in cents. A positive amount originates a credit transfer pushing funds to the receiving account. A negative amount originates a debit transfer pulling funds from the receiving account.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "addenda",
			Usage:    "Additional information to include in the transfer.",
			BodyPath: "addenda",
		},
		&requestflag.Flag[string]{
			Name:     "company-descriptive-date",
			Usage:    "The description of the date of the transfer.",
			BodyPath: "company_descriptive_date",
		},
		&requestflag.Flag[string]{
			Name:     "company-discretionary-data",
			Usage:    "Data associated with the transfer set by the sender.",
			BodyPath: "company_discretionary_data",
		},
		&requestflag.Flag[string]{
			Name:     "company-entry-description",
			Usage:    "The description of the transfer set by the sender.",
			BodyPath: "company_entry_description",
		},
		&requestflag.Flag[string]{
			Name:     "company-id",
			Usage:    "The sender's company ID.",
			BodyPath: "company_id",
		},
		&requestflag.Flag[string]{
			Name:     "company-name",
			Usage:    "The name of the sender.",
			BodyPath: "company_name",
		},
		&requestflag.Flag[string]{
			Name:     "receiver-id-number",
			Usage:    "The ID of the receiver of the transfer.",
			BodyPath: "receiver_id_number",
		},
		&requestflag.Flag[string]{
			Name:     "receiver-name",
			Usage:    "The name of the receiver of the transfer.",
			BodyPath: "receiver_name",
		},
		&requestflag.Flag[any]{
			Name:     "resolve-at",
			Usage:    "The time at which the transfer should be resolved. If not provided will resolve immediately.",
			BodyPath: "resolve_at",
		},
		&requestflag.Flag[string]{
			Name:     "standard-entry-class-code",
			Usage:    "The standard entry class code for the transfer.",
			BodyPath: "standard_entry_class_code",
		},
	},
	Action:          handleSimulationsInboundACHTransfersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"addenda": {
		&requestflag.InnerFlag[string]{
			Name:       "addenda.category",
			Usage:      "The type of addenda to simulate being sent with the transfer.",
			InnerField: "category",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "addenda.freeform",
			Usage:      "Unstructured `payment_related_information` passed through with the transfer.",
			InnerField: "freeform",
		},
	},
})

func handleSimulationsInboundACHTransfersCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationInboundACHTransferNewParams{}

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
	_, err = client.Simulations.InboundACHTransfers.New(ctx, params, options...)
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
		Title:          "simulations:inbound-ach-transfers create",
		Transform:      transform,
	})
}
