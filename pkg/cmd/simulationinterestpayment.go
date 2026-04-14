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

var simulationsInterestPaymentsCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates an interest payment to your account. In production, this happens\nautomatically on the first of each month.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "The identifier of the Account the Interest Payment should be paid to is for.",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The interest amount in cents. Must be positive.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "accrued-on-account-id",
			Usage:    "The identifier of the Account the Interest accrued on. Defaults to `account_id`.",
			BodyPath: "accrued_on_account_id",
		},
		&requestflag.Flag[any]{
			Name:     "period-end",
			Usage:    "The end of the interest period. If not provided, defaults to the current time.",
			BodyPath: "period_end",
		},
		&requestflag.Flag[any]{
			Name:     "period-start",
			Usage:    "The start of the interest period. If not provided, defaults to the current time.",
			BodyPath: "period_start",
		},
	},
	Action:          handleSimulationsInterestPaymentsCreate,
	HideHelpCommand: true,
}

func handleSimulationsInterestPaymentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationInterestPaymentNewParams{}

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
	_, err = client.Simulations.InterestPayments.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "simulations:interest-payments create", obj, format, explicitFormat, transform)
}
