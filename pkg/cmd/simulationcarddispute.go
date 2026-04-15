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

var simulationsCardDisputesAction = requestflag.WithInnerFlags(cli.Command{
	Name:    "action",
	Usage:   "After a [Card Dispute](#card-disputes) is created in production, the dispute\nwill initially be in a `pending_user_submission_reviewing` state. Since no\nreview or further action happens in sandbox, this endpoint simulates moving a\nCard Dispute through its various states.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-dispute-id",
			Usage:    "The dispute you would like to action.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "network",
			Usage:    "The network of the Card Dispute. Details specific to the network are required under the sub-object with the same identifier as the network.",
			Required: true,
			BodyPath: "network",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "visa",
			Usage:    "The Visa-specific parameters for the taking action on the dispute. Required if and only if `network` is `visa`.",
			BodyPath: "visa",
		},
	},
	Action:          handleSimulationsCardDisputesAction,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"visa": {
		&requestflag.InnerFlag[string]{
			Name:       "visa.action",
			Usage:      "The action to take. Details specific to the action are required under the sub-object with the same identifier as the action.",
			InnerField: "action",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.accept-chargeback",
			Usage:      "The parameters for accepting the chargeback. Required if and only if `action` is `accept_chargeback`.",
			InnerField: "accept_chargeback",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.accept-user-submission",
			Usage:      "The parameters for accepting the user submission. Required if and only if `action` is `accept_user_submission`.",
			InnerField: "accept_user_submission",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.decline-user-prearbitration",
			Usage:      "The parameters for declining the prearbitration. Required if and only if `action` is `decline_user_prearbitration`.",
			InnerField: "decline_user_prearbitration",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.receive-merchant-prearbitration",
			Usage:      "The parameters for receiving the prearbitration. Required if and only if `action` is `receive_merchant_prearbitration`.",
			InnerField: "receive_merchant_prearbitration",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.represent",
			Usage:      "The parameters for re-presenting the dispute. Required if and only if `action` is `represent`.",
			InnerField: "represent",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.request-further-information",
			Usage:      "The parameters for requesting further information from the user. Required if and only if `action` is `request_further_information`.",
			InnerField: "request_further_information",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.time-out-chargeback",
			Usage:      "The parameters for timing out the chargeback. Required if and only if `action` is `time_out_chargeback`.",
			InnerField: "time_out_chargeback",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.time-out-merchant-prearbitration",
			Usage:      "The parameters for timing out the merchant prearbitration. Required if and only if `action` is `time_out_merchant_prearbitration`.",
			InnerField: "time_out_merchant_prearbitration",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.time-out-representment",
			Usage:      "The parameters for timing out the re-presentment. Required if and only if `action` is `time_out_representment`.",
			InnerField: "time_out_representment",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.time-out-user-prearbitration",
			Usage:      "The parameters for timing out the user prearbitration. Required if and only if `action` is `time_out_user_prearbitration`.",
			InnerField: "time_out_user_prearbitration",
		},
	},
})

func handleSimulationsCardDisputesAction(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-dispute-id") && len(unusedArgs) > 0 {
		cmd.Set("card-dispute-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationCardDisputeActionParams{}

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
	_, err = client.Simulations.CardDisputes.Action(
		ctx,
		cmd.Value("card-dispute-id").(string),
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
		Title:          "simulations:card-disputes action",
		Transform:      transform,
	})
}
