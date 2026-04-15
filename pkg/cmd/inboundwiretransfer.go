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

var inboundWireTransfersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an Inbound Wire Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "inbound-wire-transfer-id",
			Usage:    "The identifier of the Inbound Wire Transfer to get details for.",
			Required: true,
		},
	},
	Action:          handleInboundWireTransfersRetrieve,
	HideHelpCommand: true,
}

var inboundWireTransfersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Inbound Wire Transfers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter Inbound Wire Transfers to ones belonging to the specified Account.",
			QueryPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "account-number-id",
			Usage:     "Filter Inbound Wire Transfers to ones belonging to the specified Account Number.",
			QueryPath: "account_number_id",
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
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Limit the size of the list that is returned. The default (and maximum) is 100 objects.",
			QueryPath: "limit",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "status",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "wire-drawdown-request-id",
			Usage:     "Filter Inbound Wire Transfers to ones belonging to the specified Wire Drawdown Request.",
			QueryPath: "wire_drawdown_request_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleInboundWireTransfersList,
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
	"status": {
		&requestflag.InnerFlag[[]string]{
			Name:       "status.in",
			Usage:      "Filter Inbound Wire Transfers to those with the specified status. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

var inboundWireTransfersReverse = cli.Command{
	Name:    "reverse",
	Usage:   "Reverse an Inbound Wire Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "inbound-wire-transfer-id",
			Usage:    "The identifier of the Inbound Wire Transfer to reverse.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "Reason for the reversal.",
			Required: true,
			BodyPath: "reason",
		},
	},
	Action:          handleInboundWireTransfersReverse,
	HideHelpCommand: true,
}

func handleInboundWireTransfersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbound-wire-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("inbound-wire-transfer-id", unusedArgs[0])
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
	_, err = client.InboundWireTransfers.Get(ctx, cmd.Value("inbound-wire-transfer-id").(string), options...)
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
		Title:          "inbound-wire-transfers retrieve",
		Transform:      transform,
	})
}

func handleInboundWireTransfersList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.InboundWireTransferListParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.InboundWireTransfers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "inbound-wire-transfers list",
			Transform:      transform,
		})
	} else {
		iter := client.InboundWireTransfers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "inbound-wire-transfers list",
			Transform:      transform,
		})
	}
}

func handleInboundWireTransfersReverse(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbound-wire-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("inbound-wire-transfer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.InboundWireTransferReverseParams{}

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
	_, err = client.InboundWireTransfers.Reverse(
		ctx,
		cmd.Value("inbound-wire-transfer-id").(string),
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
		Title:          "inbound-wire-transfers reverse",
		Transform:      transform,
	})
}
