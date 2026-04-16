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

var simulationsPhysicalCardsCreate = cli.Command{
	Name:    "create",
	Usage:   "This endpoint allows you to simulate receiving a tracking update for a Physical\nCard, to simulate the progress of a shipment.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "physical-card-id",
			Usage:    "The Physical Card you would like to action.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "category",
			Usage:    "The type of tracking event.",
			Required: true,
			BodyPath: "category",
		},
		&requestflag.Flag[any]{
			Name:     "carrier-estimated-delivery-at",
			Usage:    "The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time when the carrier expects the card to be delivered.",
			BodyPath: "carrier_estimated_delivery_at",
		},
		&requestflag.Flag[string]{
			Name:     "city",
			Usage:    "The city where the event took place.",
			BodyPath: "city",
		},
		&requestflag.Flag[string]{
			Name:     "postal-code",
			Usage:    "The postal code where the event took place.",
			BodyPath: "postal_code",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    "The state where the event took place.",
			BodyPath: "state",
		},
	},
	Action:          handleSimulationsPhysicalCardsCreate,
	HideHelpCommand: true,
}

var simulationsPhysicalCardsAdvanceShipment = cli.Command{
	Name:    "advance-shipment",
	Usage:   "This endpoint allows you to simulate advancing the shipment status of a Physical\nCard, to simulate e.g., that a physical card was attempted shipped but then\nfailed delivery.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "physical-card-id",
			Usage:    "The Physical Card you would like to action.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "shipment-status",
			Usage:    "The shipment status to move the Physical Card to.",
			Required: true,
			BodyPath: "shipment_status",
		},
	},
	Action:          handleSimulationsPhysicalCardsAdvanceShipment,
	HideHelpCommand: true,
}

func handleSimulationsPhysicalCardsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("physical-card-id") && len(unusedArgs) > 0 {
		cmd.Set("physical-card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationPhysicalCardNewParams{}

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
	_, err = client.Simulations.PhysicalCards.New(
		ctx,
		cmd.Value("physical-card-id").(string),
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
		Title:          "simulations:physical-cards create",
		Transform:      transform,
	})
}

func handleSimulationsPhysicalCardsAdvanceShipment(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("physical-card-id") && len(unusedArgs) > 0 {
		cmd.Set("physical-card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationPhysicalCardAdvanceShipmentParams{}

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
	_, err = client.Simulations.PhysicalCards.AdvanceShipment(
		ctx,
		cmd.Value("physical-card-id").(string),
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
		Title:          "simulations:physical-cards advance-shipment",
		Transform:      transform,
	})
}
