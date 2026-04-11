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

var exportsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create an Export",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "category",
			Usage:    "The type of Export to create.",
			Required: true,
			BodyPath: "category",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "account-statement-bai2",
			Usage:    "Options for the created export. Required if `category` is equal to `account_statement_bai2`.",
			BodyPath: "account_statement_bai2",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "account-statement-ofx",
			Usage:    "Options for the created export. Required if `category` is equal to `account_statement_ofx`.",
			BodyPath: "account_statement_ofx",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "account-verification-letter",
			Usage:    "Options for the created export. Required if `category` is equal to `account_verification_letter`.",
			BodyPath: "account_verification_letter",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "bookkeeping-account-balance-csv",
			Usage:    "Options for the created export. Required if `category` is equal to `bookkeeping_account_balance_csv`.",
			BodyPath: "bookkeeping_account_balance_csv",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "daily-account-balance-csv",
			Usage:    "Options for the created export. Required if `category` is equal to `daily_account_balance_csv`.",
			BodyPath: "daily_account_balance_csv",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "entity-csv",
			Usage:    "Options for the created export. Required if `category` is equal to `entity_csv`.",
			BodyPath: "entity_csv",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "funding-instructions",
			Usage:    "Options for the created export. Required if `category` is equal to `funding_instructions`.",
			BodyPath: "funding_instructions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "transaction-csv",
			Usage:    "Options for the created export. Required if `category` is equal to `transaction_csv`.",
			BodyPath: "transaction_csv",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "vendor-csv",
			Usage:    "Options for the created export. Required if `category` is equal to `vendor_csv`.",
			BodyPath: "vendor_csv",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "voided-check",
			Usage:    "Options for the created export. Required if `category` is equal to `voided_check`.",
			BodyPath: "voided_check",
		},
	},
	Action:          handleExportsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"account-statement-bai2": {
		&requestflag.InnerFlag[string]{
			Name:       "account-statement-bai2.account-id",
			Usage:      "The Account to create a BAI2 report for. If not provided, all open accounts will be included.",
			InnerField: "account_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "account-statement-bai2.effective-date",
			Usage:      "The date to create a BAI2 report for. If not provided, the current date will be used. The timezone is UTC. If the current date is used, the report will include intraday balances, otherwise it will include end-of-day balances for the provided date.",
			InnerField: "effective_date",
		},
		&requestflag.InnerFlag[string]{
			Name:       "account-statement-bai2.program-id",
			Usage:      "The Program to create a BAI2 report for. If not provided, all open accounts will be included.",
			InnerField: "program_id",
		},
	},
	"account-statement-ofx": {
		&requestflag.InnerFlag[string]{
			Name:       "account-statement-ofx.account-id",
			Usage:      "The Account to create a statement for.",
			InnerField: "account_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "account-statement-ofx.created-at",
			Usage:      "Filter transactions by their created date.",
			InnerField: "created_at",
		},
	},
	"account-verification-letter": {
		&requestflag.InnerFlag[string]{
			Name:       "account-verification-letter.account-number-id",
			Usage:      "The Account Number to create a letter for.",
			InnerField: "account_number_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "account-verification-letter.balance-date",
			Usage:      "The date of the balance to include in the letter. Defaults to the current date.",
			InnerField: "balance_date",
		},
	},
	"bookkeeping-account-balance-csv": {
		&requestflag.InnerFlag[string]{
			Name:       "bookkeeping-account-balance-csv.bookkeeping-account-id",
			Usage:      "Filter exported Bookkeeping Account Balances to the specified Bookkeeping Account.",
			InnerField: "bookkeeping_account_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "bookkeeping-account-balance-csv.on-or-after-date",
			Usage:      "Filter exported Balances to those on or after this date.",
			InnerField: "on_or_after_date",
		},
		&requestflag.InnerFlag[any]{
			Name:       "bookkeeping-account-balance-csv.on-or-before-date",
			Usage:      "Filter exported Balances to those on or before this date.",
			InnerField: "on_or_before_date",
		},
	},
	"daily-account-balance-csv": {
		&requestflag.InnerFlag[string]{
			Name:       "daily-account-balance-csv.account-id",
			Usage:      "Filter exported Balances to the specified Account.",
			InnerField: "account_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "daily-account-balance-csv.on-or-after-date",
			Usage:      "Filter exported Balances to those on or after this date.",
			InnerField: "on_or_after_date",
		},
		&requestflag.InnerFlag[any]{
			Name:       "daily-account-balance-csv.on-or-before-date",
			Usage:      "Filter exported Balances to those on or before this date.",
			InnerField: "on_or_before_date",
		},
	},
	"funding-instructions": {
		&requestflag.InnerFlag[string]{
			Name:       "funding-instructions.account-number-id",
			Usage:      "The Account Number to create funding instructions for.",
			InnerField: "account_number_id",
		},
	},
	"transaction-csv": {
		&requestflag.InnerFlag[string]{
			Name:       "transaction-csv.account-id",
			Usage:      "Filter exported Transactions to the specified Account.",
			InnerField: "account_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "transaction-csv.created-at",
			Usage:      "Filter results by time range on the `created_at` attribute.",
			InnerField: "created_at",
		},
	},
	"voided-check": {
		&requestflag.InnerFlag[string]{
			Name:       "voided-check.account-number-id",
			Usage:      "The Account Number for the voided check.",
			InnerField: "account_number_id",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "voided-check.payer",
			Usage:      "The payer information to be printed on the check.",
			InnerField: "payer",
		},
	},
})

var exportsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an Export",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "export-id",
			Usage:    "The identifier of the Export to retrieve.",
			Required: true,
		},
	},
	Action:          handleExportsRetrieve,
	HideHelpCommand: true,
}

var exportsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Exports",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "category",
			Usage:     "Filter Exports for those with the specified category.",
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
		&requestflag.Flag[map[string]any]{
			Name:      "form-1099-int",
			QueryPath: "form_1099_int",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "form-1099-misc",
			QueryPath: "form_1099_misc",
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
	Action:          handleExportsList,
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
	"form-1099-int": {
		&requestflag.InnerFlag[string]{
			Name:       "form-1099-int.account-id",
			Usage:      "Filter Form 1099-INT Exports to those for the specified Account.",
			InnerField: "account_id",
		},
	},
	"form-1099-misc": {
		&requestflag.InnerFlag[string]{
			Name:       "form-1099-misc.account-id",
			Usage:      "Filter Form 1099-MISC Exports to those for the specified Account.",
			InnerField: "account_id",
		},
	},
	"status": {
		&requestflag.InnerFlag[[]string]{
			Name:       "status.in",
			Usage:      "Filter Exports for those with the specified status or statuses. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

func handleExportsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.ExportNewParams{}

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
	_, err = client.Exports.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "exports create", obj, format, transform)
}

func handleExportsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("export-id") && len(unusedArgs) > 0 {
		cmd.Set("export-id", unusedArgs[0])
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
	_, err = client.Exports.Get(ctx, cmd.Value("export-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "exports retrieve", obj, format, transform)
}

func handleExportsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.ExportListParams{}

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
		_, err = client.Exports.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "exports list", obj, format, transform)
	} else {
		iter := client.Exports.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "exports list", iter, format, transform, maxItems)
	}
}
