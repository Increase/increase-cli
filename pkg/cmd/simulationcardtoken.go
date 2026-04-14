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

var simulationsCardTokensCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Simulates tokenizing a card in the sandbox environment.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "capability",
			Usage:    "The capabilities of the outbound card token.",
			BodyPath: "capabilities",
		},
		&requestflag.Flag[any]{
			Name:     "expiration",
			Usage:    "The expiration date of the card.",
			BodyPath: "expiration",
		},
		&requestflag.Flag[string]{
			Name:     "last4",
			Usage:    "The last 4 digits of the card number.",
			BodyPath: "last4",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "outcome",
			Usage:    "The outcome to simulate for card push transfers using this token.",
			BodyPath: "outcome",
		},
		&requestflag.Flag[string]{
			Name:     "prefix",
			Usage:    "The prefix of the card number, usually the first 8 digits.",
			BodyPath: "prefix",
		},
		&requestflag.Flag[int64]{
			Name:     "primary-account-number-length",
			Usage:    "The total length of the card number, including prefix and last4.",
			BodyPath: "primary_account_number_length",
		},
	},
	Action:          handleSimulationsCardTokensCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"capability": {
		&requestflag.InnerFlag[string]{
			Name:       "capability.cross-border-push-transfers",
			Usage:      "The cross-border push transfers capability.",
			InnerField: "cross_border_push_transfers",
		},
		&requestflag.InnerFlag[string]{
			Name:       "capability.domestic-push-transfers",
			Usage:      "The domestic push transfers capability.",
			InnerField: "domestic_push_transfers",
		},
		&requestflag.InnerFlag[string]{
			Name:       "capability.route",
			Usage:      "The route of the capability.",
			InnerField: "route",
		},
	},
	"outcome": {
		&requestflag.InnerFlag[string]{
			Name:       "outcome.result",
			Usage:      "Whether card push transfers or validations will be approved or declined.",
			InnerField: "result",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "outcome.decline",
			Usage:      "If the result is declined, the details of the decline.",
			InnerField: "decline",
		},
	},
})

func handleSimulationsCardTokensCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationCardTokenNewParams{}

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
	_, err = client.Simulations.CardTokens.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "simulations:card-tokens create", obj, format, explicitFormat, transform)
}
