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

var entitiesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create an Entity",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "structure",
			Usage:    "The type of Entity to create.",
			Required: true,
			BodyPath: "structure",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "corporation",
			Usage:    "Details of the corporation entity to create. Required if `structure` is equal to `corporation`.",
			BodyPath: "corporation",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "The description you choose to give the entity.",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "government-authority",
			Usage:    "Details of the Government Authority entity to create. Required if `structure` is equal to `government_authority`.",
			BodyPath: "government_authority",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "joint",
			Usage:    "Details of the joint entity to create. Required if `structure` is equal to `joint`.",
			BodyPath: "joint",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "natural-person",
			Usage:    "Details of the natural person entity to create. Required if `structure` is equal to `natural_person`. Natural people entities should be submitted with `social_security_number` or `individual_taxpayer_identification_number` identification methods.",
			BodyPath: "natural_person",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "risk-rating",
			Usage:    "An assessment of the entity's potential risk of involvement in financial crimes, such as money laundering.",
			BodyPath: "risk_rating",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "supplemental-document",
			Usage:    "Additional documentation associated with the entity.",
			BodyPath: "supplemental_documents",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "terms-agreement",
			Usage:    "The terms that the Entity agreed to. Not all programs are required to submit this data.",
			BodyPath: "terms_agreements",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "third-party-verification",
			Usage:    "If you are using a third-party service for identity verification, you can use this field to associate this Entity with the identifier that represents them in that service.",
			BodyPath: "third_party_verification",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "trust",
			Usage:    "Details of the trust entity to create. Required if `structure` is equal to `trust`.",
			BodyPath: "trust",
		},
	},
	Action:          handleEntitiesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"corporation": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "corporation.address",
			Usage:      "The entity's physical address. Mail receiving locations like PO Boxes and PMB's are disallowed.",
			InnerField: "address",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "corporation.beneficial-owners",
			Usage:      "The identifying details of each person who owns 25% or more of the business and one control person, like the CEO, CFO, or other executive. You can submit between 1 and 5 people to this list.",
			InnerField: "beneficial_owners",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "corporation.legal-identifier",
			Usage:      "The legal identifier of the corporation. This is usually the Employer Identification Number (EIN).",
			InnerField: "legal_identifier",
		},
		&requestflag.InnerFlag[string]{
			Name:       "corporation.name",
			Usage:      "The legal name of the corporation.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "corporation.beneficial-ownership-exemption-reason",
			Usage:      "If the entity is exempt from the requirement to submit beneficial owners, provide the justification. If a reason is provided, you do not need to submit a list of beneficial owners.",
			InnerField: "beneficial_ownership_exemption_reason",
		},
		&requestflag.InnerFlag[string]{
			Name:       "corporation.email",
			Usage:      "An email address for the business. Not every program requires an email for submitted Entities.",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "corporation.incorporation-state",
			Usage:      "The two-letter United States Postal Service (USPS) abbreviation for the corporation's state of incorporation.",
			InnerField: "incorporation_state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "corporation.industry-code",
			Usage:      "The North American Industry Classification System (NAICS) code for the corporation's primary line of business. This is a number, like `5132` for `Software Publishers`. A full list of classification codes is available [here](https://increase.com/documentation/data-dictionary#north-american-industry-classification-system-codes).",
			InnerField: "industry_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "corporation.website",
			Usage:      "The website of the corporation.",
			InnerField: "website",
		},
	},
	"government-authority": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "government-authority.address",
			Usage:      "The entity's physical address. Mail receiving locations like PO Boxes and PMB's are disallowed.",
			InnerField: "address",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "government-authority.authorized-persons",
			Usage:      "The identifying details of authorized officials acting on the entity's behalf.",
			InnerField: "authorized_persons",
		},
		&requestflag.InnerFlag[string]{
			Name:       "government-authority.category",
			Usage:      "The category of the government authority.",
			InnerField: "category",
		},
		&requestflag.InnerFlag[string]{
			Name:       "government-authority.name",
			Usage:      "The legal name of the government authority.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "government-authority.tax-identifier",
			Usage:      "The Employer Identification Number (EIN) for the government authority.",
			InnerField: "tax_identifier",
		},
		&requestflag.InnerFlag[string]{
			Name:       "government-authority.website",
			Usage:      "The website of the government authority.",
			InnerField: "website",
		},
	},
	"joint": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "joint.individuals",
			Usage:      "The two individuals that share control of the entity.",
			InnerField: "individuals",
		},
	},
	"natural-person": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "natural-person.address",
			Usage:      "The individual's physical address. Mail receiving locations like PO Boxes and PMB's are disallowed.",
			InnerField: "address",
		},
		&requestflag.InnerFlag[any]{
			Name:       "natural-person.date-of-birth",
			Usage:      "The person's date of birth in YYYY-MM-DD format.",
			InnerField: "date_of_birth",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "natural-person.identification",
			Usage:      "A means of verifying the person's identity.",
			InnerField: "identification",
		},
		&requestflag.InnerFlag[string]{
			Name:       "natural-person.name",
			Usage:      "The person's legal name.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "natural-person.confirmed-no-us-tax-id",
			Usage:      "The identification method for an individual can only be a passport, driver's license, or other document if you've confirmed the individual does not have a US tax id (either a Social Security Number or Individual Taxpayer Identification Number).",
			InnerField: "confirmed_no_us_tax_id",
		},
	},
	"risk-rating": {
		&requestflag.InnerFlag[any]{
			Name:       "risk-rating.rated-at",
			Usage:      "The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the risk rating was performed.",
			InnerField: "rated_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "risk-rating.rating",
			Usage:      "The rating given to this entity.",
			InnerField: "rating",
		},
	},
	"supplemental-document": {
		&requestflag.InnerFlag[string]{
			Name:       "supplemental-document.file-id",
			Usage:      "The identifier of the File containing the document.",
			InnerField: "file_id",
		},
	},
	"terms-agreement": {
		&requestflag.InnerFlag[any]{
			Name:       "terms-agreement.agreed-at",
			Usage:      "The timestamp of when the Entity agreed to the terms.",
			InnerField: "agreed_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "terms-agreement.ip-address",
			Usage:      "The IP address the Entity accessed reviewed the terms from.",
			InnerField: "ip_address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "terms-agreement.terms-url",
			Usage:      "The URL of the terms agreement. This link will be provided by your bank partner.",
			InnerField: "terms_url",
		},
	},
	"third-party-verification": {
		&requestflag.InnerFlag[string]{
			Name:       "third-party-verification.reference",
			Usage:      "The reference identifier for the third party verification.",
			InnerField: "reference",
		},
		&requestflag.InnerFlag[string]{
			Name:       "third-party-verification.vendor",
			Usage:      "The vendor that was used to perform the verification.",
			InnerField: "vendor",
		},
	},
	"trust": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "trust.address",
			Usage:      "The trust's physical address. Mail receiving locations like PO Boxes and PMB's are disallowed.",
			InnerField: "address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "trust.category",
			Usage:      "Whether the trust is `revocable` or `irrevocable`. Irrevocable trusts require their own Employer Identification Number. Revocable trusts require information about the individual `grantor` who created the trust.",
			InnerField: "category",
		},
		&requestflag.InnerFlag[string]{
			Name:       "trust.name",
			Usage:      "The legal name of the trust.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "trust.trustees",
			Usage:      "The trustees of the trust.",
			InnerField: "trustees",
		},
		&requestflag.InnerFlag[string]{
			Name:       "trust.formation-document-file-id",
			Usage:      "The identifier of the File containing the formation document of the trust.",
			InnerField: "formation_document_file_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "trust.formation-state",
			Usage:      "The two-letter United States Postal Service (USPS) abbreviation for the state in which the trust was formed.",
			InnerField: "formation_state",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "trust.grantor",
			Usage:      "The grantor of the trust. Required if `category` is equal to `revocable`.",
			InnerField: "grantor",
		},
		&requestflag.InnerFlag[string]{
			Name:       "trust.tax-identifier",
			Usage:      "The Employer Identification Number (EIN) for the trust. Required if `category` is equal to `irrevocable`.",
			InnerField: "tax_identifier",
		},
	},
})

var entitiesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an Entity",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-id",
			Usage:    "The identifier of the Entity to retrieve.",
			Required: true,
		},
	},
	Action:          handleEntitiesRetrieve,
	HideHelpCommand: true,
}

var entitiesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an Entity",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-id",
			Usage:    "The entity identifier.",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "corporation",
			Usage:    "Details of the corporation entity to update. If you specify this parameter and the entity is not a corporation, the request will fail.",
			BodyPath: "corporation",
		},
		&requestflag.Flag[any]{
			Name:     "details-confirmed-at",
			Usage:    "When your user last confirmed the Entity's details. Depending on your program, you may be required to affirmatively confirm details with your users on an annual basis.",
			BodyPath: "details_confirmed_at",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "government-authority",
			Usage:    "Details of the government authority entity to update. If you specify this parameter and the entity is not a government authority, the request will fail.",
			BodyPath: "government_authority",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "natural-person",
			Usage:    "Details of the natural person entity to update. If you specify this parameter and the entity is not a natural person, the request will fail.",
			BodyPath: "natural_person",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "risk-rating",
			Usage:    "An assessment of the entity’s potential risk of involvement in financial crimes, such as money laundering.",
			BodyPath: "risk_rating",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "terms-agreement",
			Usage:    "New terms that the Entity agreed to. Not all programs are required to submit this data. This will not archive previously submitted terms.",
			BodyPath: "terms_agreements",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "third-party-verification",
			Usage:    "If you are using a third-party service for identity verification, you can use this field to associate this Entity with the identifier that represents them in that service.",
			BodyPath: "third_party_verification",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "trust",
			Usage:    "Details of the trust entity to update. If you specify this parameter and the entity is not a trust, the request will fail.",
			BodyPath: "trust",
		},
	},
	Action:          handleEntitiesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"corporation": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "corporation.address",
			Usage:      "The entity's physical address. Mail receiving locations like PO Boxes and PMB's are disallowed.",
			InnerField: "address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "corporation.email",
			Usage:      "An email address for the business. Not every program requires an email for submitted Entities.",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "corporation.incorporation-state",
			Usage:      "The two-letter United States Postal Service (USPS) abbreviation for the corporation's state of incorporation.",
			InnerField: "incorporation_state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "corporation.industry-code",
			Usage:      "The North American Industry Classification System (NAICS) code for the corporation's primary line of business. This is a number, like `5132` for `Software Publishers`. A full list of classification codes is available [here](https://increase.com/documentation/data-dictionary#north-american-industry-classification-system-codes).",
			InnerField: "industry_code",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "corporation.legal-identifier",
			Usage:      "The legal identifier of the corporation. This is usually the Employer Identification Number (EIN).",
			InnerField: "legal_identifier",
		},
		&requestflag.InnerFlag[string]{
			Name:       "corporation.name",
			Usage:      "The legal name of the corporation.",
			InnerField: "name",
		},
	},
	"government-authority": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "government-authority.address",
			Usage:      "The entity's physical address. Mail receiving locations like PO Boxes and PMB's are disallowed.",
			InnerField: "address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "government-authority.name",
			Usage:      "The legal name of the government authority.",
			InnerField: "name",
		},
	},
	"natural-person": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "natural-person.address",
			Usage:      "The entity's physical address. Mail receiving locations like PO Boxes and PMB's are disallowed.",
			InnerField: "address",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "natural-person.confirmed-no-us-tax-id",
			Usage:      "The identification method for an individual can only be a passport, driver's license, or other document if you've confirmed the individual does not have a US tax id (either a Social Security Number or Individual Taxpayer Identification Number).",
			InnerField: "confirmed_no_us_tax_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "natural-person.identification",
			Usage:      "A means of verifying the person's identity.",
			InnerField: "identification",
		},
		&requestflag.InnerFlag[string]{
			Name:       "natural-person.name",
			Usage:      "The legal name of the natural person.",
			InnerField: "name",
		},
	},
	"risk-rating": {
		&requestflag.InnerFlag[any]{
			Name:       "risk-rating.rated-at",
			Usage:      "The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the risk rating was performed.",
			InnerField: "rated_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "risk-rating.rating",
			Usage:      "The rating given to this entity.",
			InnerField: "rating",
		},
	},
	"terms-agreement": {
		&requestflag.InnerFlag[any]{
			Name:       "terms-agreement.agreed-at",
			Usage:      "The timestamp of when the Entity agreed to the terms.",
			InnerField: "agreed_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "terms-agreement.ip-address",
			Usage:      "The IP address the Entity accessed reviewed the terms from.",
			InnerField: "ip_address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "terms-agreement.terms-url",
			Usage:      "The URL of the terms agreement. This link will be provided by your bank partner.",
			InnerField: "terms_url",
		},
	},
	"third-party-verification": {
		&requestflag.InnerFlag[string]{
			Name:       "third-party-verification.reference",
			Usage:      "The reference identifier for the third party verification.",
			InnerField: "reference",
		},
		&requestflag.InnerFlag[string]{
			Name:       "third-party-verification.vendor",
			Usage:      "The vendor that was used to perform the verification.",
			InnerField: "vendor",
		},
	},
	"trust": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "trust.address",
			Usage:      "The entity's physical address. Mail receiving locations like PO Boxes and PMB's are disallowed.",
			InnerField: "address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "trust.name",
			Usage:      "The legal name of the trust.",
			InnerField: "name",
		},
	},
})

var entitiesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Entities",
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
	Action:          handleEntitiesList,
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
			Usage:      "Filter Entities for those with the specified status or statuses. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

var entitiesArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive an Entity",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-id",
			Usage:    "The identifier of the Entity to archive. Any accounts associated with an entity must be closed before the entity can be archived.",
			Required: true,
		},
	},
	Action:          handleEntitiesArchive,
	HideHelpCommand: true,
}

func handleEntitiesCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.EntityNewParams{}

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
	_, err = client.Entities.New(ctx, params, options...)
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
		Title:          "entities create",
		Transform:      transform,
	})
}

func handleEntitiesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-id") && len(unusedArgs) > 0 {
		cmd.Set("entity-id", unusedArgs[0])
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
	_, err = client.Entities.Get(ctx, cmd.Value("entity-id").(string), options...)
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
		Title:          "entities retrieve",
		Transform:      transform,
	})
}

func handleEntitiesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-id") && len(unusedArgs) > 0 {
		cmd.Set("entity-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.EntityUpdateParams{}

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
	_, err = client.Entities.Update(
		ctx,
		cmd.Value("entity-id").(string),
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
		Title:          "entities update",
		Transform:      transform,
	})
}

func handleEntitiesList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.EntityListParams{}

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
		_, err = client.Entities.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "entities list",
			Transform:      transform,
		})
	} else {
		iter := client.Entities.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "entities list",
			Transform:      transform,
		})
	}
}

func handleEntitiesArchive(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-id") && len(unusedArgs) > 0 {
		cmd.Set("entity-id", unusedArgs[0])
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
	_, err = client.Entities.Archive(ctx, cmd.Value("entity-id").(string), options...)
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
		Title:          "entities archive",
		Transform:      transform,
	})
}
