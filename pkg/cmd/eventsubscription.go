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

var eventSubscriptionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create an Event Subscription",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The URL you'd like us to send webhooks to.",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[string]{
			Name:     "oauth-connection-id",
			Usage:    "If specified, this subscription will only receive webhooks for Events associated with the specified OAuth Connection.",
			BodyPath: "oauth_connection_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "selected-event-category",
			Usage:    "If specified, this subscription will only receive webhooks for Events with the specified `category`. If specifying a Real-Time Decision event category, only one Event Category can be specified for the Event Subscription.",
			BodyPath: "selected_event_categories",
		},
		&requestflag.Flag[string]{
			Name:     "shared-secret",
			Usage:    "The key that will be used to sign webhooks. If no value is passed, a random string will be used as default.",
			BodyPath: "shared_secret",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The status of the event subscription. Defaults to `active` if not specified.",
			BodyPath: "status",
		},
	},
	Action:          handleEventSubscriptionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"selected-event-category": {
		&requestflag.InnerFlag[string]{
			Name:       "selected-event-category.event-category",
			Usage:      "The category of the Event to subscribe to.",
			InnerField: "event_category",
		},
	},
})

var eventSubscriptionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an Event Subscription",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "event-subscription-id",
			Usage:     "The identifier of the Event Subscription.",
			Required:  true,
			PathParam: "event_subscription_id",
		},
	},
	Action:          handleEventSubscriptionsRetrieve,
	HideHelpCommand: true,
}

var eventSubscriptionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an Event Subscription",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "event-subscription-id",
			Usage:     "The identifier of the Event Subscription.",
			Required:  true,
			PathParam: "event_subscription_id",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The status to update the Event Subscription with.",
			BodyPath: "status",
		},
	},
	Action:          handleEventSubscriptionsUpdate,
	HideHelpCommand: true,
}

var eventSubscriptionsList = cli.Command{
	Name:    "list",
	Usage:   "List Event Subscriptions",
	Suggest: true,
	Flags: []cli.Flag{
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
	Action:          handleEventSubscriptionsList,
	HideHelpCommand: true,
}

func handleEventSubscriptionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := increase.EventSubscriptionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EventSubscriptions.New(ctx, params, options...)
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
		Title:          "event-subscriptions create",
		Transform:      transform,
	})
}

func handleEventSubscriptionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("event-subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("event-subscription-id", unusedArgs[0])
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
	_, err = client.EventSubscriptions.Get(ctx, cmd.Value("event-subscription-id").(string), options...)
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
		Title:          "event-subscriptions retrieve",
		Transform:      transform,
	})
}

func handleEventSubscriptionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("event-subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("event-subscription-id", unusedArgs[0])
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

	params := increase.EventSubscriptionUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EventSubscriptions.Update(
		ctx,
		cmd.Value("event-subscription-id").(string),
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
		Title:          "event-subscriptions update",
		Transform:      transform,
	})
}

func handleEventSubscriptionsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := increase.EventSubscriptionListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.EventSubscriptions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "event-subscriptions list",
			Transform:      transform,
		})
	} else {
		iter := client.EventSubscriptions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "event-subscriptions list",
			Transform:      transform,
		})
	}
}
