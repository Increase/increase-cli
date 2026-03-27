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

var simulationsProgramsCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates a [Program](#programs) being created in your group. By default, your\ngroup has one program called Commercial Banking. Note that when your group\noperates more than one program, `program_id` is a required field when creating\naccounts.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name of the program being added.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "bank",
			Usage:    "The bank for the program's accounts, defaults to First Internet Bank.",
			BodyPath: "bank",
		},
		&requestflag.Flag[int64]{
			Name:     "lending-maximum-extendable-credit",
			Usage:    "The maximum extendable credit of the program being added.",
			BodyPath: "lending_maximum_extendable_credit",
		},
		&requestflag.Flag[string]{
			Name:     "reserve-account-id",
			Usage:    "The identifier of the Account the Program should be added to is for.",
			BodyPath: "reserve_account_id",
		},
	},
	Action:          handleSimulationsProgramsCreate,
	HideHelpCommand: true,
}

func handleSimulationsProgramsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationProgramNewParams{}

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
	_, err = client.Simulations.Programs.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "simulations:programs create", obj, format, transform)
}
