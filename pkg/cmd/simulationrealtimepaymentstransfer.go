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

var simulationsRealTimePaymentsTransfersComplete = requestflag.WithInnerFlags(cli.Command{
	Name:    "complete",
	Usage:   "Simulates submission of a\n[Real-Time Payments Transfer](#real-time-payments-transfers) and handling the\nresponse from the destination financial institution. This transfer must first\nhave a `status` of `pending_submission`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "real-time-payments-transfer-id",
			Usage:    "The identifier of the Real-Time Payments Transfer you wish to complete.",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "rejection",
			Usage:    "If set, the simulation will reject the transfer.",
			BodyPath: "rejection",
		},
	},
	Action:          handleSimulationsRealTimePaymentsTransfersComplete,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"rejection": {
		&requestflag.InnerFlag[string]{
			Name:       "rejection.reject-reason-code",
			Usage:      "The reason code that the simulated rejection will have.",
			InnerField: "reject_reason_code",
		},
	},
})

func handleSimulationsRealTimePaymentsTransfersComplete(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("real-time-payments-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("real-time-payments-transfer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationRealTimePaymentsTransferCompleteParams{}

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
	_, err = client.Simulations.RealTimePaymentsTransfers.Complete(
		ctx,
		cmd.Value("real-time-payments-transfer-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "simulations:real-time-payments-transfers complete", obj, format, transform)
}
