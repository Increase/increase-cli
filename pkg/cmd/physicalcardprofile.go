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

var physicalCardProfilesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Physical Card Profile",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "carrier-image-file-id",
			Usage:    "The identifier of the File containing the physical card's carrier image.",
			Required: true,
			BodyPath: "carrier_image_file_id",
		},
		&requestflag.Flag[string]{
			Name:     "contact-phone",
			Usage:    "A phone number the user can contact to receive support for their card.",
			Required: true,
			BodyPath: "contact_phone",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "A description you can use to identify the Card Profile.",
			Required: true,
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "front-image-file-id",
			Usage:    "The identifier of the File containing the physical card's front image.",
			Required: true,
			BodyPath: "front_image_file_id",
		},
		&requestflag.Flag[string]{
			Name:     "program-id",
			Usage:    "The identifier for the Program that this Physical Card Profile falls under.",
			Required: true,
			BodyPath: "program_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "front-text",
			Usage:    "Text printed on the front of the card. Reach out to [support@increase.com](mailto:support@increase.com) for more information.",
			BodyPath: "front_text",
		},
	},
	Action:          handlePhysicalCardProfilesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"front-text": {
		&requestflag.InnerFlag[string]{
			Name:       "front-text.line1",
			Usage:      "The first line of text on the front of the card.",
			InnerField: "line1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "front-text.line2",
			Usage:      "The second line of text on the front of the card. Providing a second line moves the first line slightly higher and prints the second line in the spot where the first line would have otherwise been printed.",
			InnerField: "line2",
		},
	},
})

var physicalCardProfilesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Card Profile",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "physical-card-profile-id",
			Usage:    "The identifier of the Card Profile.",
			Required: true,
		},
	},
	Action:          handlePhysicalCardProfilesRetrieve,
	HideHelpCommand: true,
}

var physicalCardProfilesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Physical Card Profiles",
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
	Action:          handlePhysicalCardProfilesList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"status": {
		&requestflag.InnerFlag[[]string]{
			Name:       "status.in",
			Usage:      "Filter Physical Card Profiles for those with the specified statuses. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

var physicalCardProfilesArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive a Physical Card Profile",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "physical-card-profile-id",
			Usage:    "The identifier of the Physical Card Profile to archive.",
			Required: true,
		},
	},
	Action:          handlePhysicalCardProfilesArchive,
	HideHelpCommand: true,
}

var physicalCardProfilesClone = requestflag.WithInnerFlags(cli.Command{
	Name:    "clone",
	Usage:   "Clone a Physical Card Profile",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "physical-card-profile-id",
			Usage:    "The identifier of the Physical Card Profile to clone.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "carrier-image-file-id",
			Usage:    "The identifier of the File containing the physical card's carrier image.",
			BodyPath: "carrier_image_file_id",
		},
		&requestflag.Flag[string]{
			Name:     "contact-phone",
			Usage:    "A phone number the user can contact to receive support for their card.",
			BodyPath: "contact_phone",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "A description you can use to identify the Card Profile.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "front-image-file-id",
			Usage:    "The identifier of the File containing the physical card's front image.",
			BodyPath: "front_image_file_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "front-text",
			Usage:    "Text printed on the front of the card. Reach out to [support@increase.com](mailto:support@increase.com) for more information.",
			BodyPath: "front_text",
		},
		&requestflag.Flag[string]{
			Name:     "program-id",
			Usage:    "The identifier of the Program to use for the cloned Physical Card Profile.",
			BodyPath: "program_id",
		},
	},
	Action:          handlePhysicalCardProfilesClone,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"front-text": {
		&requestflag.InnerFlag[string]{
			Name:       "front-text.line1",
			Usage:      "The first line of text on the front of the card.",
			InnerField: "line1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "front-text.line2",
			Usage:      "The second line of text on the front of the card. Providing a second line moves the first line slightly higher and prints the second line in the spot where the first line would have otherwise been printed.",
			InnerField: "line2",
		},
	},
})

func handlePhysicalCardProfilesCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.PhysicalCardProfileNewParams{}

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
	_, err = client.PhysicalCardProfiles.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "physical-card-profiles create", obj, format, explicitFormat, transform)
}

func handlePhysicalCardProfilesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("physical-card-profile-id") && len(unusedArgs) > 0 {
		cmd.Set("physical-card-profile-id", unusedArgs[0])
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
	_, err = client.PhysicalCardProfiles.Get(ctx, cmd.Value("physical-card-profile-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "physical-card-profiles retrieve", obj, format, explicitFormat, transform)
}

func handlePhysicalCardProfilesList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.PhysicalCardProfileListParams{}

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
		_, err = client.PhysicalCardProfiles.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "physical-card-profiles list", obj, format, explicitFormat, transform)
	} else {
		iter := client.PhysicalCardProfiles.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "physical-card-profiles list", iter, format, explicitFormat, transform, maxItems)
	}
}

func handlePhysicalCardProfilesArchive(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("physical-card-profile-id") && len(unusedArgs) > 0 {
		cmd.Set("physical-card-profile-id", unusedArgs[0])
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
	_, err = client.PhysicalCardProfiles.Archive(ctx, cmd.Value("physical-card-profile-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "physical-card-profiles archive", obj, format, explicitFormat, transform)
}

func handlePhysicalCardProfilesClone(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("physical-card-profile-id") && len(unusedArgs) > 0 {
		cmd.Set("physical-card-profile-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.PhysicalCardProfileCloneParams{}

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
	_, err = client.PhysicalCardProfiles.Clone(
		ctx,
		cmd.Value("physical-card-profile-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "physical-card-profiles clone", obj, format, explicitFormat, transform)
}
