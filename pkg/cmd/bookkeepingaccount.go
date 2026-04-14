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

var bookkeepingAccountsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a Bookkeeping Account",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name you choose for the account.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "The account, if `compliance_category` is `commingled_cash`.",
			BodyPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:     "compliance-category",
			Usage:    "The account compliance category.",
			BodyPath: "compliance_category",
		},
		&requestflag.Flag[string]{
			Name:     "entity-id",
			Usage:    "The entity, if `compliance_category` is `customer_balance`.",
			BodyPath: "entity_id",
		},
	},
	Action:          handleBookkeepingAccountsCreate,
	HideHelpCommand: true,
}

var bookkeepingAccountsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a Bookkeeping Account",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bookkeeping-account-id",
			Usage:    "The bookkeeping account you would like to update.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name you choose for the account.",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handleBookkeepingAccountsUpdate,
	HideHelpCommand: true,
}

var bookkeepingAccountsList = cli.Command{
	Name:    "list",
	Usage:   "List Bookkeeping Accounts",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBookkeepingAccountsList,
	HideHelpCommand: true,
}

var bookkeepingAccountsBalance = cli.Command{
	Name:    "balance",
	Usage:   "Retrieve a Bookkeeping Account Balance",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bookkeeping-account-id",
			Usage:    "The identifier of the Bookkeeping Account to retrieve.",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "at-time",
			Usage:     "The moment to query the balance at. If not set, returns the current balances.",
			QueryPath: "at_time",
		},
	},
	Action:          handleBookkeepingAccountsBalance,
	HideHelpCommand: true,
}

func handleBookkeepingAccountsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.BookkeepingAccountNewParams{}

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
	_, err = client.BookkeepingAccounts.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "bookkeeping-accounts create", obj, format, explicitFormat, transform)
}

func handleBookkeepingAccountsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("bookkeeping-account-id") && len(unusedArgs) > 0 {
		cmd.Set("bookkeeping-account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.BookkeepingAccountUpdateParams{}

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
	_, err = client.BookkeepingAccounts.Update(
		ctx,
		cmd.Value("bookkeeping-account-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "bookkeeping-accounts update", obj, format, explicitFormat, transform)
}

func handleBookkeepingAccountsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.BookkeepingAccountListParams{}

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
		_, err = client.BookkeepingAccounts.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "bookkeeping-accounts list", obj, format, explicitFormat, transform)
	} else {
		iter := client.BookkeepingAccounts.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "bookkeeping-accounts list", iter, format, explicitFormat, transform, maxItems)
	}
}

func handleBookkeepingAccountsBalance(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("bookkeeping-account-id") && len(unusedArgs) > 0 {
		cmd.Set("bookkeeping-account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.BookkeepingAccountBalanceParams{}

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
	_, err = client.BookkeepingAccounts.Balance(
		ctx,
		cmd.Value("bookkeeping-account-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "bookkeeping-accounts balance", obj, format, explicitFormat, transform)
}
