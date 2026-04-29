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

var cardValidationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a Card Validation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "The identifier of the Account from which to send the validation.",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:     "card-token-id",
			Usage:    "The Increase identifier for the Card Token that represents the card number you're validating.",
			Required: true,
			BodyPath: "card_token_id",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-category-code",
			Usage:    "A four-digit code (MCC) identifying the type of business or service provided by the merchant.",
			Required: true,
			BodyPath: "merchant_category_code",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-city-name",
			Usage:    "The city where the merchant (typically your business) is located.",
			Required: true,
			BodyPath: "merchant_city_name",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-name",
			Usage:    "The merchant name that will appear in the cardholder’s statement descriptor. Typically your business name.",
			Required: true,
			BodyPath: "merchant_name",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-postal-code",
			Usage:    "The postal code for the merchant’s (typically your business’s) location.",
			Required: true,
			BodyPath: "merchant_postal_code",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-state",
			Usage:    "The U.S. state where the merchant (typically your business) is located.",
			Required: true,
			BodyPath: "merchant_state",
		},
		&requestflag.Flag[string]{
			Name:     "cardholder-first-name",
			Usage:    "The cardholder's first name.",
			BodyPath: "cardholder_first_name",
		},
		&requestflag.Flag[string]{
			Name:     "cardholder-last-name",
			Usage:    "The cardholder's last name.",
			BodyPath: "cardholder_last_name",
		},
		&requestflag.Flag[string]{
			Name:     "cardholder-middle-name",
			Usage:    "The cardholder's middle name.",
			BodyPath: "cardholder_middle_name",
		},
		&requestflag.Flag[string]{
			Name:     "cardholder-postal-code",
			Usage:    "The postal code of the cardholder's address.",
			BodyPath: "cardholder_postal_code",
		},
		&requestflag.Flag[string]{
			Name:     "cardholder-street-address",
			Usage:    "The cardholder's street address.",
			BodyPath: "cardholder_street_address",
		},
	},
	Action:          handleCardValidationsCreate,
	HideHelpCommand: true,
}

var cardValidationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Card Validation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "card-validation-id",
			Usage:     "The identifier of the Card Validation.",
			Required:  true,
			PathParam: "card_validation_id",
		},
	},
	Action:          handleCardValidationsRetrieve,
	HideHelpCommand: true,
}

var cardValidationsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Card Validations",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter Card Validations to ones belonging to the specified Account.",
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
	Action:          handleCardValidationsList,
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
			Usage:      "Filter Card Validations by status. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

func handleCardValidationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := increase.CardValidationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.CardValidations.New(ctx, params, options...)
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
		Title:          "card-validations create",
		Transform:      transform,
	})
}

func handleCardValidationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-validation-id") && len(unusedArgs) > 0 {
		cmd.Set("card-validation-id", unusedArgs[0])
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
	_, err = client.CardValidations.Get(ctx, cmd.Value("card-validation-id").(string), options...)
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
		Title:          "card-validations retrieve",
		Transform:      transform,
	})
}

func handleCardValidationsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := increase.CardValidationListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.CardValidations.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "card-validations list",
			Transform:      transform,
		})
	} else {
		iter := client.CardValidations.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "card-validations list",
			Transform:      transform,
		})
	}
}
