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

var simulationsAccountTransfersComplete = cli.Command{
	Name:    "complete",
	Usage:   "If your account is configured to require approval for each transfer, this\nendpoint simulates the approval of an [Account Transfer](#account-transfers).\nYou can also approve sandbox Account Transfers in the dashboard. This transfer\nmust first have a `status` of `pending_approval`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-transfer-id",
			Usage:    "The identifier of the Account Transfer you wish to complete.",
			Required: true,
		},
	},
	Action:          handleSimulationsAccountTransfersComplete,
	HideHelpCommand: true,
}

func handleSimulationsAccountTransfersComplete(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("account-transfer-id", unusedArgs[0])
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
	_, err = client.Simulations.AccountTransfers.Complete(ctx, cmd.Value("account-transfer-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "simulations:account-transfers complete", obj, format, transform)
}
