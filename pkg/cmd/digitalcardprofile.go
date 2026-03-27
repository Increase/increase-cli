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

var digitalCardProfilesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Digital Card Profile",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "app-icon-file-id",
			Usage:    "The identifier of the File containing the card's icon image.",
			Required: true,
			BodyPath: "app_icon_file_id",
		},
		&requestflag.Flag[string]{
			Name:     "background-image-file-id",
			Usage:    "The identifier of the File containing the card's front image.",
			Required: true,
			BodyPath: "background_image_file_id",
		},
		&requestflag.Flag[string]{
			Name:     "card-description",
			Usage:    "A user-facing description for the card itself.",
			Required: true,
			BodyPath: "card_description",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "A description you can use to identify the Card Profile.",
			Required: true,
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "issuer-name",
			Usage:    "A user-facing description for whoever is issuing the card.",
			Required: true,
			BodyPath: "issuer_name",
		},
		&requestflag.Flag[string]{
			Name:     "contact-email",
			Usage:    "An email address the user can contact to receive support for their card.",
			BodyPath: "contact_email",
		},
		&requestflag.Flag[string]{
			Name:     "contact-phone",
			Usage:    "A phone number the user can contact to receive support for their card.",
			BodyPath: "contact_phone",
		},
		&requestflag.Flag[string]{
			Name:     "contact-website",
			Usage:    "A website the user can visit to view and receive support for their card.",
			BodyPath: "contact_website",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "text-color",
			Usage:    "The Card's text color, specified as an RGB triple. The default is white.",
			BodyPath: "text_color",
		},
	},
	Action:          handleDigitalCardProfilesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"text-color": {
		&requestflag.InnerFlag[int64]{
			Name:       "text-color.blue",
			Usage:      "The value of the blue channel in the RGB color.",
			InnerField: "blue",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "text-color.green",
			Usage:      "The value of the green channel in the RGB color.",
			InnerField: "green",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "text-color.red",
			Usage:      "The value of the red channel in the RGB color.",
			InnerField: "red",
		},
	},
})

var digitalCardProfilesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Digital Card Profile",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "digital-card-profile-id",
			Usage:    "The identifier of the Digital Card Profile.",
			Required: true,
		},
	},
	Action:          handleDigitalCardProfilesRetrieve,
	HideHelpCommand: true,
}

var digitalCardProfilesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Card Profiles",
	Suggest: true,
	Flags: []cli.Flag{
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
	Action:          handleDigitalCardProfilesList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"status": {
		&requestflag.InnerFlag[[]string]{
			Name:       "status.in",
			Usage:      "Filter Digital Card Profiles for those with the specified digital wallet status or statuses. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

var digitalCardProfilesArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive a Digital Card Profile",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "digital-card-profile-id",
			Usage:    "The identifier of the Digital Card Profile to archive.",
			Required: true,
		},
	},
	Action:          handleDigitalCardProfilesArchive,
	HideHelpCommand: true,
}

var digitalCardProfilesClone = requestflag.WithInnerFlags(cli.Command{
	Name:    "clone",
	Usage:   "Clones a Digital Card Profile",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "digital-card-profile-id",
			Usage:    "The identifier of the Digital Card Profile to clone.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "app-icon-file-id",
			Usage:    "The identifier of the File containing the card's icon image.",
			BodyPath: "app_icon_file_id",
		},
		&requestflag.Flag[string]{
			Name:     "background-image-file-id",
			Usage:    "The identifier of the File containing the card's front image.",
			BodyPath: "background_image_file_id",
		},
		&requestflag.Flag[string]{
			Name:     "card-description",
			Usage:    "A user-facing description for the card itself.",
			BodyPath: "card_description",
		},
		&requestflag.Flag[string]{
			Name:     "contact-email",
			Usage:    "An email address the user can contact to receive support for their card.",
			BodyPath: "contact_email",
		},
		&requestflag.Flag[string]{
			Name:     "contact-phone",
			Usage:    "A phone number the user can contact to receive support for their card.",
			BodyPath: "contact_phone",
		},
		&requestflag.Flag[string]{
			Name:     "contact-website",
			Usage:    "A website the user can visit to view and receive support for their card.",
			BodyPath: "contact_website",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "A description you can use to identify the Card Profile.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "issuer-name",
			Usage:    "A user-facing description for whoever is issuing the card.",
			BodyPath: "issuer_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "text-color",
			Usage:    "The Card's text color, specified as an RGB triple. The default is white.",
			BodyPath: "text_color",
		},
	},
	Action:          handleDigitalCardProfilesClone,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"text-color": {
		&requestflag.InnerFlag[int64]{
			Name:       "text-color.blue",
			Usage:      "The value of the blue channel in the RGB color.",
			InnerField: "blue",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "text-color.green",
			Usage:      "The value of the green channel in the RGB color.",
			InnerField: "green",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "text-color.red",
			Usage:      "The value of the red channel in the RGB color.",
			InnerField: "red",
		},
	},
})

func handleDigitalCardProfilesCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.DigitalCardProfileNewParams{}

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
	_, err = client.DigitalCardProfiles.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "digital-card-profiles create", obj, format, transform)
}

func handleDigitalCardProfilesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("digital-card-profile-id") && len(unusedArgs) > 0 {
		cmd.Set("digital-card-profile-id", unusedArgs[0])
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
	_, err = client.DigitalCardProfiles.Get(ctx, cmd.Value("digital-card-profile-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "digital-card-profiles retrieve", obj, format, transform)
}

func handleDigitalCardProfilesList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.DigitalCardProfileListParams{}

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
		_, err = client.DigitalCardProfiles.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "digital-card-profiles list", obj, format, transform)
	} else {
		iter := client.DigitalCardProfiles.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "digital-card-profiles list", iter, format, transform, maxItems)
	}
}

func handleDigitalCardProfilesArchive(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("digital-card-profile-id") && len(unusedArgs) > 0 {
		cmd.Set("digital-card-profile-id", unusedArgs[0])
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
	_, err = client.DigitalCardProfiles.Archive(ctx, cmd.Value("digital-card-profile-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "digital-card-profiles archive", obj, format, transform)
}

func handleDigitalCardProfilesClone(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("digital-card-profile-id") && len(unusedArgs) > 0 {
		cmd.Set("digital-card-profile-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.DigitalCardProfileCloneParams{}

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
	_, err = client.DigitalCardProfiles.Clone(
		ctx,
		cmd.Value("digital-card-profile-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "digital-card-profiles clone", obj, format, transform)
}
