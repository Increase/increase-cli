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

var simulationsCardSettlementsCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates the settlement of an authorization by a card acquirer. After a card\nauthorization is created, the merchant will eventually send a settlement. This\nsimulates that event, which may occur many days after the purchase in\nproduction. The amount settled can be different from the amount originally\nauthorized, for example, when adding a tip to a restaurant bill.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-id",
			Usage:    "The identifier of the Card to create a settlement on.",
			Required: true,
			BodyPath: "card_id",
		},
		&requestflag.Flag[string]{
			Name:     "pending-transaction-id",
			Usage:    "The identifier of the Pending Transaction for the Card Authorization you wish to settle.",
			Required: true,
			BodyPath: "pending_transaction_id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The amount to be settled. This defaults to the amount of the Pending Transaction being settled.",
			BodyPath: "amount",
		},
	},
	Action:          handleSimulationsCardSettlementsCreate,
	HideHelpCommand: true,
}

func handleSimulationsCardSettlementsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationCardSettlementNewParams{}

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
	_, err = client.Simulations.CardSettlements.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "simulations:card-settlements create", obj, format, transform)
}
