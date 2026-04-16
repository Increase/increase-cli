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

var realTimePaymentsTransfersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a Real-Time Payments Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The transfer amount in USD cents. For Real-Time Payments transfers, must be positive.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-name",
			Usage:    "The name of the transfer's recipient.",
			Required: true,
			BodyPath: "creditor_name",
		},
		&requestflag.Flag[string]{
			Name:     "source-account-number-id",
			Usage:    "The identifier of the Account Number from which to send the transfer.",
			Required: true,
			BodyPath: "source_account_number_id",
		},
		&requestflag.Flag[string]{
			Name:     "unstructured-remittance-information",
			Usage:    "Unstructured information that will show on the recipient's bank statement.",
			Required: true,
			BodyPath: "unstructured_remittance_information",
		},
		&requestflag.Flag[string]{
			Name:     "account-number",
			Usage:    "The destination account number.",
			BodyPath: "account_number",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-name",
			Usage:    "The name of the transfer's sender. If not provided, defaults to the name of the account's entity.",
			BodyPath: "debtor_name",
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
			Usage:    "The destination American Bankers' Association (ABA) Routing Transit Number (RTN).",
			BodyPath: "routing_number",
		},
		&requestflag.Flag[string]{
			Name:     "ultimate-creditor-name",
			Usage:    "The name of the ultimate recipient of the transfer. Set this if the creditor is an intermediary receiving the payment for someone else.",
			BodyPath: "ultimate_creditor_name",
		},
		&requestflag.Flag[string]{
			Name:     "ultimate-debtor-name",
			Usage:    "The name of the ultimate sender of the transfer. Set this if the funds are being sent on behalf of someone who is not the account holder at Increase.",
			BodyPath: "ultimate_debtor_name",
		},
	},
	Action:          handleRealTimePaymentsTransfersCreate,
	HideHelpCommand: true,
}

var realTimePaymentsTransfersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Real-Time Payments Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "real-time-payments-transfer-id",
			Usage:    "The identifier of the Real-Time Payments Transfer.",
			Required: true,
		},
	},
	Action:          handleRealTimePaymentsTransfersRetrieve,
	HideHelpCommand: true,
}

var realTimePaymentsTransfersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Real-Time Payments Transfers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter Real-Time Payments Transfers to those belonging to the specified Account.",
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
			Usage:     "Filter Real-Time Payments Transfers to those made to the specified External Account.",
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
	Action:          handleRealTimePaymentsTransfersList,
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

var realTimePaymentsTransfersApprove = cli.Command{
	Name:    "approve",
	Usage:   "Approves a Real-Time Payments Transfer in a pending_approval state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "real-time-payments-transfer-id",
			Usage:    "The identifier of the Real-Time Payments Transfer to approve.",
			Required: true,
		},
	},
	Action:          handleRealTimePaymentsTransfersApprove,
	HideHelpCommand: true,
}

var realTimePaymentsTransfersCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancels a Real-Time Payments Transfer in a pending_approval state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "real-time-payments-transfer-id",
			Usage:    "The identifier of the pending Real-Time Payments Transfer to cancel.",
			Required: true,
		},
	},
	Action:          handleRealTimePaymentsTransfersCancel,
	HideHelpCommand: true,
}

func handleRealTimePaymentsTransfersCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.RealTimePaymentsTransferNewParams{}

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
	_, err = client.RealTimePaymentsTransfers.New(ctx, params, options...)
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
		Title:          "real-time-payments-transfers create",
		Transform:      transform,
	})
}

func handleRealTimePaymentsTransfersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("real-time-payments-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("real-time-payments-transfer-id", unusedArgs[0])
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
	_, err = client.RealTimePaymentsTransfers.Get(ctx, cmd.Value("real-time-payments-transfer-id").(string), options...)
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
		Title:          "real-time-payments-transfers retrieve",
		Transform:      transform,
	})
}

func handleRealTimePaymentsTransfersList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.RealTimePaymentsTransferListParams{}

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
		_, err = client.RealTimePaymentsTransfers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "real-time-payments-transfers list",
			Transform:      transform,
		})
	} else {
		iter := client.RealTimePaymentsTransfers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "real-time-payments-transfers list",
			Transform:      transform,
		})
	}
}

func handleRealTimePaymentsTransfersApprove(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("real-time-payments-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("real-time-payments-transfer-id", unusedArgs[0])
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
	_, err = client.RealTimePaymentsTransfers.Approve(ctx, cmd.Value("real-time-payments-transfer-id").(string), options...)
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
		Title:          "real-time-payments-transfers approve",
		Transform:      transform,
	})
}

func handleRealTimePaymentsTransfersCancel(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("real-time-payments-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("real-time-payments-transfer-id", unusedArgs[0])
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
	_, err = client.RealTimePaymentsTransfers.Cancel(ctx, cmd.Value("real-time-payments-transfer-id").(string), options...)
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
		Title:          "real-time-payments-transfers cancel",
		Transform:      transform,
	})
}
