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

var cardPushTransfersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Card Push Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "business-application-identifier",
			Usage:    "The Business Application Identifier describes the type of transaction being performed. Your program must be approved for the specified Business Application Identifier in order to use it.",
			Required: true,
			BodyPath: "business_application_identifier",
		},
		&requestflag.Flag[string]{
			Name:     "card-token-id",
			Usage:    "The Increase identifier for the Card Token that represents the card number you're pushing funds to.",
			Required: true,
			BodyPath: "card_token_id",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-category-code",
			Usage:    "The merchant category code (MCC) of the merchant (generally your business) sending the transfer. This is a four-digit code that describes the type of business or service provided by the merchant. Your program must be approved for the specified MCC in order to use it.",
			Required: true,
			BodyPath: "merchant_category_code",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-city-name",
			Usage:    "The city name of the merchant (generally your business) sending the transfer.",
			Required: true,
			BodyPath: "merchant_city_name",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-name",
			Usage:    "The merchant name shows up as the statement descriptor for the transfer. This is typically the name of your business or organization.",
			Required: true,
			BodyPath: "merchant_name",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-name-prefix",
			Usage:    "For certain Business Application Identifiers, the statement descriptor is `merchant_name_prefix*sender_name`, where the `merchant_name_prefix` is a one to four character prefix that identifies the merchant.",
			Required: true,
			BodyPath: "merchant_name_prefix",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-postal-code",
			Usage:    "The postal code of the merchant (generally your business) sending the transfer.",
			Required: true,
			BodyPath: "merchant_postal_code",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-state",
			Usage:    "The state of the merchant (generally your business) sending the transfer.",
			Required: true,
			BodyPath: "merchant_state",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "presentment-amount",
			Usage:    "The amount to transfer. The receiving bank will convert this to the cardholder's currency. The amount that is applied to your Increase account matches the currency of your account.",
			Required: true,
			BodyPath: "presentment_amount",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-name",
			Usage:    "The name of the funds recipient.",
			Required: true,
			BodyPath: "recipient_name",
		},
		&requestflag.Flag[string]{
			Name:     "sender-address-city",
			Usage:    "The city of the sender.",
			Required: true,
			BodyPath: "sender_address_city",
		},
		&requestflag.Flag[string]{
			Name:     "sender-address-line1",
			Usage:    "The address line 1 of the sender.",
			Required: true,
			BodyPath: "sender_address_line1",
		},
		&requestflag.Flag[string]{
			Name:     "sender-address-postal-code",
			Usage:    "The postal code of the sender.",
			Required: true,
			BodyPath: "sender_address_postal_code",
		},
		&requestflag.Flag[string]{
			Name:     "sender-address-state",
			Usage:    "The state of the sender.",
			Required: true,
			BodyPath: "sender_address_state",
		},
		&requestflag.Flag[string]{
			Name:     "sender-name",
			Usage:    "The name of the funds originator.",
			Required: true,
			BodyPath: "sender_name",
		},
		&requestflag.Flag[string]{
			Name:     "source-account-number-id",
			Usage:    "The identifier of the Account Number from which to send the transfer.",
			Required: true,
			BodyPath: "source_account_number_id",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-legal-business-name",
			Usage:    "The legal business name of the merchant (generally your business) sending the transfer. Required if the card is issued in Canada.",
			BodyPath: "merchant_legal_business_name",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-street-address",
			Usage:    "The street address of the merchant (generally your business) sending the transfer. Required if the card is issued in Canada.",
			BodyPath: "merchant_street_address",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-address-city",
			Usage:    "The city of the recipient. Required if the card is issued in Canada.",
			BodyPath: "recipient_address_city",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-address-line1",
			Usage:    "The first line of the recipient's address. Required if the card is issued in Canada.",
			BodyPath: "recipient_address_line1",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-address-postal-code",
			Usage:    "The postal code of the recipient. Required if the card is issued in Canada.",
			BodyPath: "recipient_address_postal_code",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-address-state",
			Usage:    "The state or province of the recipient. Required if the card is issued in Canada.",
			BodyPath: "recipient_address_state",
		},
		&requestflag.Flag[bool]{
			Name:     "require-approval",
			Usage:    "Whether the transfer requires explicit approval via the dashboard or API.",
			BodyPath: "require_approval",
		},
	},
	Action:          handleCardPushTransfersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"presentment-amount": {
		&requestflag.InnerFlag[string]{
			Name:       "presentment-amount.currency",
			Usage:      "The ISO 4217 currency code representing the currency of the amount.",
			InnerField: "currency",
		},
		&requestflag.InnerFlag[string]{
			Name:       "presentment-amount.value",
			Usage:      "The amount value as a decimal string in the currency's major unit. For example, for USD, '1234.56' represents 1234 dollars and 56 cents. For JPY, '1234' represents 1234 yen. A currency with minor units requires at least one decimal place and supports up to the number of decimal places defined by the currency's minor units. A currency without minor units does not support any decimal places.",
			InnerField: "value",
		},
	},
})

var cardPushTransfersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Card Push Transfer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-push-transfer-id",
			Usage:    "The identifier of the Card Push Transfer.",
			Required: true,
		},
	},
	Action:          handleCardPushTransfersRetrieve,
	HideHelpCommand: true,
}

var cardPushTransfersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Card Push Transfers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter Card Push Transfers to ones belonging to the specified Account.",
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
	Action:          handleCardPushTransfersList,
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
			Usage:      "Filter Card Push Transfers by status. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

var cardPushTransfersApprove = cli.Command{
	Name:    "approve",
	Usage:   "Approves a Card Push Transfer in a pending_approval state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-push-transfer-id",
			Usage:    "The identifier of the Card Push Transfer to approve.",
			Required: true,
		},
	},
	Action:          handleCardPushTransfersApprove,
	HideHelpCommand: true,
}

var cardPushTransfersCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancels a Card Push Transfer in a pending_approval state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-push-transfer-id",
			Usage:    "The identifier of the pending Card Push Transfer to cancel.",
			Required: true,
		},
	},
	Action:          handleCardPushTransfersCancel,
	HideHelpCommand: true,
}

func handleCardPushTransfersCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CardPushTransferNewParams{}

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
	_, err = client.CardPushTransfers.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "card-push-transfers create", obj, format, transform)
}

func handleCardPushTransfersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-push-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("card-push-transfer-id", unusedArgs[0])
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
	_, err = client.CardPushTransfers.Get(ctx, cmd.Value("card-push-transfer-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "card-push-transfers retrieve", obj, format, transform)
}

func handleCardPushTransfersList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CardPushTransferListParams{}

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
		_, err = client.CardPushTransfers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "card-push-transfers list", obj, format, transform)
	} else {
		iter := client.CardPushTransfers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "card-push-transfers list", iter, format, transform, maxItems)
	}
}

func handleCardPushTransfersApprove(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-push-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("card-push-transfer-id", unusedArgs[0])
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
	_, err = client.CardPushTransfers.Approve(ctx, cmd.Value("card-push-transfer-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "card-push-transfers approve", obj, format, transform)
}

func handleCardPushTransfersCancel(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-push-transfer-id") && len(unusedArgs) > 0 {
		cmd.Set("card-push-transfer-id", unusedArgs[0])
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
	_, err = client.CardPushTransfers.Cancel(ctx, cmd.Value("card-push-transfer-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "card-push-transfers cancel", obj, format, transform)
}
