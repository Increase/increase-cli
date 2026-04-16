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

var wireTransfersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Wire Transfer",
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
		&requestflag.Flag[map[string]any]{
			Name:     "creditor",
			Usage:    "The person or business that is receiving the funds from the transfer.",
			Required: true,
			BodyPath: "creditor",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "remittance",
			Usage:    "Additional remittance information related to the wire transfer.",
			Required: true,
			BodyPath: "remittance",
		},
		&requestflag.Flag[string]{
			Name:     "account-number",
			Usage:    "The account number for the destination account.",
			BodyPath: "account_number",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "debtor",
			Usage:    "The person or business whose funds are being transferred. This is only necessary if you're transferring from a commingled account. Otherwise, we'll use the associated entity's details.",
			BodyPath: "debtor",
		},
		&requestflag.Flag[string]{
			Name:     "external-account-id",
			Usage:    "The ID of an External Account to initiate a transfer to. If this parameter is provided, `account_number` and `routing_number` must be absent.",
			BodyPath: "external_account_id",
		},
		&requestflag.Flag[string]{
			Name:     "inbound-wire-drawdown-request-id",
			Usage:    "The ID of an Inbound Wire Drawdown Request in response to which this transfer is being sent.",
			BodyPath: "inbound_wire_drawdown_request_id",
		},
		&requestflag.Flag[bool]{
			Name:     "require-approval",
			Usage:    "Whether the transfer requires explicit approval via the dashboard or API.",
			BodyPath: "require_approval",
		},
		&requestflag.Flag[string]{
			Name:     "routing-number",
			Usage:    "The American Bankers' Association (ABA) Routing Transit Number (RTN) for the destination account.",
			BodyPath: "routing_number",
		},
		&requestflag.Flag[string]{
			Name:     "source-account-number-id",
			Usage:    "The ID of an Account Number that will be passed to the wire's recipient",
			BodyPath: "source_account_number_id",
		},
	},
	Action:          handleWireTransfersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"creditor": {
		&requestflag.InnerFlag[string]{
			Name:       "creditor.name",
			Usage:      "The person or business's name.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "creditor.address",
			Usage:      "The person or business's address.",
			InnerField: "address",
		},
	},
	"remittance": {
		&requestflag.InnerFlag[string]{
			Name:       "remittance.category",
			Usage:      "The type of remittance information being passed.",
			InnerField: "category",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "remittance.tax",
			Usage:      "Internal Revenue Service (IRS) tax repayment information. Required if `category` is equal to `tax`.",
			InnerField: "tax",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "remittance.unstructured",
			Usage:      "Unstructured remittance information. Required if `category` is equal to `unstructured`.",
			InnerField: "unstructured",
		},
	},
	"debtor": {
		&requestflag.InnerFlag[string]{
			Name:       "debtor.name",
			Usage:      "The person or business's name.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "debtor.address",
			Usage:      "The person or business's address.",
			InnerField: "address",
		},
	},
})

var wireTransfersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Wire Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "wire-transfer-id",
			Usage:    "The identifier of the Wire Transfer.",
			Required: true,
		},
	},
	Action:          handleWireTransfersRetrieve,
	HideHelpCommand: true,
}

var wireTransfersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Wire Transfers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter Wire Transfers to those belonging to the specified Account.",
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
			Usage:     "Filter Wire Transfers to those made to the specified External Account.",
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
	Action:          handleWireTransfersList,
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

var wireTransfersApprove = cli.Command{
	Name:    "approve",
	Usage:   "Approve a Wire Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "wire-transfer-id",
			Usage:    "The identifier of the Wire Transfer to approve.",
			Required: true,
		},
	},
	Action:          handleWireTransfersApprove,
	HideHelpCommand: true,
}

var wireTransfersCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a pending Wire Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "wire-transfer-id",
			Usage:    "The identifier of the pending Wire Transfer to cancel.",
			Required: true,
		},
	},
	Action:          handleWireTransfersCancel,
	HideHelpCommand: true,
}

func handleWireTransfersCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.WireTransferNewParams{}

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
	_, err = client.WireTransfers.New(ctx, params, options...)
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
		Title:          "wire-transfers create",
		Transform:      transform,
	})
}

func handleWireTransfersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("wire-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("wire-transfer-id", unusedArgs[0])
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
	_, err = client.WireTransfers.Get(ctx, cmd.Value("wire-transfer-id").(string), options...)
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
		Title:          "wire-transfers retrieve",
		Transform:      transform,
	})
}

func handleWireTransfersList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.WireTransferListParams{}

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
		_, err = client.WireTransfers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "wire-transfers list",
			Transform:      transform,
		})
	} else {
		iter := client.WireTransfers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "wire-transfers list",
			Transform:      transform,
		})
	}
}

func handleWireTransfersApprove(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("wire-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("wire-transfer-id", unusedArgs[0])
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
	_, err = client.WireTransfers.Approve(ctx, cmd.Value("wire-transfer-id").(string), options...)
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
		Title:          "wire-transfers approve",
		Transform:      transform,
	})
}

func handleWireTransfersCancel(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("wire-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("wire-transfer-id", unusedArgs[0])
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
	_, err = client.WireTransfers.Cancel(ctx, cmd.Value("wire-transfer-id").(string), options...)
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
		Title:          "wire-transfers cancel",
		Transform:      transform,
	})
}
