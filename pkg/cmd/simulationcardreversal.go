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

var simulationsCardReversalsCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates the reversal of an authorization by a card acquirer. An authorization\ncan be partially reversed multiple times, up until the total authorized amount.\nMarks the pending transaction as complete if the authorization is fully\nreversed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-payment-id",
			Usage:    "The identifier of the Card Payment to create a reversal on.",
			Required: true,
			BodyPath: "card_payment_id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The amount of the reversal in minor units in the card authorization's currency. This defaults to the authorization amount.",
			BodyPath: "amount",
		},
	},
	Action:          handleSimulationsCardReversalsCreate,
	HideHelpCommand: true,
}

func handleSimulationsCardReversalsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationCardReversalNewParams{}

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
	_, err = client.Simulations.CardReversals.New(ctx, params, options...)
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
		Title:          "simulations:card-reversals create",
		Transform:      transform,
	})
}
