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

var physicalCardsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Physical Card",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-id",
			Usage:    "The underlying card representing this physical card.",
			Required: true,
			BodyPath: "card_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "cardholder",
			Usage:    "Details about the cardholder, as it will appear on the physical card.",
			Required: true,
			BodyPath: "cardholder",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "shipment",
			Usage:    "The details used to ship this physical card.",
			Required: true,
			BodyPath: "shipment",
		},
		&requestflag.Flag[string]{
			Name:     "physical-card-profile-id",
			Usage:    "The physical card profile to use for this physical card. The latest default physical card profile will be used if not provided.",
			BodyPath: "physical_card_profile_id",
		},
	},
	Action:          handlePhysicalCardsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"cardholder": {
		&requestflag.InnerFlag[string]{
			Name:       "cardholder.first-name",
			Usage:      "The cardholder's first name.",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "cardholder.last-name",
			Usage:      "The cardholder's last name.",
			InnerField: "last_name",
		},
	},
	"shipment": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "shipment.address",
			Usage:      "The address to where the card should be shipped.",
			InnerField: "address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shipment.method",
			Usage:      "The shipping method to use.",
			InnerField: "method",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shipment.schedule",
			Usage:      "When this physical card should be produced by the card printer. The default timeline is the day after the card printer receives the order, except for `FEDEX_PRIORITY_OVERNIGHT` cards, which default to `SAME_DAY`. To use faster production methods, please reach out to [support@increase.com](mailto:support@increase.com).",
			InnerField: "schedule",
		},
	},
})

var physicalCardsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Physical Card",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "physical-card-id",
			Usage:    "The identifier of the Physical Card.",
			Required: true,
		},
	},
	Action:          handlePhysicalCardsRetrieve,
	HideHelpCommand: true,
}

var physicalCardsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a Physical Card",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "physical-card-id",
			Usage:    "The Physical Card identifier.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The status to update the Physical Card to.",
			Required: true,
			BodyPath: "status",
		},
	},
	Action:          handlePhysicalCardsUpdate,
	HideHelpCommand: true,
}

var physicalCardsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Physical Cards",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "card-id",
			Usage:     "Filter Physical Cards to ones belonging to the specified Card.",
			QueryPath: "card_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "created-at",
			QueryPath: "created_at",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Return the page of entries after this one.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "idempotency-key",
			Usage:     "Filter records to the one with the specified `idempotency_key` you chose for that object. This value is unique across Increase and is used to ensure that a request is only processed once. Learn more about [idempotency](https://increase.com/documentation/idempotency-keys).",
			QueryPath: "idempotency_key",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Limit the size of the list that is returned. The default (and maximum) is 100 objects.",
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handlePhysicalCardsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"created-at": {
		&requestflag.InnerFlag[any]{
			Name:       "created-at.after",
			Usage:      "Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.",
			InnerField: "after",
		},
		&requestflag.InnerFlag[any]{
			Name:       "created-at.before",
			Usage:      "Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.",
			InnerField: "before",
		},
		&requestflag.InnerFlag[any]{
			Name:       "created-at.on-or-after",
			Usage:      "Return results on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.",
			InnerField: "on_or_after",
		},
		&requestflag.InnerFlag[any]{
			Name:       "created-at.on-or-before",
			Usage:      "Return results on or before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.",
			InnerField: "on_or_before",
		},
	},
})

func handlePhysicalCardsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.PhysicalCardNewParams{}

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
	_, err = client.PhysicalCards.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "physical-cards create", obj, format, transform)
}

func handlePhysicalCardsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("physical-card-id") && len(unusedArgs) > 0 {
		cmd.Set("physical-card-id", unusedArgs[0])
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
	_, err = client.PhysicalCards.Get(ctx, cmd.Value("physical-card-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "physical-cards retrieve", obj, format, transform)
}

func handlePhysicalCardsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("physical-card-id") && len(unusedArgs) > 0 {
		cmd.Set("physical-card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.PhysicalCardUpdateParams{}

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
	_, err = client.PhysicalCards.Update(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "physical-cards update", obj, format, transform)
}

func handlePhysicalCardsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.PhysicalCardListParams{}

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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.PhysicalCards.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "physical-cards list", obj, format, transform)
	} else {
		iter := client.PhysicalCards.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "physical-cards list", iter, format, transform, maxItems)
	}
}
