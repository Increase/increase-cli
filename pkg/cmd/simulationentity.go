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

var simulationsEntitiesValidation = requestflag.WithInnerFlags(cli.Command{
	Name:    "validation",
	Usage:   "Set the status for an\n[Entity's validation](/documentation/api/entities#entity-object.validation). In\nproduction, Know Your Customer validations\n[run automatically](/documentation/entity-validation#entity-validation). While\ndeveloping, it can be helpful to override the behavior in Sandbox.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-id",
			Usage:     "The identifier of the Entity whose validation status to update.",
			Required:  true,
			PathParam: "entity_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "issue",
			Usage:    "The validation issues to attach. Only allowed when `status` is `invalid`.",
			Required: true,
			BodyPath: "issues",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The validation status to set on the Entity.",
			Required: true,
			BodyPath: "status",
		},
	},
	Action:          handleSimulationsEntitiesValidation,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"issue": {
		&requestflag.InnerFlag[string]{
			Name:       "issue.category",
			Usage:      "The type of issue.",
			InnerField: "category",
		},
	},
})

func handleSimulationsEntitiesValidation(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-id") && len(unusedArgs) > 0 {
		cmd.Set("entity-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := increase.SimulationEntityValidationParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simulations.Entities.Validation(
		ctx,
		cmd.Value("entity-id").(string),
		params,
		options...,
	)
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "simulations:entities validation",
		Transform:      transform,
	})
}
