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

var cardDisputesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Card Dispute",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "disputed-transaction-id",
			Usage:    "The Transaction you wish to dispute. This Transaction must have a `source_type` of `card_settlement`.",
			Required: true,
			BodyPath: "disputed_transaction_id",
		},
		&requestflag.Flag[string]{
			Name:     "network",
			Usage:    "The network of the disputed transaction. Details specific to the network are required under the sub-object with the same identifier as the network.",
			Required: true,
			BodyPath: "network",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The monetary amount of the part of the transaction that is being disputed. This is optional and will default to the full amount of the transaction if not provided. If provided, the amount must be less than or equal to the amount of the transaction.",
			BodyPath: "amount",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "attachment-file",
			Usage:    "The files to be attached to the initial dispute submission.",
			BodyPath: "attachment_files",
		},
		&requestflag.Flag[string]{
			Name:     "explanation",
			Usage:    "The free-form explanation provided to Increase to provide more context for the user submission. This field is not sent directly to the card networks.",
			BodyPath: "explanation",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "visa",
			Usage:    "The Visa-specific parameters for the dispute. Required if and only if `network` is `visa`.",
			BodyPath: "visa",
		},
	},
	Action:          handleCardDisputesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"attachment-file": {
		&requestflag.InnerFlag[string]{
			Name:       "attachment-file.file-id",
			Usage:      "The ID of the file to be attached. The file must have a `purpose` of `card_dispute_attachment`.",
			InnerField: "file_id",
		},
	},
	"visa": {
		&requestflag.InnerFlag[string]{
			Name:       "visa.category",
			Usage:      "Category.",
			InnerField: "category",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.authorization",
			Usage:      "Authorization. Required if and only if `category` is `authorization`.",
			InnerField: "authorization",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-canceled-merchandise",
			Usage:      "Canceled merchandise. Required if and only if `category` is `consumer_canceled_merchandise`.",
			InnerField: "consumer_canceled_merchandise",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-canceled-recurring-transaction",
			Usage:      "Canceled recurring transaction. Required if and only if `category` is `consumer_canceled_recurring_transaction`.",
			InnerField: "consumer_canceled_recurring_transaction",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-canceled-services",
			Usage:      "Canceled services. Required if and only if `category` is `consumer_canceled_services`.",
			InnerField: "consumer_canceled_services",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-counterfeit-merchandise",
			Usage:      "Counterfeit merchandise. Required if and only if `category` is `consumer_counterfeit_merchandise`.",
			InnerField: "consumer_counterfeit_merchandise",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-credit-not-processed",
			Usage:      "Credit not processed. Required if and only if `category` is `consumer_credit_not_processed`.",
			InnerField: "consumer_credit_not_processed",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-damaged-or-defective-merchandise",
			Usage:      "Damaged or defective merchandise. Required if and only if `category` is `consumer_damaged_or_defective_merchandise`.",
			InnerField: "consumer_damaged_or_defective_merchandise",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-merchandise-misrepresentation",
			Usage:      "Merchandise misrepresentation. Required if and only if `category` is `consumer_merchandise_misrepresentation`.",
			InnerField: "consumer_merchandise_misrepresentation",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-merchandise-not-as-described",
			Usage:      "Merchandise not as described. Required if and only if `category` is `consumer_merchandise_not_as_described`.",
			InnerField: "consumer_merchandise_not_as_described",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-merchandise-not-received",
			Usage:      "Merchandise not received. Required if and only if `category` is `consumer_merchandise_not_received`.",
			InnerField: "consumer_merchandise_not_received",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-non-receipt-of-cash",
			Usage:      "Non-receipt of cash. Required if and only if `category` is `consumer_non_receipt_of_cash`.",
			InnerField: "consumer_non_receipt_of_cash",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-original-credit-transaction-not-accepted",
			Usage:      "Original Credit Transaction (OCT) not accepted. Required if and only if `category` is `consumer_original_credit_transaction_not_accepted`.",
			InnerField: "consumer_original_credit_transaction_not_accepted",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-quality-merchandise",
			Usage:      "Merchandise quality issue. Required if and only if `category` is `consumer_quality_merchandise`.",
			InnerField: "consumer_quality_merchandise",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-quality-services",
			Usage:      "Services quality issue. Required if and only if `category` is `consumer_quality_services`.",
			InnerField: "consumer_quality_services",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-services-misrepresentation",
			Usage:      "Services misrepresentation. Required if and only if `category` is `consumer_services_misrepresentation`.",
			InnerField: "consumer_services_misrepresentation",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-services-not-as-described",
			Usage:      "Services not as described. Required if and only if `category` is `consumer_services_not_as_described`.",
			InnerField: "consumer_services_not_as_described",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.consumer-services-not-received",
			Usage:      "Services not received. Required if and only if `category` is `consumer_services_not_received`.",
			InnerField: "consumer_services_not_received",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.fraud",
			Usage:      "Fraud. Required if and only if `category` is `fraud`.",
			InnerField: "fraud",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.processing-error",
			Usage:      "Processing error. Required if and only if `category` is `processing_error`.",
			InnerField: "processing_error",
		},
	},
})

var cardDisputesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Card Dispute",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-dispute-id",
			Usage:    "The identifier of the Card Dispute.",
			Required: true,
		},
	},
	Action:          handleCardDisputesRetrieve,
	HideHelpCommand: true,
}

var cardDisputesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Card Disputes",
	Suggest: true,
	Flags: []cli.Flag{
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
	Action:          handleCardDisputesList,
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
			Usage:      "Filter Card Disputes for those with the specified status or statuses. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

var cardDisputesSubmitUserSubmission = requestflag.WithInnerFlags(cli.Command{
	Name:    "submit-user-submission",
	Usage:   "Submit a User Submission for a Card Dispute",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-dispute-id",
			Usage:    "The identifier of the Card Dispute to submit a user submission for.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "network",
			Usage:    "The network of the Card Dispute. Details specific to the network are required under the sub-object with the same identifier as the network.",
			Required: true,
			BodyPath: "network",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The adjusted monetary amount of the part of the transaction that is being disputed. This is optional and will default to the most recent amount provided. If provided, the amount must be less than or equal to the amount of the transaction.",
			BodyPath: "amount",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "attachment-file",
			Usage:    "The files to be attached to the user submission.",
			BodyPath: "attachment_files",
		},
		&requestflag.Flag[string]{
			Name:     "explanation",
			Usage:    "The free-form explanation provided to Increase to provide more context for the user submission. This field is not sent directly to the card networks.",
			BodyPath: "explanation",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "visa",
			Usage:    "The Visa-specific parameters for the dispute. Required if and only if `network` is `visa`.",
			BodyPath: "visa",
		},
	},
	Action:          handleCardDisputesSubmitUserSubmission,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"attachment-file": {
		&requestflag.InnerFlag[string]{
			Name:       "attachment-file.file-id",
			Usage:      "The ID of the file to be attached. The file must have a `purpose` of `card_dispute_attachment`.",
			InnerField: "file_id",
		},
	},
	"visa": {
		&requestflag.InnerFlag[string]{
			Name:       "visa.category",
			Usage:      "The category of the user submission. Details specific to the category are required under the sub-object with the same identifier as the category.",
			InnerField: "category",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.chargeback",
			Usage:      "The chargeback details for the user submission. Required if and only if `category` is `chargeback`.",
			InnerField: "chargeback",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.merchant-prearbitration-decline",
			Usage:      "The merchant pre-arbitration decline details for the user submission. Required if and only if `category` is `merchant_prearbitration_decline`.",
			InnerField: "merchant_prearbitration_decline",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "visa.user-prearbitration",
			Usage:      "The user pre-arbitration details for the user submission. Required if and only if `category` is `user_prearbitration`.",
			InnerField: "user_prearbitration",
		},
	},
})

var cardDisputesWithdraw = cli.Command{
	Name:    "withdraw",
	Usage:   "Withdraw a Card Dispute",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-dispute-id",
			Usage:    "The identifier of the Card Dispute to withdraw.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "explanation",
			Usage:    "The explanation for withdrawing the Card Dispute.",
			BodyPath: "explanation",
		},
	},
	Action:          handleCardDisputesWithdraw,
	HideHelpCommand: true,
}

func handleCardDisputesCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CardDisputeNewParams{}

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
	_, err = client.CardDisputes.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "card-disputes create", obj, format, transform)
}

func handleCardDisputesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-dispute-id") && len(unusedArgs) > 0 {
		cmd.Set("card-dispute-id", unusedArgs[0])
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
	_, err = client.CardDisputes.Get(ctx, cmd.Value("card-dispute-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "card-disputes retrieve", obj, format, transform)
}

func handleCardDisputesList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CardDisputeListParams{}

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
		_, err = client.CardDisputes.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "card-disputes list", obj, format, transform)
	} else {
		iter := client.CardDisputes.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "card-disputes list", iter, format, transform, maxItems)
	}
}

func handleCardDisputesSubmitUserSubmission(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-dispute-id") && len(unusedArgs) > 0 {
		cmd.Set("card-dispute-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CardDisputeSubmitUserSubmissionParams{}

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
	_, err = client.CardDisputes.SubmitUserSubmission(
		ctx,
		cmd.Value("card-dispute-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "card-disputes submit-user-submission", obj, format, transform)
}

func handleCardDisputesWithdraw(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-dispute-id") && len(unusedArgs) > 0 {
		cmd.Set("card-dispute-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CardDisputeWithdrawParams{}

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
	_, err = client.CardDisputes.Withdraw(
		ctx,
		cmd.Value("card-dispute-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "card-disputes withdraw", obj, format, transform)
}
