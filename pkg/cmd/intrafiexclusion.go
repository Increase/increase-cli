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

var intrafiExclusionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an IntraFi Exclusion",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-id",
			Usage:    "The identifier of the Entity whose deposits will be excluded.",
			Required: true,
			BodyPath: "entity_id",
		},
		&requestflag.Flag[string]{
			Name:     "fdic-certificate-number",
			Usage:    "The FDIC certificate number of the financial institution to be excluded. An FDIC certificate number uniquely identifies a financial institution, and is different than a routing number. To find one, we recommend searching by Bank Name using the [FDIC's bankfind tool](https://banks.data.fdic.gov/bankfind-suite/bankfind).",
			Required: true,
			BodyPath: "fdic_certificate_number",
		},
	},
	Action:          handleIntrafiExclusionsCreate,
	HideHelpCommand: true,
}

var intrafiExclusionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get an IntraFi Exclusion",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "intrafi-exclusion-id",
			Usage:    "The identifier of the IntraFi Exclusion to retrieve.",
			Required: true,
		},
	},
	Action:          handleIntrafiExclusionsRetrieve,
	HideHelpCommand: true,
}

var intrafiExclusionsList = cli.Command{
	Name:    "list",
	Usage:   "List IntraFi Exclusions",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Return the page of entries after this one.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "entity-id",
			Usage:     "Filter IntraFi Exclusions for those belonging to the specified Entity.",
			QueryPath: "entity_id",
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
	Action:          handleIntrafiExclusionsList,
	HideHelpCommand: true,
}

var intrafiExclusionsArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive an IntraFi Exclusion",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "intrafi-exclusion-id",
			Usage:    "The identifier of the IntraFi Exclusion request to archive. It may take 5 business days for an exclusion removal to be processed. Removing an exclusion does not guarantee that funds will be swept to the previously-excluded bank.",
			Required: true,
		},
	},
	Action:          handleIntrafiExclusionsArchive,
	HideHelpCommand: true,
}

func handleIntrafiExclusionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.IntrafiExclusionNewParams{}

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
	_, err = client.IntrafiExclusions.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "intrafi-exclusions create", obj, format, transform)
}

func handleIntrafiExclusionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("intrafi-exclusion-id") && len(unusedArgs) > 0 {
		cmd.Set("intrafi-exclusion-id", unusedArgs[0])
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
	_, err = client.IntrafiExclusions.Get(ctx, cmd.Value("intrafi-exclusion-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "intrafi-exclusions retrieve", obj, format, transform)
}

func handleIntrafiExclusionsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.IntrafiExclusionListParams{}

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
		_, err = client.IntrafiExclusions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "intrafi-exclusions list", obj, format, transform)
	} else {
		iter := client.IntrafiExclusions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "intrafi-exclusions list", iter, format, transform, maxItems)
	}
}

func handleIntrafiExclusionsArchive(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("intrafi-exclusion-id") && len(unusedArgs) > 0 {
		cmd.Set("intrafi-exclusion-id", unusedArgs[0])
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
	_, err = client.IntrafiExclusions.Archive(ctx, cmd.Value("intrafi-exclusion-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "intrafi-exclusions archive", obj, format, transform)
}
