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

var externalAccountsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an External Account",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-number",
			Usage:    "The account number for the destination account.",
			Required: true,
			BodyPath: "account_number",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "The name you choose for the Account.",
			Required: true,
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "routing-number",
			Usage:    "The American Bankers' Association (ABA) Routing Transit Number (RTN) for the destination account.",
			Required: true,
			BodyPath: "routing_number",
		},
		&requestflag.Flag[string]{
			Name:     "account-holder",
			Usage:    "The type of entity that owns the External Account.",
			BodyPath: "account_holder",
		},
		&requestflag.Flag[string]{
			Name:     "funding",
			Usage:    "The type of the destination account. Defaults to `checking`.",
			BodyPath: "funding",
		},
	},
	Action:          handleExternalAccountsCreate,
	HideHelpCommand: true,
}

var externalAccountsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an External Account",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "external-account-id",
			Usage:    "The identifier of the External Account.",
			Required: true,
		},
	},
	Action:          handleExternalAccountsRetrieve,
	HideHelpCommand: true,
}

var externalAccountsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an External Account",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "external-account-id",
			Usage:    "The external account identifier.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "account-holder",
			Usage:    "The type of entity that owns the External Account.",
			BodyPath: "account_holder",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "The description you choose to give the external account.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "funding",
			Usage:    "The funding type of the External Account.",
			BodyPath: "funding",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The status of the External Account.",
			BodyPath: "status",
		},
	},
	Action:          handleExternalAccountsUpdate,
	HideHelpCommand: true,
}

var externalAccountsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List External Accounts",
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
			Name:      "routing-number",
			Usage:     "Filter External Accounts to those with the specified Routing Number.",
			QueryPath: "routing_number",
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
	Action:          handleExternalAccountsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"status": {
		&requestflag.InnerFlag[[]string]{
			Name:       "status.in",
			Usage:      "Filter External Accounts for those with the specified status or statuses. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

func handleExternalAccountsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.ExternalAccountNewParams{}

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
	_, err = client.ExternalAccounts.New(ctx, params, options...)
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
		Title:          "external-accounts create",
		Transform:      transform,
	})
}

func handleExternalAccountsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("external-account-id") && len(unusedArgs) > 0 {
		cmd.Set("external-account-id", unusedArgs[0])
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
	_, err = client.ExternalAccounts.Get(ctx, cmd.Value("external-account-id").(string), options...)
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
		Title:          "external-accounts retrieve",
		Transform:      transform,
	})
}

func handleExternalAccountsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("external-account-id") && len(unusedArgs) > 0 {
		cmd.Set("external-account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.ExternalAccountUpdateParams{}

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
	_, err = client.ExternalAccounts.Update(
		ctx,
		cmd.Value("external-account-id").(string),
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
		Title:          "external-accounts update",
		Transform:      transform,
	})
}

func handleExternalAccountsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.ExternalAccountListParams{}

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
		_, err = client.ExternalAccounts.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "external-accounts list",
			Transform:      transform,
		})
	} else {
		iter := client.ExternalAccounts.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "external-accounts list",
			Transform:      transform,
		})
	}
}
