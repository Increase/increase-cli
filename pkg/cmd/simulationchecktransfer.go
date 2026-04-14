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

var simulationsCheckTransfersMail = cli.Command{
	Name:    "mail",
	Usage:   "Simulates the mailing of a [Check Transfer](#check-transfers), which happens\nperiodically throughout the day in production but can be sped up in sandbox.\nThis transfer must first have a `status` of `pending_approval` or\n`pending_submission`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "check-transfer-id",
			Usage:    "The identifier of the Check Transfer you wish to mail.",
			Required: true,
		},
	},
	Action:          handleSimulationsCheckTransfersMail,
	HideHelpCommand: true,
}

func handleSimulationsCheckTransfersMail(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("check-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("check-transfer-id", unusedArgs[0])
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
	_, err = client.Simulations.CheckTransfers.Mail(ctx, cmd.Value("check-transfer-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "simulations:check-transfers mail", obj, format, explicitFormat, transform)
}
