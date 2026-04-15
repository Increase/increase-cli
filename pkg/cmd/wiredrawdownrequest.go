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

var wireDrawdownRequestsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Wire Drawdown Request",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-number-id",
			Usage:    "The Account Number to which the debtor should send funds.",
			Required: true,
			BodyPath: "account_number_id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The amount requested from the debtor, in USD cents.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "creditor-address",
			Usage:    "The creditor's address.",
			Required: true,
			BodyPath: "creditor_address",
		},
		&requestflag.Flag[string]{
			Name:     "creditor-name",
			Usage:    "The creditor's name.",
			Required: true,
			BodyPath: "creditor_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "debtor-address",
			Usage:    "The debtor's address.",
			Required: true,
			BodyPath: "debtor_address",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-name",
			Usage:    "The debtor's name.",
			Required: true,
			BodyPath: "debtor_name",
		},
		&requestflag.Flag[string]{
			Name:     "unstructured-remittance-information",
			Usage:    "Remittance information the debtor will see as part of the request.",
			Required: true,
			BodyPath: "unstructured_remittance_information",
		},
		&requestflag.Flag[string]{
			Name:     "charge-bearer",
			Usage:    "Determines who bears the cost of the drawdown request. Defaults to `shared` if not specified.",
			BodyPath: "charge_bearer",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-account-number",
			Usage:    "The debtor's account number.",
			BodyPath: "debtor_account_number",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-external-account-id",
			Usage:    "The ID of an External Account to initiate a transfer to. If this parameter is provided, `debtor_account_number` and `debtor_routing_number` must be absent.",
			BodyPath: "debtor_external_account_id",
		},
		&requestflag.Flag[string]{
			Name:     "debtor-routing-number",
			Usage:    "The debtor's routing number.",
			BodyPath: "debtor_routing_number",
		},
		&requestflag.Flag[string]{
			Name:     "end-to-end-identification",
			Usage:    "A free-form reference string set by the sender mirrored back in the subsequent wire transfer.",
			BodyPath: "end_to_end_identification",
		},
	},
	Action:          handleWireDrawdownRequestsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"creditor-address": {
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.city",
			Usage:      "The city, district, town, or village of the address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.country",
			Usage:      "The two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code for the country of the address.",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.line1",
			Usage:      "The first line of the address. This is usually the street number and street.",
			InnerField: "line1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.line2",
			Usage:      "The second line of the address. This might be the floor or room number.",
			InnerField: "line2",
		},
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.postal-code",
			Usage:      "The ZIP code of the address.",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "creditor-address.state",
			Usage:      "The address state.",
			InnerField: "state",
		},
	},
	"debtor-address": {
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.city",
			Usage:      "The city, district, town, or village of the address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.country",
			Usage:      "The two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code for the country of the address.",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.line1",
			Usage:      "The first line of the address. This is usually the street number and street.",
			InnerField: "line1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.line2",
			Usage:      "The second line of the address. This might be the floor or room number.",
			InnerField: "line2",
		},
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.postal-code",
			Usage:      "The ZIP code of the address.",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "debtor-address.state",
			Usage:      "The address state.",
			InnerField: "state",
		},
	},
})

var wireDrawdownRequestsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Wire Drawdown Request",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "wire-drawdown-request-id",
			Usage:    "The identifier of the Wire Drawdown Request to retrieve.",
			Required: true,
		},
	},
	Action:          handleWireDrawdownRequestsRetrieve,
	HideHelpCommand: true,
}

var wireDrawdownRequestsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Wire Drawdown Requests",
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
	Action:          handleWireDrawdownRequestsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"status": {
		&requestflag.InnerFlag[[]string]{
			Name:       "status.in",
			Usage:      "Filter Wire Drawdown Requests for those with the specified status. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

func handleWireDrawdownRequestsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.WireDrawdownRequestNewParams{}

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
	_, err = client.WireDrawdownRequests.New(ctx, params, options...)
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
		Title:          "wire-drawdown-requests create",
		Transform:      transform,
	})
}

func handleWireDrawdownRequestsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("wire-drawdown-request-id") && len(unusedArgs) > 0 {
		cmd.Set("wire-drawdown-request-id", unusedArgs[0])
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
	_, err = client.WireDrawdownRequests.Get(ctx, cmd.Value("wire-drawdown-request-id").(string), options...)
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
		Title:          "wire-drawdown-requests retrieve",
		Transform:      transform,
	})
}

func handleWireDrawdownRequestsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.WireDrawdownRequestListParams{}

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
		_, err = client.WireDrawdownRequests.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "wire-drawdown-requests list",
			Transform:      transform,
		})
	} else {
		iter := client.WireDrawdownRequests.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "wire-drawdown-requests list",
			Transform:      transform,
		})
	}
}
