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

var simulationsAccountRevenuePaymentsCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates an account revenue payment to your account. In production, this\nhappens automatically on the first of each month.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "The identifier of the Account the Account Revenue Payment should be paid to.",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The account revenue amount in cents. Must be positive.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "accrued-on-account-id",
			Usage:    "The identifier of the Account the account revenue accrued on. Defaults to `account_id`.",
			BodyPath: "accrued_on_account_id",
		},
		&requestflag.Flag[any]{
			Name:     "period-end",
			Usage:    "The end of the account revenue period. If not provided, defaults to the current time.",
			BodyPath: "period_end",
		},
		&requestflag.Flag[any]{
			Name:     "period-start",
			Usage:    "The start of the account revenue period. If not provided, defaults to the current time.",
			BodyPath: "period_start",
		},
	},
	Action:          handleSimulationsAccountRevenuePaymentsCreate,
	HideHelpCommand: true,
}

func handleSimulationsAccountRevenuePaymentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := increase.SimulationAccountRevenuePaymentNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simulations.AccountRevenuePayments.New(ctx, params, options...)
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
		Title:          "simulations:account-revenue-payments create",
		Transform:      transform,
	})
}
