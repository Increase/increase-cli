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

var simulationsCardRefundsCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates refunding a card transaction. The full value of the original sandbox\ntransaction is refunded.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The refund amount in cents. Pulled off the `pending_transaction` or the `transaction` if not provided.",
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "pending-transaction-id",
			Usage:    "The identifier of the Pending Transaction for the refund authorization. If this is provided, `transaction` must not be provided as a refund with a refund authorized can not be linked to a regular transaction.",
			BodyPath: "pending_transaction_id",
		},
		&requestflag.Flag[string]{
			Name:     "transaction-id",
			Usage:    "The identifier for the Transaction to refund. The Transaction's source must have a category of card_settlement.",
			BodyPath: "transaction_id",
		},
	},
	Action:          handleSimulationsCardRefundsCreate,
	HideHelpCommand: true,
}

func handleSimulationsCardRefundsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationCardRefundNewParams{}

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
	_, err = client.Simulations.CardRefunds.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "simulations:card-refunds create", obj, format, transform)
}
