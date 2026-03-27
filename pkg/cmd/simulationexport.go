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

var simulationsExportsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Many exports are created by you via POST /exports or in the Dashboard. Some\nexports are created automatically by Increase. For example, tax documents are\npublished once a year. In sandbox, you can trigger the arrival of an export that\nwould normally only be created automatically via this simulation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "category",
			Usage:    "The type of Export to create.",
			Required: true,
			BodyPath: "category",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "form-1099-int",
			Usage:    "Options for the created export. Required if `category` is equal to `form_1099_int`.",
			BodyPath: "form_1099_int",
		},
	},
	Action:          handleSimulationsExportsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"form-1099-int": {
		&requestflag.InnerFlag[string]{
			Name:       "form-1099-int.account-id",
			Usage:      "The identifier of the Account the tax document is for.",
			InnerField: "account_id",
		},
	},
})

func handleSimulationsExportsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationExportNewParams{}

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
	_, err = client.Simulations.Exports.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "simulations:exports create", obj, format, transform)
}
