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

var simulationsWireDrawdownRequestsRefuse = cli.Command{
	Name:    "refuse",
	Usage:   "Simulates a Wire Drawdown Request being refused by the debtor.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "wire-drawdown-request-id",
			Usage:    "The identifier of the Wire Drawdown Request you wish to refuse.",
			Required: true,
		},
	},
	Action:          handleSimulationsWireDrawdownRequestsRefuse,
	HideHelpCommand: true,
}

var simulationsWireDrawdownRequestsSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Simulates a Wire Drawdown Request being submitted to Fedwire.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "wire-drawdown-request-id",
			Usage:    "The identifier of the Wire Drawdown Request you wish to submit.",
			Required: true,
		},
	},
	Action:          handleSimulationsWireDrawdownRequestsSubmit,
	HideHelpCommand: true,
}

func handleSimulationsWireDrawdownRequestsRefuse(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("wire-drawdown-request-id") && len(unusedArgs) > 0 {
		cmd.Set("wire-drawdown-request-id", unusedArgs[0])
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
	_, err = client.Simulations.WireDrawdownRequests.Refuse(ctx, cmd.Value("wire-drawdown-request-id").(string), options...)
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
		Title:          "simulations:wire-drawdown-requests refuse",
		Transform:      transform,
	})
}

func handleSimulationsWireDrawdownRequestsSubmit(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("wire-drawdown-request-id") && len(unusedArgs) > 0 {
		cmd.Set("wire-drawdown-request-id", unusedArgs[0])
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
	_, err = client.Simulations.WireDrawdownRequests.Submit(ctx, cmd.Value("wire-drawdown-request-id").(string), options...)
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
		Title:          "simulations:wire-drawdown-requests submit",
		Transform:      transform,
	})
}
