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

var swiftTransfersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Swift Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "The identifier for the account that will send the transfer.",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:     "account-number",
			Usage:    "The creditor's account number.",
			Required: true,
			BodyPath: "account_number",
		},
		&requestflag.Flag[string]{
			Name:     "bank-identification-code",
			Usage:    "The bank identification code (BIC) of the creditor. If it ends with the three-character branch code, this must be 11 characters long. Otherwise this must be 8 characters and the branch code will be assumed to be `XXX`.",
			Required: true,
			BodyPath: "bank_identification_code",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "creditor-address",
			Usage:    "The creditor's address.",
			Required: true,
			BodyPath: "creditor_address",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-name",
			Usage:    "The creditor's name.",
			Required: true,
			BodyPath: "creditor_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "debtor-address",
			Usage:    "The debtor's address.",
			Required: true,
			BodyPath: "debtor_address",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-name",
			Usage:    "The debtor's name.",
			Required: true,
			BodyPath: "debtor_name",
		},
		&requestflag.Flag[int64]{
			Name:     "instructed-amount",
			Usage:    "The amount, in minor units of `instructed_currency`, to send to the creditor.",
			Required: true,
			BodyPath: "instructed_amount",
		},
		&requestflag.Flag[string]{
			Name:     "instructed-currency",
			Usage:    "The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code of the instructed amount.",
			Required: true,
			BodyPath: "instructed_currency",
		},
		&requestflag.Flag[string]{
			Name:     "source-account-number-id",
			Usage:    "The Account Number to include in the transfer as the debtor's account number.",
			Required: true,
			BodyPath: "source_account_number_id",
		},
		&requestflag.Flag[string]{
			Name:     "unstructured-remittance-information",
			Usage:    "Unstructured remittance information to include in the transfer.",
			Required: true,
			BodyPath: "unstructured_remittance_information",
		},
		&requestflag.Flag[bool]{
			Name:     "require-approval",
			Usage:    "Whether the transfer requires explicit approval via the dashboard or API.",
			BodyPath: "require_approval",
		},
		&requestflag.Flag[string]{
			Name:     "routing-number",
			Usage:    "The creditor's bank account routing or transit number. Required in certain countries.",
			BodyPath: "routing_number",
		},
	},
	Action:          handleSwiftTransfersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"creditor-address": {
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.city",
			Usage:      "The city, district, town, or village of the address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.country",
			Usage:      "The two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code for the country of the address.",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.line1",
			Usage:      "The first line of the address. This is usually the street number and street.",
			InnerField: "line1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.line2",
			Usage:      "The second line of the address. This might be the floor or room number.",
			InnerField: "line2",
		},
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.postal-code",
			Usage:      "The ZIP or postal code of the address. Required in certain countries.",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.state",
			Usage:      "The state, province, or region of the address. Required in certain countries.",
			InnerField: "state",
		},
	},
	"debtor-address": {
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.city",
			Usage:      "The city, district, town, or village of the address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.country",
			Usage:      "The two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code for the country of the address.",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.line1",
			Usage:      "The first line of the address. This is usually the street number and street.",
			InnerField: "line1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.line2",
			Usage:      "The second line of the address. This might be the floor or room number.",
			InnerField: "line2",
		},
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.postal-code",
			Usage:      "The ZIP or postal code of the address. Required in certain countries.",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.state",
			Usage:      "The state, province, or region of the address. Required in certain countries.",
			InnerField: "state",
		},
	},
})

var swiftTransfersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Swift Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "swift-transfer-id",
			Usage:    "The identifier of the Swift Transfer.",
			Required: true,
		},
	},
	Action:          handleSwiftTransfersRetrieve,
	HideHelpCommand: true,
}

var swiftTransfersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Swift Transfers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter Swift Transfers to those that originated from the specified Account.",
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
	Action:          handleSwiftTransfersList,
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
			Usage:      "Return results whose value is in the provided list. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

var swiftTransfersApprove = cli.Command{
	Name:    "approve",
	Usage:   "Approve a Swift Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "swift-transfer-id",
			Usage:    "The identifier of the Swift Transfer to approve.",
			Required: true,
		},
	},
	Action:          handleSwiftTransfersApprove,
	HideHelpCommand: true,
}

var swiftTransfersCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a pending Swift Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "swift-transfer-id",
			Usage:    "The identifier of the pending Swift Transfer to cancel.",
			Required: true,
		},
	},
	Action:          handleSwiftTransfersCancel,
	HideHelpCommand: true,
}

func handleSwiftTransfersCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SwiftTransferNewParams{}

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
	_, err = client.SwiftTransfers.New(ctx, params, options...)
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
		Title:          "swift-transfers create",
		Transform:      transform,
	})
}

func handleSwiftTransfersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("swift-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("swift-transfer-id", unusedArgs[0])
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
	_, err = client.SwiftTransfers.Get(ctx, cmd.Value("swift-transfer-id").(string), options...)
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
		Title:          "swift-transfers retrieve",
		Transform:      transform,
	})
}

func handleSwiftTransfersList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SwiftTransferListParams{}

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
		_, err = client.SwiftTransfers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "swift-transfers list",
			Transform:      transform,
		})
	} else {
		iter := client.SwiftTransfers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "swift-transfers list",
			Transform:      transform,
		})
	}
}

func handleSwiftTransfersApprove(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("swift-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("swift-transfer-id", unusedArgs[0])
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
	_, err = client.SwiftTransfers.Approve(ctx, cmd.Value("swift-transfer-id").(string), options...)
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
		Title:          "swift-transfers approve",
		Transform:      transform,
	})
}

func handleSwiftTransfersCancel(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("swift-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("swift-transfer-id", unusedArgs[0])
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
	_, err = client.SwiftTransfers.Cancel(ctx, cmd.Value("swift-transfer-id").(string), options...)
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
		Title:          "swift-transfers cancel",
		Transform:      transform,
	})
}
