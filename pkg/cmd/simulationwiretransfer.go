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

var simulationsWireTransfersReverse = cli.Command{
	Name:    "reverse",
	Usage:   "Simulates the reversal of a [Wire Transfer](#wire-transfers) by the Federal\nReserve due to error conditions. This will also create a\n[Transaction](#transaction) to account for the returned funds. This Wire\nTransfer must first have a `status` of `complete`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "wire-transfer-id",
			Usage:    "The identifier of the Wire Transfer you wish to reverse.",
			Required: true,
		},
	},
	Action:          handleSimulationsWireTransfersReverse,
	HideHelpCommand: true,
}

var simulationsWireTransfersSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Simulates the submission of a [Wire Transfer](#wire-transfers) to the Federal\nReserve. This transfer must first have a `status` of `pending_approval` or\n`pending_creating`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "wire-transfer-id",
			Usage:    "The identifier of the Wire Transfer you wish to submit.",
			Required: true,
		},
	},
	Action:          handleSimulationsWireTransfersSubmit,
	HideHelpCommand: true,
}

func handleSimulationsWireTransfersReverse(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("wire-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("wire-transfer-id", unusedArgs[0])
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
	_, err = client.Simulations.WireTransfers.Reverse(ctx, cmd.Value("wire-transfer-id").(string), options...)
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
		Title:          "simulations:wire-transfers reverse",
		Transform:      transform,
	})
}

func handleSimulationsWireTransfersSubmit(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("wire-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("wire-transfer-id", unusedArgs[0])
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
	_, err = client.Simulations.WireTransfers.Submit(ctx, cmd.Value("wire-transfer-id").(string), options...)
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
		Title:          "simulations:wire-transfers submit",
		Transform:      transform,
	})
}
