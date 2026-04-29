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

var inboundCheckDepositsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an Inbound Check Deposit",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbound-check-deposit-id",
			Usage:     "The identifier of the Inbound Check Deposit to get details for.",
			Required:  true,
			PathParam: "inbound_check_deposit_id",
		},
	},
	Action:          handleInboundCheckDepositsRetrieve,
	HideHelpCommand: true,
}

var inboundCheckDepositsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Inbound Check Deposits",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter Inbound Check Deposits to those belonging to the specified Account.",
			QueryPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:      "check-transfer-id",
			Usage:     "Filter Inbound Check Deposits to those belonging to the specified Check Transfer.",
			QueryPath: "check_transfer_id",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleInboundCheckDepositsList,
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
})

var inboundCheckDepositsDecline = cli.Command{
	Name:    "decline",
	Usage:   "Decline an Inbound Check Deposit",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbound-check-deposit-id",
			Usage:     "The identifier of the Inbound Check Deposit to decline.",
			Required:  true,
			PathParam: "inbound_check_deposit_id",
		},
	},
	Action:          handleInboundCheckDepositsDecline,
	HideHelpCommand: true,
}

var inboundCheckDepositsReturn = cli.Command{
	Name:    "return",
	Usage:   "Return an Inbound Check Deposit",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "inbound-check-deposit-id",
			Usage:     "The identifier of the Inbound Check Deposit to return.",
			Required:  true,
			PathParam: "inbound_check_deposit_id",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "The reason to return the Inbound Check Deposit.",
			Required: true,
			BodyPath: "reason",
		},
	},
	Action:          handleInboundCheckDepositsReturn,
	HideHelpCommand: true,
}

func handleInboundCheckDepositsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbound-check-deposit-id") && len(unusedArgs) > 0 {
		cmd.Set("inbound-check-deposit-id", unusedArgs[0])
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
	_, err = client.InboundCheckDeposits.Get(ctx, cmd.Value("inbound-check-deposit-id").(string), options...)
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "inbound-check-deposits retrieve",
		Transform:      transform,
	})
}

func handleInboundCheckDepositsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := increase.InboundCheckDepositListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.InboundCheckDeposits.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "inbound-check-deposits list",
			Transform:      transform,
		})
	} else {
		iter := client.InboundCheckDeposits.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "inbound-check-deposits list",
			Transform:      transform,
		})
	}
}

func handleInboundCheckDepositsDecline(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbound-check-deposit-id") && len(unusedArgs) > 0 {
		cmd.Set("inbound-check-deposit-id", unusedArgs[0])
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
	_, err = client.InboundCheckDeposits.Decline(ctx, cmd.Value("inbound-check-deposit-id").(string), options...)
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "inbound-check-deposits decline",
		Transform:      transform,
	})
}

func handleInboundCheckDepositsReturn(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("inbound-check-deposit-id") && len(unusedArgs) > 0 {
		cmd.Set("inbound-check-deposit-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := increase.InboundCheckDepositReturnParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.InboundCheckDeposits.Return(
		ctx,
		cmd.Value("inbound-check-deposit-id").(string),
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "inbound-check-deposits return",
		Transform:      transform,
	})
}
