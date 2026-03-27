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

var lockboxesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a Lockbox",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "The Account checks sent to this Lockbox should be deposited into.",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "The description you choose for the Lockbox, for display purposes.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-name",
			Usage:    "The name of the recipient that will receive mail at this location.",
			BodyPath: "recipient_name",
		},
	},
	Action:          handleLockboxesCreate,
	HideHelpCommand: true,
}

var lockboxesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Lockbox",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "lockbox-id",
			Usage:    "The identifier of the Lockbox to retrieve.",
			Required: true,
		},
	},
	Action:          handleLockboxesRetrieve,
	HideHelpCommand: true,
}

var lockboxesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a Lockbox",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "lockbox-id",
			Usage:    "The identifier of the Lockbox.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "check-deposit-behavior",
			Usage:    "This indicates if checks mailed to this lockbox will be deposited.",
			BodyPath: "check_deposit_behavior",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "The description you choose for the Lockbox.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-name",
			Usage:    "The recipient name you choose for the Lockbox.",
			BodyPath: "recipient_name",
		},
	},
	Action:          handleLockboxesUpdate,
	HideHelpCommand: true,
}

var lockboxesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Lockboxes",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter Lockboxes to those associated with the provided Account.",
			QueryPath: "account_id",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleLockboxesList,
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

func handleLockboxesCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.LockboxNewParams{}

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
	_, err = client.Lockboxes.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "lockboxes create", obj, format, transform)
}

func handleLockboxesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("lockbox-id") && len(unusedArgs) > 0 {
		cmd.Set("lockbox-id", unusedArgs[0])
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
	_, err = client.Lockboxes.Get(ctx, cmd.Value("lockbox-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "lockboxes retrieve", obj, format, transform)
}

func handleLockboxesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("lockbox-id") && len(unusedArgs) > 0 {
		cmd.Set("lockbox-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.LockboxUpdateParams{}

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
	_, err = client.Lockboxes.Update(
		ctx,
		cmd.Value("lockbox-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "lockboxes update", obj, format, transform)
}

func handleLockboxesList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.LockboxListParams{}

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
		_, err = client.Lockboxes.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "lockboxes list", obj, format, transform)
	} else {
		iter := client.Lockboxes.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "lockboxes list", iter, format, transform, maxItems)
	}
}
