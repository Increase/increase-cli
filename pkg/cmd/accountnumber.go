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

var accountNumbersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create an Account Number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "The Account the Account Number should belong to.",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name you choose for the Account Number.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound-ach",
			Usage:    "Options related to how this Account Number should handle inbound ACH transfers.",
			BodyPath: "inbound_ach",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound-checks",
			Usage:    "Options related to how this Account Number should handle inbound check withdrawals.",
			BodyPath: "inbound_checks",
		},
	},
	Action:          handleAccountNumbersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"inbound-ach": {
		&requestflag.InnerFlag[string]{
			Name:       "inbound-ach.debit-status",
			Usage:      "Whether ACH debits are allowed against this Account Number. Note that ACH debits will be declined if this is `allowed` but the Account Number is not active. If you do not specify this field, the default is `allowed`.",
			InnerField: "debit_status",
		},
	},
	"inbound-checks": {
		&requestflag.InnerFlag[string]{
			Name:       "inbound-checks.status",
			Usage:      "How Increase should process checks with this account number printed on them. If you do not specify this field, the default is `check_transfers_only`.",
			InnerField: "status",
		},
	},
})

var accountNumbersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an Account Number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-number-id",
			Usage:    "The identifier of the Account Number to retrieve.",
			Required: true,
		},
	},
	Action:          handleAccountNumbersRetrieve,
	HideHelpCommand: true,
}

var accountNumbersUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an Account Number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-number-id",
			Usage:    "The identifier of the Account Number.",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound-ach",
			Usage:    "Options related to how this Account Number handles inbound ACH transfers.",
			BodyPath: "inbound_ach",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound-checks",
			Usage:    "Options related to how this Account Number should handle inbound check withdrawals.",
			BodyPath: "inbound_checks",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name you choose for the Account Number.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "This indicates if transfers can be made to the Account Number.",
			BodyPath: "status",
		},
	},
	Action:          handleAccountNumbersUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"inbound-ach": {
		&requestflag.InnerFlag[string]{
			Name:       "inbound-ach.debit-status",
			Usage:      "Whether ACH debits are allowed against this Account Number. Note that ACH debits will be declined if this is `allowed` but the Account Number is not active.",
			InnerField: "debit_status",
		},
	},
	"inbound-checks": {
		&requestflag.InnerFlag[string]{
			Name:       "inbound-checks.status",
			Usage:      "How Increase should process checks with this account number printed on them.",
			InnerField: "status",
		},
	},
})

var accountNumbersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Account Numbers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter Account Numbers to those belonging to the specified Account.",
			QueryPath: "account_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "ach-debit-status",
			QueryPath: "ach_debit_status",
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
	Action:          handleAccountNumbersList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"ach-debit-status": {
		&requestflag.InnerFlag[[]string]{
			Name:       "ach-debit-status.in",
			Usage:      "The ACH Debit status to retrieve Account Numbers for. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
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
			Usage:      "The status to retrieve Account Numbers for. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

func handleAccountNumbersCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.AccountNumberNewParams{}

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
	_, err = client.AccountNumbers.New(ctx, params, options...)
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
		Title:          "account-numbers create",
		Transform:      transform,
	})
}

func handleAccountNumbersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-number-id") && len(unusedArgs) > 0 {
		cmd.Set("account-number-id", unusedArgs[0])
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
	_, err = client.AccountNumbers.Get(ctx, cmd.Value("account-number-id").(string), options...)
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
		Title:          "account-numbers retrieve",
		Transform:      transform,
	})
}

func handleAccountNumbersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-number-id") && len(unusedArgs) > 0 {
		cmd.Set("account-number-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.AccountNumberUpdateParams{}

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
	_, err = client.AccountNumbers.Update(
		ctx,
		cmd.Value("account-number-id").(string),
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
		Title:          "account-numbers update",
		Transform:      transform,
	})
}

func handleAccountNumbersList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.AccountNumberListParams{}

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
		_, err = client.AccountNumbers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "account-numbers list",
			Transform:      transform,
		})
	} else {
		iter := client.AccountNumbers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "account-numbers list",
			Transform:      transform,
		})
	}
}
