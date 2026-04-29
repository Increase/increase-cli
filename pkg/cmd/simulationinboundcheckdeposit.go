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

var simulationsInboundCheckDepositsCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates an Inbound Check Deposit against your account. This imitates someone\ndepositing a check at their bank that was issued from your account. It may or\nmay not be associated with a Check Transfer. Increase will evaluate the Inbound\nCheck Deposit as we would in production and either create a Transaction or a\nDeclined Transaction as a result. You can inspect the resulting Inbound Check\nDeposit object to see the result.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-number-id",
			Usage:    "The identifier of the Account Number the Inbound Check Deposit will be against.",
			Required: true,
			BodyPath: "account_number_id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The check amount in cents.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "check-number",
			Usage:    "The check number on the check to be deposited.",
			Required: true,
			BodyPath: "check_number",
		},
		&requestflag.Flag[string]{
			Name:     "payee-name-analysis",
			Usage:    "Simulate the outcome of [payee name checking](https://increase.com/documentation/positive-pay#payee-name-mismatches). Defaults to `not_evaluated`.",
			BodyPath: "payee_name_analysis",
		},
	},
	Action:          handleSimulationsInboundCheckDepositsCreate,
	HideHelpCommand: true,
}

var simulationsInboundCheckDepositsAdjustment = cli.Command{
	Name:    "adjustment",
	Usage:   "Simulates an adjustment on an Inbound Check Deposit. The Inbound Check Deposit\nmust have a `status` of `accepted`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbound-check-deposit-id",
			Usage:     "The identifier of the Inbound Check Deposit to adjust.",
			Required:  true,
			PathParam: "inbound_check_deposit_id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The adjustment amount in cents. Defaults to the amount of the Inbound Check Deposit.",
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "The reason for the adjustment. Defaults to `wrong_payee_credit`.",
			BodyPath: "reason",
		},
	},
	Action:          handleSimulationsInboundCheckDepositsAdjustment,
	HideHelpCommand: true,
}

func handleSimulationsInboundCheckDepositsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := increase.SimulationInboundCheckDepositNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simulations.InboundCheckDeposits.New(ctx, params, options...)
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
		Title:          "simulations:inbound-check-deposits create",
		Transform:      transform,
	})
}

func handleSimulationsInboundCheckDepositsAdjustment(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbound-check-deposit-id") && len(unusedArgs) > 0 {
		cmd.Set("inbound-check-deposit-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := increase.SimulationInboundCheckDepositAdjustmentParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simulations.InboundCheckDeposits.Adjustment(
		ctx,
		cmd.Value("inbound-check-deposit-id").(string),
		params,
		options...,
	)
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
		Title:          "simulations:inbound-check-deposits adjustment",
		Transform:      transform,
	})
}
