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

var pendingTransactionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a pending transaction on an account. This can be useful to hold funds\nfor an external payment or known future transaction outside of Increase (only\nnegative amounts are supported). The resulting Pending Transaction will have a\n`category` of `user_initiated_hold` and can be released via the API to unlock\nthe held funds.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "The Account to place the hold on.",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The amount to hold in the minor unit of the account's currency. For dollars, for example, this is cents. This should be a negative amount - to hold $1.00 from the account, you would pass -100.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "The description you choose to give the hold.",
			BodyPath: "description",
		},
	},
	Action:          handlePendingTransactionsCreate,
	HideHelpCommand: true,
}

var pendingTransactionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Pending Transaction",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "pending-transaction-id",
			Usage:    "The identifier of the Pending Transaction.",
			Required: true,
		},
	},
	Action:          handlePendingTransactionsRetrieve,
	HideHelpCommand: true,
}

var pendingTransactionsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Pending Transactions",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter pending transactions to those belonging to the specified Account.",
			QueryPath: "account_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "category",
			QueryPath: "category",
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
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Limit the size of the list that is returned. The default (and maximum) is 100 objects.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "route-id",
			Usage:     "Filter pending transactions to those belonging to the specified Route.",
			QueryPath: "route_id",
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
	Action:          handlePendingTransactionsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"category": {
		&requestflag.InnerFlag[[]string]{
			Name:       "category.in",
			Usage:      "Return results whose value is in the provided list. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
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
			Usage:      "Filter Pending Transactions for those with the specified status. By default only Pending Transactions in with status `pending` will be returned. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

var pendingTransactionsRelease = cli.Command{
	Name:    "release",
	Usage:   "Release a Pending Transaction you had previously created. The Pending\nTransaction must have a `category` of `user_initiated_hold` and a `status` of\n`pending`. This will unlock the held funds and mark the Pending Transaction as\ncomplete.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "pending-transaction-id",
			Usage:    "The identifier of the Pending Transaction to release.",
			Required: true,
		},
	},
	Action:          handlePendingTransactionsRelease,
	HideHelpCommand: true,
}

func handlePendingTransactionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.PendingTransactionNewParams{}

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
	_, err = client.PendingTransactions.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "pending-transactions create", obj, format, explicitFormat, transform)
}

func handlePendingTransactionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pending-transaction-id") && len(unusedArgs) > 0 {
		cmd.Set("pending-transaction-id", unusedArgs[0])
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
	_, err = client.PendingTransactions.Get(ctx, cmd.Value("pending-transaction-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "pending-transactions retrieve", obj, format, explicitFormat, transform)
}

func handlePendingTransactionsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.PendingTransactionListParams{}

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
		_, err = client.PendingTransactions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "pending-transactions list", obj, format, explicitFormat, transform)
	} else {
		iter := client.PendingTransactions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "pending-transactions list", iter, format, explicitFormat, transform, maxItems)
	}
}

func handlePendingTransactionsRelease(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pending-transaction-id") && len(unusedArgs) > 0 {
		cmd.Set("pending-transaction-id", unusedArgs[0])
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
	_, err = client.PendingTransactions.Release(ctx, cmd.Value("pending-transaction-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "pending-transactions release", obj, format, explicitFormat, transform)
}
