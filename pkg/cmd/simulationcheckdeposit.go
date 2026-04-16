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

var simulationsCheckDepositsAdjustment = cli.Command{
	Name:    "adjustment",
	Usage:   "Simulates the creation of a\n[Check Deposit Adjustment](#check-deposit-adjustments) on a\n[Check Deposit](#check-deposits). This Check Deposit must first have a `status`\nof `submitted`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "check-deposit-id",
			Usage:    "The identifier of the Check Deposit you wish to adjust.",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The adjustment amount in the minor unit of the Check Deposit's currency (e.g., cents). A negative amount means that the funds are being clawed back by the other bank and is a debit to your account. Defaults to the negative of the Check Deposit amount.",
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "The reason for the adjustment. Defaults to `non_conforming_item`, which is often used for a low quality image that the recipient wasn't able to handle.",
			BodyPath: "reason",
		},
	},
	Action:          handleSimulationsCheckDepositsAdjustment,
	HideHelpCommand: true,
}

var simulationsCheckDepositsReject = cli.Command{
	Name:    "reject",
	Usage:   "Simulates the rejection of a [Check Deposit](#check-deposits) by Increase due to\nfactors like poor image quality. This Check Deposit must first have a `status`\nof `pending`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "check-deposit-id",
			Usage:    "The identifier of the Check Deposit you wish to reject.",
			Required: true,
		},
	},
	Action:          handleSimulationsCheckDepositsReject,
	HideHelpCommand: true,
}

var simulationsCheckDepositsReturn = cli.Command{
	Name:    "return",
	Usage:   "Simulates the return of a [Check Deposit](#check-deposits). This Check Deposit\nmust first have a `status` of `submitted`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "check-deposit-id",
			Usage:    "The identifier of the Check Deposit you wish to return.",
			Required: true,
		},
	},
	Action:          handleSimulationsCheckDepositsReturn,
	HideHelpCommand: true,
}

var simulationsCheckDepositsSubmit = requestflag.WithInnerFlags(cli.Command{
	Name:    "submit",
	Usage:   "Simulates the submission of a [Check Deposit](#check-deposits) to the Federal\nReserve. This Check Deposit must first have a `status` of `pending`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "check-deposit-id",
			Usage:    "The identifier of the Check Deposit you wish to submit.",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "scan",
			Usage:    "If set, the simulation will use these values for the check's scanned MICR data.",
			BodyPath: "scan",
		},
	},
	Action:          handleSimulationsCheckDepositsSubmit,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"scan": {
		&requestflag.InnerFlag[string]{
			Name:       "scan.account-number",
			Usage:      "The account number to be returned in the check deposit's scan data.",
			InnerField: "account_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "scan.routing-number",
			Usage:      "The routing number to be returned in the check deposit's scan data.",
			InnerField: "routing_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "scan.auxiliary-on-us",
			Usage:      "The auxiliary on-us data to be returned in the check deposit's scan data.",
			InnerField: "auxiliary_on_us",
		},
	},
})

func handleSimulationsCheckDepositsAdjustment(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("check-deposit-id") && len(unusedArgs) > 0 {
		cmd.Set("check-deposit-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationCheckDepositAdjustmentParams{}

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
	_, err = client.Simulations.CheckDeposits.Adjustment(
		ctx,
		cmd.Value("check-deposit-id").(string),
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
		Title:          "simulations:check-deposits adjustment",
		Transform:      transform,
	})
}

func handleSimulationsCheckDepositsReject(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("check-deposit-id") && len(unusedArgs) > 0 {
		cmd.Set("check-deposit-id", unusedArgs[0])
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
	_, err = client.Simulations.CheckDeposits.Reject(ctx, cmd.Value("check-deposit-id").(string), options...)
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
		Title:          "simulations:check-deposits reject",
		Transform:      transform,
	})
}

func handleSimulationsCheckDepositsReturn(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("check-deposit-id") && len(unusedArgs) > 0 {
		cmd.Set("check-deposit-id", unusedArgs[0])
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
	_, err = client.Simulations.CheckDeposits.Return(ctx, cmd.Value("check-deposit-id").(string), options...)
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
		Title:          "simulations:check-deposits return",
		Transform:      transform,
	})
}

func handleSimulationsCheckDepositsSubmit(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("check-deposit-id") && len(unusedArgs) > 0 {
		cmd.Set("check-deposit-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationCheckDepositSubmitParams{}

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
	_, err = client.Simulations.CheckDeposits.Submit(
		ctx,
		cmd.Value("check-deposit-id").(string),
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
		Title:          "simulations:check-deposits submit",
		Transform:      transform,
	})
}
