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

var beneficialOwnersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a beneficial owner",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-id",
			Usage:    "The identifier of the Entity to associate with the new Beneficial Owner.",
			Required: true,
			BodyPath: "entity_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "individual",
			Usage:    "Personal details for the beneficial owner.",
			Required: true,
			BodyPath: "individual",
		},
		&requestflag.Flag[[]string]{
			Name:     "prong",
			Usage:    "Why this person is considered a beneficial owner of the entity. At least one option is required, if a person is both a control person and owner, submit an array containing both.",
			Required: true,
			BodyPath: "prongs",
		},
		&requestflag.Flag[string]{
			Name:     "company-title",
			Usage:    "This person's role or title within the entity.",
			BodyPath: "company_title",
		},
	},
	Action:          handleBeneficialOwnersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"individual": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "individual.address",
			Usage:      "The individual's physical address. Mail receiving locations like PO Boxes and PMB's are disallowed.",
			InnerField: "address",
		},
		&requestflag.InnerFlag[any]{
			Name:       "individual.date-of-birth",
			Usage:      "The person's date of birth in YYYY-MM-DD format.",
			InnerField: "date_of_birth",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "individual.identification",
			Usage:      "A means of verifying the person's identity.",
			InnerField: "identification",
		},
		&requestflag.InnerFlag[string]{
			Name:       "individual.name",
			Usage:      "The person's legal name.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "individual.confirmed-no-us-tax-id",
			Usage:      "The identification method for an individual can only be a passport, driver's license, or other document if you've confirmed the individual does not have a US tax id (either a Social Security Number or Individual Taxpayer Identification Number).",
			InnerField: "confirmed_no_us_tax_id",
		},
	},
})

var beneficialOwnersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Beneficial Owner",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-beneficial-owner-id",
			Usage:    "The identifier of the Beneficial Owner to retrieve.",
			Required: true,
		},
	},
	Action:          handleBeneficialOwnersRetrieve,
	HideHelpCommand: true,
}

var beneficialOwnersUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a Beneficial Owner",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-beneficial-owner-id",
			Usage:    "The identifier of the Beneficial Owner to update.",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "address",
			Usage:    "The individual's physical address. Mail receiving locations like PO Boxes and PMB's are disallowed.",
			BodyPath: "address",
		},
		&requestflag.Flag[bool]{
			Name:     "confirmed-no-us-tax-id",
			Usage:    "The identification method for an individual can only be a passport, driver's license, or other document if you've confirmed the individual does not have a US tax id (either a Social Security Number or Individual Taxpayer Identification Number).",
			BodyPath: "confirmed_no_us_tax_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "identification",
			Usage:    "A means of verifying the person's identity.",
			BodyPath: "identification",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The individual's legal name.",
			BodyPath: "name",
		},
	},
	Action:          handleBeneficialOwnersUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"address": {
		&requestflag.InnerFlag[string]{
			Name:       "address.city",
			Usage:      "The city, district, town, or village of the address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.country",
			Usage:      "The two-letter ISO 3166-1 alpha-2 code for the country of the address.",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.line1",
			Usage:      "The first line of the address. This is usually the street number and street.",
			InnerField: "line1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.line2",
			Usage:      "The second line of the address. This might be the floor or room number.",
			InnerField: "line2",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.state",
			Usage:      "The two-letter United States Postal Service (USPS) abbreviation for the US state, province, or region of the address. Required in certain countries.",
			InnerField: "state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.zip",
			Usage:      "The ZIP or postal code of the address. Required in certain countries.",
			InnerField: "zip",
		},
	},
	"identification": {
		&requestflag.InnerFlag[string]{
			Name:       "identification.method",
			Usage:      "A method that can be used to verify the individual's identity.",
			InnerField: "method",
		},
		&requestflag.InnerFlag[string]{
			Name:       "identification.number",
			Usage:      "An identification number that can be used to verify the individual's identity, such as a social security number.",
			InnerField: "number",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "identification.drivers-license",
			Usage:      "Information about the United States driver's license used for identification. Required if `method` is equal to `drivers_license`.",
			InnerField: "drivers_license",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "identification.other",
			Usage:      "Information about the identification document provided. Required if `method` is equal to `other`.",
			InnerField: "other",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "identification.passport",
			Usage:      "Information about the passport used for identification. Required if `method` is equal to `passport`.",
			InnerField: "passport",
		},
	},
})

var beneficialOwnersList = cli.Command{
	Name:    "list",
	Usage:   "List Beneficial Owners",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-id",
			Usage:     "The identifier of the Entity to list beneficial owners for. Only `corporation` entities have beneficial owners.",
			Required:  true,
			QueryPath: "entity_id",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBeneficialOwnersList,
	HideHelpCommand: true,
}

var beneficialOwnersArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive a Beneficial Owner",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-beneficial-owner-id",
			Usage:    "The identifier of the Beneficial Owner to archive.",
			Required: true,
		},
	},
	Action:          handleBeneficialOwnersArchive,
	HideHelpCommand: true,
}

func handleBeneficialOwnersCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.BeneficialOwnerNewParams{}

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
	_, err = client.BeneficialOwners.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "beneficial-owners create", obj, format, explicitFormat, transform)
}

func handleBeneficialOwnersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-beneficial-owner-id") && len(unusedArgs) > 0 {
		cmd.Set("entity-beneficial-owner-id", unusedArgs[0])
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
	_, err = client.BeneficialOwners.Get(ctx, cmd.Value("entity-beneficial-owner-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "beneficial-owners retrieve", obj, format, explicitFormat, transform)
}

func handleBeneficialOwnersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-beneficial-owner-id") && len(unusedArgs) > 0 {
		cmd.Set("entity-beneficial-owner-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.BeneficialOwnerUpdateParams{}

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
	_, err = client.BeneficialOwners.Update(
		ctx,
		cmd.Value("entity-beneficial-owner-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "beneficial-owners update", obj, format, explicitFormat, transform)
}

func handleBeneficialOwnersList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.BeneficialOwnerListParams{}

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
		_, err = client.BeneficialOwners.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "beneficial-owners list", obj, format, explicitFormat, transform)
	} else {
		iter := client.BeneficialOwners.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "beneficial-owners list", iter, format, explicitFormat, transform, maxItems)
	}
}

func handleBeneficialOwnersArchive(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-beneficial-owner-id") && len(unusedArgs) > 0 {
		cmd.Set("entity-beneficial-owner-id", unusedArgs[0])
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
	_, err = client.BeneficialOwners.Archive(ctx, cmd.Value("entity-beneficial-owner-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "beneficial-owners archive", obj, format, explicitFormat, transform)
}
