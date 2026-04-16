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

var checkTransfersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Check Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "The identifier for the account that will send the transfer.",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The transfer amount in USD cents.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "fulfillment-method",
			Usage:    "Whether Increase will print and mail the check or if you will do it yourself.",
			Required: true,
			BodyPath: "fulfillment_method",
		},
		&requestflag.Flag[string]{
			Name:     "source-account-number-id",
			Usage:    "The identifier of the Account Number from which to send the transfer and print on the check.",
			Required: true,
			BodyPath: "source_account_number_id",
		},
		&requestflag.Flag[string]{
			Name:     "balance-check",
			Usage:    "How the account's available balance should be checked. If omitted, the default behavior is `balance_check: full`.",
			BodyPath: "balance_check",
		},
		&requestflag.Flag[string]{
			Name:     "check-number",
			Usage:    "The check number Increase should use for the check. This should not contain leading zeroes and must be unique across the `source_account_number`. If this is omitted, Increase will generate a check number for you.",
			BodyPath: "check_number",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "physical-check",
			Usage:    "Details relating to the physical check that Increase will print and mail. This is required if `fulfillment_method` is equal to `physical_check`. It must not be included if any other `fulfillment_method` is provided.",
			BodyPath: "physical_check",
		},
		&requestflag.Flag[bool]{
			Name:     "require-approval",
			Usage:    "Whether the transfer requires explicit approval via the dashboard or API.",
			BodyPath: "require_approval",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "third-party",
			Usage:    "Details relating to the custom fulfillment you will perform. This is required if `fulfillment_method` is equal to `third_party`. It must not be included if any other `fulfillment_method` is provided.",
			BodyPath: "third_party",
		},
		&requestflag.Flag[any]{
			Name:     "valid-until-date",
			Usage:    "If provided, the check will be valid on or before this date. After this date, the check transfer will be automatically stopped and deposits will not be accepted. For checks printed by Increase, this date is included on the check as its expiry.",
			BodyPath: "valid_until_date",
		},
	},
	Action:          handleCheckTransfersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"physical-check": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "physical-check.mailing-address",
			Usage:      "Details for where Increase will mail the check.",
			InnerField: "mailing_address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "physical-check.memo",
			Usage:      "The descriptor that will be printed on the memo field on the check.",
			InnerField: "memo",
		},
		&requestflag.InnerFlag[string]{
			Name:       "physical-check.recipient-name",
			Usage:      "The name that will be printed on the check in the 'To:' field.",
			InnerField: "recipient_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "physical-check.attachment-file-id",
			Usage:      "The ID of a File to be attached to the check. This must have `purpose: check_attachment`. For details on pricing and restrictions, see https://increase.com/documentation/originating-checks#printing-checks .",
			InnerField: "attachment_file_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "physical-check.check-voucher-image-file-id",
			Usage:      "The ID of a File to be used as the check voucher image. This must have `purpose: check_voucher_image`. For details on pricing and restrictions, see https://increase.com/documentation/originating-checks#printing-checks .",
			InnerField: "check_voucher_image_file_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "physical-check.note",
			Usage:      "The descriptor that will be printed on the letter included with the check.",
			InnerField: "note",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "physical-check.payer",
			Usage:      "The payer of the check. This will be printed on the top-left portion of the check and defaults to the return address if unspecified. This should be an array of up to 4 elements, each of which represents a line of the payer.",
			InnerField: "payer",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "physical-check.return-address",
			Usage:      "The return address to be printed on the check. If omitted this will default to an Increase-owned address that will mark checks as delivery failed and shred them.",
			InnerField: "return_address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "physical-check.shipping-method",
			Usage:      "How to ship the check. For details on pricing, timing, and restrictions, see https://increase.com/documentation/originating-checks#printing-checks .",
			InnerField: "shipping_method",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "physical-check.signature",
			Usage:      "The signature that will appear on the check. If not provided, the check will be printed with 'No Signature Required'. At most one of `text` and `image_file_id` may be provided.",
			InnerField: "signature",
		},
	},
	"third-party": {
		&requestflag.InnerFlag[string]{
			Name:       "third-party.recipient-name",
			Usage:      "The pay-to name you will print on the check. If provided, this is used for [Positive Pay](/documentation/positive-pay). If this is omitted, Increase will be unable to validate the payer name when the check is deposited.",
			InnerField: "recipient_name",
		},
	},
})

var checkTransfersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Check Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "check-transfer-id",
			Usage:    "The identifier of the Check Transfer.",
			Required: true,
		},
	},
	Action:          handleCheckTransfersRetrieve,
	HideHelpCommand: true,
}

var checkTransfersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Check Transfers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter Check Transfers to those that originated from the specified Account.",
			QueryPath: "account_id",
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
		&requestflag.Flag[map[string]any]{
			Name:      "status",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleCheckTransfersList,
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
			Usage:      "Filter Check Transfers to those that have the specified status. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

var checkTransfersApprove = cli.Command{
	Name:    "approve",
	Usage:   "Approve a Check Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "check-transfer-id",
			Usage:    "The identifier of the Check Transfer to approve.",
			Required: true,
		},
	},
	Action:          handleCheckTransfersApprove,
	HideHelpCommand: true,
}

var checkTransfersCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a Check Transfer with the `pending_approval` status. See\n[Transfer Approvals](/documentation/transfer-approvals) for more information.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "check-transfer-id",
			Usage:    "The identifier of the pending Check Transfer to cancel.",
			Required: true,
		},
	},
	Action:          handleCheckTransfersCancel,
	HideHelpCommand: true,
}

var checkTransfersStopPayment = cli.Command{
	Name:    "stop-payment",
	Usage:   "Stop payment on a Check Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "check-transfer-id",
			Usage:    "The identifier of the Check Transfer.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "The reason why this transfer should be stopped.",
			BodyPath: "reason",
		},
	},
	Action:          handleCheckTransfersStopPayment,
	HideHelpCommand: true,
}

func handleCheckTransfersCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CheckTransferNewParams{}

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
	_, err = client.CheckTransfers.New(ctx, params, options...)
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
		Title:          "check-transfers create",
		Transform:      transform,
	})
}

func handleCheckTransfersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("check-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("check-transfer-id", unusedArgs[0])
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
	_, err = client.CheckTransfers.Get(ctx, cmd.Value("check-transfer-id").(string), options...)
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
		Title:          "check-transfers retrieve",
		Transform:      transform,
	})
}

func handleCheckTransfersList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CheckTransferListParams{}

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
		_, err = client.CheckTransfers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "check-transfers list",
			Transform:      transform,
		})
	} else {
		iter := client.CheckTransfers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "check-transfers list",
			Transform:      transform,
		})
	}
}

func handleCheckTransfersApprove(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("check-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("check-transfer-id", unusedArgs[0])
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
	_, err = client.CheckTransfers.Approve(ctx, cmd.Value("check-transfer-id").(string), options...)
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
		Title:          "check-transfers approve",
		Transform:      transform,
	})
}

func handleCheckTransfersCancel(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("check-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("check-transfer-id", unusedArgs[0])
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
	_, err = client.CheckTransfers.Cancel(ctx, cmd.Value("check-transfer-id").(string), options...)
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
		Title:          "check-transfers cancel",
		Transform:      transform,
	})
}

func handleCheckTransfersStopPayment(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("check-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("check-transfer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CheckTransferStopPaymentParams{}

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
	_, err = client.CheckTransfers.StopPayment(
		ctx,
		cmd.Value("check-transfer-id").(string),
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
		Title:          "check-transfers stop-payment",
		Transform:      transform,
	})
}
