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

var simulationsACHTransfersAcknowledge = cli.Command{
	Name:    "acknowledge",
	Usage:   "Simulates the acknowledgement of an [ACH Transfer](#ach-transfers) by the\nFederal Reserve. This transfer must first have a `status` of `submitted`. In\nproduction, the Federal Reserve generally acknowledges submitted ACH files\nwithin 30 minutes. Since sandbox ACH Transfers are not submitted to the Federal\nReserve, this endpoint allows you to skip that delay and add the acknowledgement\nsubresource to the ACH Transfer.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ach-transfer-id",
			Usage:     "The identifier of the ACH Transfer you wish to become acknowledged.",
			Required:  true,
			PathParam: "ach_transfer_id",
		},
	},
	Action:          handleSimulationsACHTransfersAcknowledge,
	HideHelpCommand: true,
}

var simulationsACHTransfersCreateNotificationOfChange = cli.Command{
	Name:    "create-notification-of-change",
	Usage:   "Simulates receiving a Notification of Change for an\n[ACH Transfer](#ach-transfers).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ach-transfer-id",
			Usage:     "The identifier of the ACH Transfer you wish to create a notification of change for.",
			Required:  true,
			PathParam: "ach_transfer_id",
		},
		&requestflag.Flag[string]{
			Name:     "change-code",
			Usage:    "The reason for the notification of change.",
			Required: true,
			BodyPath: "change_code",
		},
		&requestflag.Flag[string]{
			Name:     "corrected-data",
			Usage:    "The corrected data for the notification of change (e.g., a new routing number).",
			Required: true,
			BodyPath: "corrected_data",
		},
	},
	Action:          handleSimulationsACHTransfersCreateNotificationOfChange,
	HideHelpCommand: true,
}

var simulationsACHTransfersReturn = cli.Command{
	Name:    "return",
	Usage:   "Simulates the return of an [ACH Transfer](#ach-transfers) by the Federal Reserve\ndue to an error condition. This will also create a Transaction to account for\nthe returned funds. This transfer must first have a `status` of `submitted`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ach-transfer-id",
			Usage:     "The identifier of the ACH Transfer you wish to return.",
			Required:  true,
			PathParam: "ach_transfer_id",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "The reason why the Federal Reserve or destination bank returned this transfer. Defaults to `no_account`.",
			BodyPath: "reason",
		},
	},
	Action:          handleSimulationsACHTransfersReturn,
	HideHelpCommand: true,
}

var simulationsACHTransfersSettle = cli.Command{
	Name:    "settle",
	Usage:   "Simulates the settlement of an [ACH Transfer](#ach-transfers) by the Federal\nReserve. This transfer must first have a `status` of `pending_submission` or\n`submitted`. For convenience, if the transfer is in `status`:\n`pending_submission`, the simulation will also submit the transfer. Without this\nsimulation the transfer will eventually settle on its own following the same\nFederal Reserve timeline as in production. Additionally, you can specify the\nbehavior of the inbound funds hold that is created when the ACH Transfer is\nsettled. If no behavior is specified, the inbound funds hold will be released\nimmediately in order for the funds to be available for use.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ach-transfer-id",
			Usage:     "The identifier of the ACH Transfer you wish to become settled.",
			Required:  true,
			PathParam: "ach_transfer_id",
		},
		&requestflag.Flag[string]{
			Name:     "inbound-funds-hold-behavior",
			Usage:    "The behavior of the inbound funds hold that is created when the ACH Transfer is settled. If no behavior is specified, the inbound funds hold will be released immediately in order for the funds to be available for use.",
			BodyPath: "inbound_funds_hold_behavior",
		},
	},
	Action:          handleSimulationsACHTransfersSettle,
	HideHelpCommand: true,
}

var simulationsACHTransfersSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Simulates the submission of an [ACH Transfer](#ach-transfers) to the Federal\nReserve. This transfer must first have a `status` of `pending_approval` or\n`pending_submission`. In production, Increase submits ACH Transfers to the\nFederal Reserve three times per day on weekdays. Since sandbox ACH Transfers are\nnot submitted to the Federal Reserve, this endpoint allows you to skip that\ndelay and transition the ACH Transfer to a status of `submitted`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ach-transfer-id",
			Usage:     "The identifier of the ACH Transfer you wish to submit.",
			Required:  true,
			PathParam: "ach_transfer_id",
		},
	},
	Action:          handleSimulationsACHTransfersSubmit,
	HideHelpCommand: true,
}

func handleSimulationsACHTransfersAcknowledge(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ach-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("ach-transfer-id", unusedArgs[0])
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
	_, err = client.Simulations.ACHTransfers.Acknowledge(ctx, cmd.Value("ach-transfer-id").(string), options...)
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
		Title:          "simulations:ach-transfers acknowledge",
		Transform:      transform,
	})
}

func handleSimulationsACHTransfersCreateNotificationOfChange(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ach-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("ach-transfer-id", unusedArgs[0])
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

	params := increase.SimulationACHTransferNewNotificationOfChangeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simulations.ACHTransfers.NewNotificationOfChange(
		ctx,
		cmd.Value("ach-transfer-id").(string),
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
		Title:          "simulations:ach-transfers create-notification-of-change",
		Transform:      transform,
	})
}

func handleSimulationsACHTransfersReturn(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ach-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("ach-transfer-id", unusedArgs[0])
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

	params := increase.SimulationACHTransferReturnParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simulations.ACHTransfers.Return(
		ctx,
		cmd.Value("ach-transfer-id").(string),
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
		Title:          "simulations:ach-transfers return",
		Transform:      transform,
	})
}

func handleSimulationsACHTransfersSettle(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ach-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("ach-transfer-id", unusedArgs[0])
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

	params := increase.SimulationACHTransferSettleParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simulations.ACHTransfers.Settle(
		ctx,
		cmd.Value("ach-transfer-id").(string),
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
		Title:          "simulations:ach-transfers settle",
		Transform:      transform,
	})
}

func handleSimulationsACHTransfersSubmit(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ach-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("ach-transfer-id", unusedArgs[0])
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
	_, err = client.Simulations.ACHTransfers.Submit(ctx, cmd.Value("ach-transfer-id").(string), options...)
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
		Title:          "simulations:ach-transfers submit",
		Transform:      transform,
	})
}
