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

var simulationsEntityOnboardingSessionsSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Simulates the submission of an\n[Entity Onboarding Session](#entity-onboarding-sessions). This session must have\na `status` of `active`. After submission, the session will transition to\n`expired` and a new Entity will be created.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-onboarding-session-id",
			Usage:    "The identifier of the Entity Onboarding Session you wish to submit.",
			Required: true,
		},
	},
	Action:          handleSimulationsEntityOnboardingSessionsSubmit,
	HideHelpCommand: true,
}

func handleSimulationsEntityOnboardingSessionsSubmit(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-onboarding-session-id") && len(unusedArgs) > 0 {
		cmd.Set("entity-onboarding-session-id", unusedArgs[0])
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
	_, err = client.Simulations.EntityOnboardingSessions.Submit(ctx, cmd.Value("entity-onboarding-session-id").(string), options...)
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
		Title:          "simulations:entity-onboarding-sessions submit",
		Transform:      transform,
	})
}
