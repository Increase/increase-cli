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

var achTransfersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create an ACH Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "The Increase identifier for the account that will send the transfer.",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The transfer amount in USD cents. A positive amount originates a credit transfer pushing funds to the receiving account. A negative amount originates a debit transfer pulling funds from the receiving account.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "statement-descriptor",
			Usage:    "A description you choose to give the transfer. This will be saved with the transfer details, displayed in the dashboard, and returned by the API. If `individual_name` and `company_name` are not explicitly set by this API, the `statement_descriptor` will be sent in those fields to the receiving bank to help the customer recognize the transfer. You are highly encouraged to pass `individual_name` and `company_name` instead of relying on this fallback.",
			Required: true,
			BodyPath: "statement_descriptor",
		},
		&requestflag.Flag[string]{
			Name:     "account-number",
			Usage:    "The account number for the destination account.",
			BodyPath: "account_number",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "addenda",
			Usage:    "Additional information that will be sent to the recipient. This is included in the transfer data sent to the receiving bank.",
			BodyPath: "addenda",
		},
		&requestflag.Flag[string]{
			Name:     "company-descriptive-date",
			Usage:    "The description of the date of the transfer, usually in the format `YYMMDD`. This is included in the transfer data sent to the receiving bank.",
			BodyPath: "company_descriptive_date",
		},
		&requestflag.Flag[string]{
			Name:     "company-discretionary-data",
			Usage:    "The data you choose to associate with the transfer. This is included in the transfer data sent to the receiving bank.",
			BodyPath: "company_discretionary_data",
		},
		&requestflag.Flag[string]{
			Name:     "company-entry-description",
			Usage:    "A description of the transfer, included in the transfer data sent to the receiving bank. Standardized formatting may be required, for example `PAYROLL` for payroll-related Prearranged Payments and Deposits (PPD) credit transfers.",
			BodyPath: "company_entry_description",
		},
		&requestflag.Flag[string]{
			Name:     "company-name",
			Usage:    "The name by which the recipient knows you. This is included in the transfer data sent to the receiving bank.",
			BodyPath: "company_name",
		},
		&requestflag.Flag[string]{
			Name:     "destination-account-holder",
			Usage:    "The type of entity that owns the account to which the ACH Transfer is being sent.",
			BodyPath: "destination_account_holder",
		},
		&requestflag.Flag[string]{
			Name:     "external-account-id",
			Usage:    "The ID of an External Account to initiate a transfer to. If this parameter is provided, `account_number`, `routing_number`, and `funding` must be absent.",
			BodyPath: "external_account_id",
		},
		&requestflag.Flag[string]{
			Name:     "funding",
			Usage:    "The type of the account to which the transfer will be sent.",
			BodyPath: "funding",
		},
		&requestflag.Flag[string]{
			Name:     "individual-id",
			Usage:    "Your identifier for the transfer recipient.",
			BodyPath: "individual_id",
		},
		&requestflag.Flag[string]{
			Name:     "individual-name",
			Usage:    "The name of the transfer recipient. This value is informational and not verified by the recipient's bank.",
			BodyPath: "individual_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "preferred-effective-date",
			Usage:    "Configuration for how the effective date of the transfer will be set. This determines same-day vs future-dated settlement timing. If not set, defaults to a `settlement_schedule` of `same_day`. If set, exactly one of the child attributes must be set.",
			BodyPath: "preferred_effective_date",
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
			Name:     "standard-entry-class-code",
			Usage:    "The [Standard Entry Class (SEC) code](/documentation/ach-standard-entry-class-codes) to use for the transfer.",
			BodyPath: "standard_entry_class_code",
		},
		&requestflag.Flag[string]{
			Name:     "transaction-timing",
			Usage:    "The timing of the transaction.",
			BodyPath: "transaction_timing",
		},
	},
	Action:          handleACHTransfersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"addenda": {
		&requestflag.InnerFlag[string]{
			Name:       "addenda.category",
			Usage:      "The type of addenda to pass with the transfer.",
			InnerField: "category",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "addenda.freeform",
			Usage:      "Unstructured `payment_related_information` passed through with the transfer. Required if and only if `category` is `freeform`.",
			InnerField: "freeform",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "addenda.payment-order-remittance-advice",
			Usage:      "Structured ASC X12 820 remittance advice records. Please reach out to [support@increase.com](mailto:support@increase.com) for more information. Required if and only if `category` is `payment_order_remittance_advice`.",
			InnerField: "payment_order_remittance_advice",
		},
	},
	"preferred-effective-date": {
		&requestflag.InnerFlag[any]{
			Name:       "preferred-effective-date.date",
			Usage:      "A specific date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format to use as the effective date when submitting this transfer.",
			InnerField: "date",
		},
		&requestflag.InnerFlag[string]{
			Name:       "preferred-effective-date.settlement-schedule",
			Usage:      "A schedule by which Increase will choose an effective date for the transfer.",
			InnerField: "settlement_schedule",
		},
	},
})

var achTransfersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an ACH Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "ach-transfer-id",
			Usage:    "The identifier of the ACH Transfer.",
			Required: true,
		},
	},
	Action:          handleACHTransfersRetrieve,
	HideHelpCommand: true,
}

var achTransfersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List ACH Transfers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter ACH Transfers to those that originated from the specified Account.",
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
			Usage:     "Filter ACH Transfers to those made to the specified External Account.",
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
	Action:          handleACHTransfersList,
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

var achTransfersApprove = cli.Command{
	Name:    "approve",
	Usage:   "Approves an ACH Transfer in a pending_approval state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "ach-transfer-id",
			Usage:    "The identifier of the ACH Transfer to approve.",
			Required: true,
		},
	},
	Action:          handleACHTransfersApprove,
	HideHelpCommand: true,
}

var achTransfersCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancels an ACH Transfer in a pending_approval state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "ach-transfer-id",
			Usage:    "The identifier of the pending ACH Transfer to cancel.",
			Required: true,
		},
	},
	Action:          handleACHTransfersCancel,
	HideHelpCommand: true,
}

func handleACHTransfersCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.ACHTransferNewParams{}

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
	_, err = client.ACHTransfers.New(ctx, params, options...)
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
		Title:          "ach-transfers create",
		Transform:      transform,
	})
}

func handleACHTransfersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ach-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("ach-transfer-id", unusedArgs[0])
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
	_, err = client.ACHTransfers.Get(ctx, cmd.Value("ach-transfer-id").(string), options...)
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
		Title:          "ach-transfers retrieve",
		Transform:      transform,
	})
}

func handleACHTransfersList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.ACHTransferListParams{}

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
		_, err = client.ACHTransfers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "ach-transfers list",
			Transform:      transform,
		})
	} else {
		iter := client.ACHTransfers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "ach-transfers list",
			Transform:      transform,
		})
	}
}

func handleACHTransfersApprove(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ach-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("ach-transfer-id", unusedArgs[0])
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
	_, err = client.ACHTransfers.Approve(ctx, cmd.Value("ach-transfer-id").(string), options...)
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
		Title:          "ach-transfers approve",
		Transform:      transform,
	})
}

func handleACHTransfersCancel(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ach-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("ach-transfer-id", unusedArgs[0])
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
	_, err = client.ACHTransfers.Cancel(ctx, cmd.Value("ach-transfer-id").(string), options...)
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
		Title:          "ach-transfers cancel",
		Transform:      transform,
	})
}
