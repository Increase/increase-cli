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

var simulationsCardFuelConfirmationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates the fuel confirmation of an authorization by a card acquirer. This\nhappens asynchronously right after a fuel pump transaction is completed. A fuel\nconfirmation can only happen once per authorization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The amount of the fuel_confirmation in minor units in the card authorization's currency.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "card-payment-id",
			Usage:    "The identifier of the Card Payment to create a fuel_confirmation on.",
			Required: true,
			BodyPath: "card_payment_id",
		},
	},
	Action:          handleSimulationsCardFuelConfirmationsCreate,
	HideHelpCommand: true,
}

func handleSimulationsCardFuelConfirmationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationCardFuelConfirmationNewParams{}

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
	_, err = client.Simulations.CardFuelConfirmations.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "simulations:card-fuel-confirmations create", obj, format, explicitFormat, transform)
}
