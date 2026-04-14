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

var inboundACHTransfersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an Inbound ACH Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "inbound-ach-transfer-id",
			Usage:    "The identifier of the Inbound ACH Transfer to get details for.",
			Required: true,
		},
	},
	Action:          handleInboundACHTransfersRetrieve,
	HideHelpCommand: true,
}

var inboundACHTransfersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Inbound ACH Transfers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter Inbound ACH Transfers to ones belonging to the specified Account.",
			QueryPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "account-number-id",
			Usage:     "Filter Inbound ACH Transfers to ones belonging to the specified Account Number.",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleInboundACHTransfersList,
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
			Usage:      "Filter Inbound ACH Transfers to those with the specified status. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

var inboundACHTransfersCreateNotificationOfChange = cli.Command{
	Name:    "create-notification-of-change",
	Usage:   "Create a notification of change for an Inbound ACH Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "inbound-ach-transfer-id",
			Usage:    "The identifier of the Inbound ACH Transfer for which to create a notification of change.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "updated-account-number",
			Usage:    "The updated account number to send in the notification of change.",
			BodyPath: "updated_account_number",
		},
		&requestflag.Flag[string]{
			Name:     "updated-routing-number",
			Usage:    "The updated routing number to send in the notification of change.",
			BodyPath: "updated_routing_number",
		},
	},
	Action:          handleInboundACHTransfersCreateNotificationOfChange,
	HideHelpCommand: true,
}

var inboundACHTransfersDecline = cli.Command{
	Name:    "decline",
	Usage:   "Decline an Inbound ACH Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "inbound-ach-transfer-id",
			Usage:    "The identifier of the Inbound ACH Transfer to decline.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "The reason why this transfer will be returned. If this parameter is unset, the return codes will be `payment_stopped` for debits and `credit_entry_refused_by_receiver` for credits.",
			BodyPath: "reason",
		},
	},
	Action:          handleInboundACHTransfersDecline,
	HideHelpCommand: true,
}

var inboundACHTransfersTransferReturn = cli.Command{
	Name:    "transfer-return",
	Usage:   "Return an Inbound ACH Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "inbound-ach-transfer-id",
			Usage:    "The identifier of the Inbound ACH Transfer to return to the originating financial institution.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "The reason why this transfer will be returned. The most usual return codes are `payment_stopped` for debits and `credit_entry_refused_by_receiver` for credits.",
			Required: true,
			BodyPath: "reason",
		},
	},
	Action:          handleInboundACHTransfersTransferReturn,
	HideHelpCommand: true,
}

func handleInboundACHTransfersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbound-ach-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("inbound-ach-transfer-id", unusedArgs[0])
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
	_, err = client.InboundACHTransfers.Get(ctx, cmd.Value("inbound-ach-transfer-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "inbound-ach-transfers retrieve", obj, format, explicitFormat, transform)
}

func handleInboundACHTransfersList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.InboundACHTransferListParams{}

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
		_, err = client.InboundACHTransfers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "inbound-ach-transfers list", obj, format, explicitFormat, transform)
	} else {
		iter := client.InboundACHTransfers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "inbound-ach-transfers list", iter, format, explicitFormat, transform, maxItems)
	}
}

func handleInboundACHTransfersCreateNotificationOfChange(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbound-ach-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("inbound-ach-transfer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.InboundACHTransferNewNotificationOfChangeParams{}

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
	_, err = client.InboundACHTransfers.NewNotificationOfChange(
		ctx,
		cmd.Value("inbound-ach-transfer-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "inbound-ach-transfers create-notification-of-change", obj, format, explicitFormat, transform)
}

func handleInboundACHTransfersDecline(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbound-ach-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("inbound-ach-transfer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.InboundACHTransferDeclineParams{}

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
	_, err = client.InboundACHTransfers.Decline(
		ctx,
		cmd.Value("inbound-ach-transfer-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "inbound-ach-transfers decline", obj, format, explicitFormat, transform)
}

func handleInboundACHTransfersTransferReturn(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbound-ach-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("inbound-ach-transfer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.InboundACHTransferTransferReturnParams{}

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
	_, err = client.InboundACHTransfers.TransferReturn(
		ctx,
		cmd.Value("inbound-ach-transfer-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "inbound-ach-transfers transfer-return", obj, format, explicitFormat, transform)
}
