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

var supplementalDocumentsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a supplemental document for an Entity",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-id",
			Usage:    "The identifier of the Entity to associate with the supplemental document.",
			Required: true,
			BodyPath: "entity_id",
		},
		&requestflag.Flag[string]{
			Name:     "file-id",
			Usage:    "The identifier of the File containing the document.",
			Required: true,
			BodyPath: "file_id",
		},
	},
	Action:          handleSupplementalDocumentsCreate,
	HideHelpCommand: true,
}

var supplementalDocumentsList = cli.Command{
	Name:    "list",
	Usage:   "List Entity Supplemental Document Submissions",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-id",
			Usage:     "The identifier of the Entity to list supplemental documents for.",
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
	Action:          handleSupplementalDocumentsList,
	HideHelpCommand: true,
}

func handleSupplementalDocumentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SupplementalDocumentNewParams{}

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
	_, err = client.SupplementalDocuments.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "supplemental-documents create", obj, format, transform)
}

func handleSupplementalDocumentsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SupplementalDocumentListParams{}

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
		_, err = client.SupplementalDocuments.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "supplemental-documents list", obj, format, transform)
	} else {
		iter := client.SupplementalDocuments.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "supplemental-documents list", iter, format, transform, maxItems)
	}
}
