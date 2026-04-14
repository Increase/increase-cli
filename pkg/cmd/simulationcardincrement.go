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

var simulationsCardIncrementsCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates the increment of an authorization by a card acquirer. An authorization\ncan be incremented multiple times.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The amount of the increment in minor units in the card authorization's currency.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "card-payment-id",
			Usage:    "The identifier of the Card Payment to create an increment on.",
			Required: true,
			BodyPath: "card_payment_id",
		},
		&requestflag.Flag[string]{
			Name:     "event-subscription-id",
			Usage:    "The identifier of the Event Subscription to use. If provided, will override the default real time event subscription. Because you can only create one real time decision event subscription, you can use this field to route events to any specified event subscription for testing purposes.",
			BodyPath: "event_subscription_id",
		},
	},
	Action:          handleSimulationsCardIncrementsCreate,
	HideHelpCommand: true,
}

func handleSimulationsCardIncrementsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationCardIncrementNewParams{}

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
	_, err = client.Simulations.CardIncrements.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "simulations:card-increments create", obj, format, explicitFormat, transform)
}
