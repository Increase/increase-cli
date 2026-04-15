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

var fednowTransfersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a FedNow Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The amount, in minor units, to send to the creditor.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-name",
			Usage:    "The creditor's name.",
			Required: true,
			BodyPath: "creditor_name",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-name",
			Usage:    "The debtor's name.",
			Required: true,
			BodyPath: "debtor_name",
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
		&requestflag.Flag[string]{
			Name:     "account-number",
			Usage:    "The creditor's account number.",
			BodyPath: "account_number",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "creditor-address",
			Usage:    "The creditor's address.",
			BodyPath: "creditor_address",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "debtor-address",
			Usage:    "The debtor's address.",
			BodyPath: "debtor_address",
		},
		&requestflag.Flag[string]{
			Name:     "external-account-id",
			Usage:    "The ID of an External Account to initiate a transfer to. If this parameter is provided, `account_number` and `routing_number` must be absent.",
			BodyPath: "external_account_id",
		},
		&requestflag.Flag[bool]{
			Name:     "require-approval",
			Usage:    "Whether the transfer requires explicit approval via the dashboard or API.",
			BodyPath: "require_approval",
		},
		&requestflag.Flag[string]{
			Name:     "routing-number",
			Usage:    "The creditor's bank account routing number.",
			BodyPath: "routing_number",
		},
	},
	Action:          handleFednowTransfersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"creditor-address": {
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.city",
			Usage:      "The city, district, town, or village of the address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.postal-code",
			Usage:      "The postal code component of the address.",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.state",
			Usage:      "The US state component of the address.",
			InnerField: "state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.line1",
			Usage:      "The first line of the address. This is usually the street number and street.",
			InnerField: "line1",
		},
	},
	"debtor-address": {
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.city",
			Usage:      "The city, district, town, or village of the address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.postal-code",
			Usage:      "The postal code component of the address.",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.state",
			Usage:      "The US state component of the address.",
			InnerField: "state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.line1",
			Usage:      "The first line of the address. This is usually the street number and street.",
			InnerField: "line1",
		},
	},
})

var fednowTransfersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a FedNow Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "fednow-transfer-id",
			Usage:    "The identifier of the FedNow Transfer.",
			Required: true,
		},
	},
	Action:          handleFednowTransfersRetrieve,
	HideHelpCommand: true,
}

var fednowTransfersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List FedNow Transfers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter FedNow Transfers to those that originated from the specified Account.",
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
			Name:      "external-account-id",
			Usage:     "Filter FedNow Transfers to those made to the specified External Account.",
			QueryPath: "external_account_id",
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
	Action:          handleFednowTransfersList,
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

var fednowTransfersApprove = cli.Command{
	Name:    "approve",
	Usage:   "Approve a FedNow Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "fednow-transfer-id",
			Usage:    "The identifier of the FedNow Transfer to approve.",
			Required: true,
		},
	},
	Action:          handleFednowTransfersApprove,
	HideHelpCommand: true,
}

var fednowTransfersCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a pending FedNow Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "fednow-transfer-id",
			Usage:    "The identifier of the pending FedNow Transfer to cancel.",
			Required: true,
		},
	},
	Action:          handleFednowTransfersCancel,
	HideHelpCommand: true,
}

func handleFednowTransfersCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.FednowTransferNewParams{}

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
	_, err = client.FednowTransfers.New(ctx, params, options...)
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
		Title:          "fednow-transfers create",
		Transform:      transform,
	})
}

func handleFednowTransfersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("fednow-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("fednow-transfer-id", unusedArgs[0])
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
	_, err = client.FednowTransfers.Get(ctx, cmd.Value("fednow-transfer-id").(string), options...)
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
		Title:          "fednow-transfers retrieve",
		Transform:      transform,
	})
}

func handleFednowTransfersList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.FednowTransferListParams{}

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
		_, err = client.FednowTransfers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "fednow-transfers list",
			Transform:      transform,
		})
	} else {
		iter := client.FednowTransfers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "fednow-transfers list",
			Transform:      transform,
		})
	}
}

func handleFednowTransfersApprove(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("fednow-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("fednow-transfer-id", unusedArgs[0])
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
	_, err = client.FednowTransfers.Approve(ctx, cmd.Value("fednow-transfer-id").(string), options...)
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
		Title:          "fednow-transfers approve",
		Transform:      transform,
	})
}

func handleFednowTransfersCancel(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("fednow-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("fednow-transfer-id", unusedArgs[0])
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
	_, err = client.FednowTransfers.Cancel(ctx, cmd.Value("fednow-transfer-id").(string), options...)
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
		Title:          "fednow-transfers cancel",
		Transform:      transform,
	})
}
