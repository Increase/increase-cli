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

var bookkeepingEntrySetsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Bookkeeping Entry Set",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "entry",
			Usage:    "The bookkeeping entries.",
			Required: true,
			BodyPath: "entries",
		},
		&requestflag.Flag[any]{
			Name:     "date",
			Usage:    "The date of the transaction. Optional if `transaction_id` is provided, in which case we use the `date` of that transaction. Required otherwise.",
			BodyPath: "date",
		},
		&requestflag.Flag[string]{
			Name:     "transaction-id",
			Usage:    "The identifier of the Transaction related to this entry set, if any.",
			BodyPath: "transaction_id",
		},
	},
	Action:          handleBookkeepingEntrySetsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"entry": {
		&requestflag.InnerFlag[string]{
			Name:       "entry.account-id",
			Usage:      "The identifier for the Bookkeeping Account impacted by this entry.",
			InnerField: "account_id",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "entry.amount",
			Usage:      "The entry amount in the minor unit of the account currency. For dollars, for example, this is cents. Debit entries have positive amounts; credit entries have negative amounts.",
			InnerField: "amount",
		},
	},
})

var bookkeepingEntrySetsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Bookkeeping Entry Set",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bookkeeping-entry-set-id",
			Usage:    "The identifier of the Bookkeeping Entry Set.",
			Required: true,
		},
	},
	Action:          handleBookkeepingEntrySetsRetrieve,
	HideHelpCommand: true,
}

var bookkeepingEntrySetsList = cli.Command{
	Name:    "list",
	Usage:   "List Bookkeeping Entry Sets",
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
		&requestflag.Flag[string]{
			Name:      "transaction-id",
			Usage:     "Filter to the Bookkeeping Entry Set that maps to this Transaction.",
			QueryPath: "transaction_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBookkeepingEntrySetsList,
	HideHelpCommand: true,
}

func handleBookkeepingEntrySetsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.BookkeepingEntrySetNewParams{}

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
	_, err = client.BookkeepingEntrySets.New(ctx, params, options...)
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
		Title:          "bookkeeping-entry-sets create",
		Transform:      transform,
	})
}

func handleBookkeepingEntrySetsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("bookkeeping-entry-set-id") && len(unusedArgs) > 0 {
		cmd.Set("bookkeeping-entry-set-id", unusedArgs[0])
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
	_, err = client.BookkeepingEntrySets.Get(ctx, cmd.Value("bookkeeping-entry-set-id").(string), options...)
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
		Title:          "bookkeeping-entry-sets retrieve",
		Transform:      transform,
	})
}

func handleBookkeepingEntrySetsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.BookkeepingEntrySetListParams{}

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
		_, err = client.BookkeepingEntrySets.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "bookkeeping-entry-sets list",
			Transform:      transform,
		})
	} else {
		iter := client.BookkeepingEntrySets.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "bookkeeping-entry-sets list",
			Transform:      transform,
		})
	}
}
