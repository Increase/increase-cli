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

var achPrenotificationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an ACH Prenotification",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "The Increase identifier for the account that will send the ACH Prenotification.",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:     "account-number",
			Usage:    "The account number for the destination account.",
			Required: true,
			BodyPath: "account_number",
		},
		&requestflag.Flag[string]{
			Name:     "routing-number",
			Usage:    "The American Bankers' Association (ABA) Routing Transit Number (RTN) for the destination account.",
			Required: true,
			BodyPath: "routing_number",
		},
		&requestflag.Flag[string]{
			Name:     "addendum",
			Usage:    "Additional information that will be sent to the recipient.",
			BodyPath: "addendum",
		},
		&requestflag.Flag[string]{
			Name:     "company-descriptive-date",
			Usage:    "The description of the date of the ACH Prenotification.",
			BodyPath: "company_descriptive_date",
		},
		&requestflag.Flag[string]{
			Name:     "company-discretionary-data",
			Usage:    "The data you choose to associate with the ACH Prenotification.",
			BodyPath: "company_discretionary_data",
		},
		&requestflag.Flag[string]{
			Name:     "company-entry-description",
			Usage:    "The description you wish to be shown to the recipient.",
			BodyPath: "company_entry_description",
		},
		&requestflag.Flag[string]{
			Name:     "company-name",
			Usage:    "The name by which the recipient knows you.",
			BodyPath: "company_name",
		},
		&requestflag.Flag[string]{
			Name:     "credit-debit-indicator",
			Usage:    "Whether the Prenotification is for a future debit or credit.",
			BodyPath: "credit_debit_indicator",
		},
		&requestflag.Flag[any]{
			Name:     "effective-date",
			Usage:    "The ACH Prenotification effective date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.",
			BodyPath: "effective_date",
		},
		&requestflag.Flag[string]{
			Name:     "individual-id",
			Usage:    "Your identifier for the recipient.",
			BodyPath: "individual_id",
		},
		&requestflag.Flag[string]{
			Name:     "individual-name",
			Usage:    "The name of therecipient. This value is informational and not verified by the recipient's bank.",
			BodyPath: "individual_name",
		},
		&requestflag.Flag[string]{
			Name:     "standard-entry-class-code",
			Usage:    "The [Standard Entry Class (SEC) code](/documentation/ach-standard-entry-class-codes) to use for the ACH Prenotification.",
			BodyPath: "standard_entry_class_code",
		},
	},
	Action:          handleACHPrenotificationsCreate,
	HideHelpCommand: true,
}

var achPrenotificationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an ACH Prenotification",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "ach-prenotification-id",
			Usage:    "The identifier of the ACH Prenotification to retrieve.",
			Required: true,
		},
	},
	Action:          handleACHPrenotificationsRetrieve,
	HideHelpCommand: true,
}

var achPrenotificationsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List ACH Prenotifications",
	Suggest: true,
	Flags: []cli.Flag{
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
	Action:          handleACHPrenotificationsList,
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

func handleACHPrenotificationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.ACHPrenotificationNewParams{}

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
	_, err = client.ACHPrenotifications.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ach-prenotifications create", obj, format, transform)
}

func handleACHPrenotificationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ach-prenotification-id") && len(unusedArgs) > 0 {
		cmd.Set("ach-prenotification-id", unusedArgs[0])
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
	_, err = client.ACHPrenotifications.Get(ctx, cmd.Value("ach-prenotification-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ach-prenotifications retrieve", obj, format, transform)
}

func handleACHPrenotificationsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.ACHPrenotificationListParams{}

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
		_, err = client.ACHPrenotifications.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "ach-prenotifications list", obj, format, transform)
	} else {
		iter := client.ACHPrenotifications.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "ach-prenotifications list", iter, format, transform, maxItems)
	}
}
