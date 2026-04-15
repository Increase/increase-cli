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

var cardsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Card",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "The Account the card should belong to.",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "authorization-controls",
			Usage:    "Controls that restrict how this card can be used.",
			BodyPath: "authorization_controls",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "billing-address",
			Usage:    "The card's billing address.",
			BodyPath: "billing_address",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "The description you choose to give the card.",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "digital-wallet",
			Usage:    "The contact information used in the two-factor steps for digital wallet card creation. To add the card to a digital wallet, you may supply an email or phone number with this request. Otherwise, subscribe and then action a Real Time Decision with the category `digital_wallet_token_requested` or `digital_wallet_authentication_requested`.",
			BodyPath: "digital_wallet",
		},
		&requestflag.Flag[string]{
			Name:     "entity-id",
			Usage:    "The Entity the card belongs to. You only need to supply this in rare situations when the card is not for the Account holder.",
			BodyPath: "entity_id",
		},
	},
	Action:          handleCardsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"authorization-controls": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "authorization-controls.merchant-acceptor-identifier",
			Usage:      "Restricts which Merchant Acceptor IDs are allowed or blocked for authorizations on this card.",
			InnerField: "merchant_acceptor_identifier",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "authorization-controls.merchant-category-code",
			Usage:      "Restricts which Merchant Category Codes are allowed or blocked for authorizations on this card.",
			InnerField: "merchant_category_code",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "authorization-controls.merchant-country",
			Usage:      "Restricts which merchant countries are allowed or blocked for authorizations on this card.",
			InnerField: "merchant_country",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "authorization-controls.usage",
			Usage:      "Controls how many times this card can be used.",
			InnerField: "usage",
		},
	},
	"billing-address": {
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.city",
			Usage:      "The city of the billing address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.line1",
			Usage:      "The first line of the billing address.",
			InnerField: "line1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.postal-code",
			Usage:      "The postal code of the billing address.",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.state",
			Usage:      "The US state of the billing address.",
			InnerField: "state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.line2",
			Usage:      "The second line of the billing address.",
			InnerField: "line2",
		},
	},
	"digital-wallet": {
		&requestflag.InnerFlag[string]{
			Name:       "digital-wallet.digital-card-profile-id",
			Usage:      "The digital card profile assigned to this digital card.",
			InnerField: "digital_card_profile_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "digital-wallet.email",
			Usage:      "An email address that can be used to contact and verify the cardholder via one-time passcode over email.",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "digital-wallet.phone",
			Usage:      "A phone number that can be used to contact and verify the cardholder via one-time passcode over SMS.",
			InnerField: "phone",
		},
	},
})

var cardsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Card",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-id",
			Usage:    "The identifier of the Card.",
			Required: true,
		},
	},
	Action:          handleCardsRetrieve,
	HideHelpCommand: true,
}

var cardsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a Card",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-id",
			Usage:    "The card identifier.",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "authorization-controls",
			Usage:    "Controls that restrict how this card can be used.",
			BodyPath: "authorization_controls",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "billing-address",
			Usage:    "The card's updated billing address.",
			BodyPath: "billing_address",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "The description you choose to give the card.",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "digital-wallet",
			Usage:    "The contact information used in the two-factor steps for digital wallet card creation. At least one field must be present to complete the digital wallet steps.",
			BodyPath: "digital_wallet",
		},
		&requestflag.Flag[string]{
			Name:     "entity-id",
			Usage:    "The Entity the card belongs to. You only need to supply this in rare situations when the card is not for the Account holder.",
			BodyPath: "entity_id",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The status to update the Card with.",
			BodyPath: "status",
		},
	},
	Action:          handleCardsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"authorization-controls": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "authorization-controls.merchant-acceptor-identifier",
			Usage:      "Restricts which Merchant Acceptor IDs are allowed or blocked for authorizations on this card.",
			InnerField: "merchant_acceptor_identifier",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "authorization-controls.merchant-category-code",
			Usage:      "Restricts which Merchant Category Codes are allowed or blocked for authorizations on this card.",
			InnerField: "merchant_category_code",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "authorization-controls.merchant-country",
			Usage:      "Restricts which merchant countries are allowed or blocked for authorizations on this card.",
			InnerField: "merchant_country",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "authorization-controls.usage",
			Usage:      "Controls how many times this card can be used.",
			InnerField: "usage",
		},
	},
	"billing-address": {
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.city",
			Usage:      "The city of the billing address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.line1",
			Usage:      "The first line of the billing address.",
			InnerField: "line1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.postal-code",
			Usage:      "The postal code of the billing address.",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.state",
			Usage:      "The US state of the billing address.",
			InnerField: "state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.line2",
			Usage:      "The second line of the billing address.",
			InnerField: "line2",
		},
	},
	"digital-wallet": {
		&requestflag.InnerFlag[string]{
			Name:       "digital-wallet.digital-card-profile-id",
			Usage:      "The digital card profile assigned to this digital card.",
			InnerField: "digital_card_profile_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "digital-wallet.email",
			Usage:      "An email address that can be used to verify the cardholder via one-time passcode over email.",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "digital-wallet.phone",
			Usage:      "A phone number that can be used to verify the cardholder via one-time passcode over SMS.",
			InnerField: "phone",
		},
	},
})

var cardsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Cards",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter Cards to ones belonging to the specified Account.",
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
	Action:          handleCardsList,
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
			Usage:      "Filter Cards by status. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

var cardsCreateDetailsIframe = cli.Command{
	Name:    "create-details-iframe",
	Usage:   "Create an iframe URL for a Card to display the card details. More details about\nstyling and usage can be found in the\n[documentation](/documentation/embedded-card-component).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-id",
			Usage:    "The identifier of the Card to create an iframe for.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "physical-card-id",
			Usage:    "The identifier of the Physical Card to create an iframe for. This will inform the appearance of the card rendered in the iframe.",
			BodyPath: "physical_card_id",
		},
	},
	Action:          handleCardsCreateDetailsIframe,
	HideHelpCommand: true,
}

var cardsDetails = cli.Command{
	Name:    "details",
	Usage:   "Sensitive details for a Card include the primary account number, expiry, card\nverification code, and PIN.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-id",
			Usage:    "The identifier of the Card to retrieve details for.",
			Required: true,
		},
	},
	Action:          handleCardsDetails,
	HideHelpCommand: true,
}

var cardsUpdatePin = cli.Command{
	Name:    "update-pin",
	Usage:   "Update a Card's PIN",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-id",
			Usage:    "The identifier of the Card to update the PIN for.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "pin",
			Usage:    "The 4-digit PIN for the card, for use with ATMs.",
			Required: true,
			BodyPath: "pin",
		},
	},
	Action:          handleCardsUpdatePin,
	HideHelpCommand: true,
}

func handleCardsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CardNewParams{}

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
	_, err = client.Cards.New(ctx, params, options...)
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
		Title:          "cards create",
		Transform:      transform,
	})
}

func handleCardsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
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
	_, err = client.Cards.Get(ctx, cmd.Value("card-id").(string), options...)
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
		Title:          "cards retrieve",
		Transform:      transform,
	})
}

func handleCardsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CardUpdateParams{}

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
	_, err = client.Cards.Update(
		ctx,
		cmd.Value("card-id").(string),
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
		Title:          "cards update",
		Transform:      transform,
	})
}

func handleCardsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CardListParams{}

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
		_, err = client.Cards.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "cards list",
			Transform:      transform,
		})
	} else {
		iter := client.Cards.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "cards list",
			Transform:      transform,
		})
	}
}

func handleCardsCreateDetailsIframe(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CardNewDetailsIframeParams{}

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
	_, err = client.Cards.NewDetailsIframe(
		ctx,
		cmd.Value("card-id").(string),
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
		Title:          "cards create-details-iframe",
		Transform:      transform,
	})
}

func handleCardsDetails(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
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
	_, err = client.Cards.Details(ctx, cmd.Value("card-id").(string), options...)
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
		Title:          "cards details",
		Transform:      transform,
	})
}

func handleCardsUpdatePin(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.CardUpdatePinParams{}

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
	_, err = client.Cards.UpdatePin(
		ctx,
		cmd.Value("card-id").(string),
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
		Title:          "cards update-pin",
		Transform:      transform,
	})
}
