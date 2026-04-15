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

var intrafiAccountEnrollmentsCreate = cli.Command{
	Name:    "create",
	Usage:   "Enroll an account in the IntraFi deposit sweep network",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "The identifier for the account to be added to IntraFi.",
			Required: true,
			BodyPath: "account_id",
		},
		&requestflag.Flag[string]{
			Name:     "email-address",
			Usage:    "The contact email for the account owner, to be shared with IntraFi.",
			Required: true,
			BodyPath: "email_address",
		},
	},
	Action:          handleIntrafiAccountEnrollmentsCreate,
	HideHelpCommand: true,
}

var intrafiAccountEnrollmentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get an IntraFi Account Enrollment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "intrafi-account-enrollment-id",
			Usage:    "The identifier of the IntraFi Account Enrollment to retrieve.",
			Required: true,
		},
	},
	Action:          handleIntrafiAccountEnrollmentsRetrieve,
	HideHelpCommand: true,
}

var intrafiAccountEnrollmentsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List IntraFi Account Enrollments",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "Filter IntraFi Account Enrollments to the one belonging to an account.",
			QueryPath: "account_id",
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
	Action:          handleIntrafiAccountEnrollmentsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"status": {
		&requestflag.InnerFlag[[]string]{
			Name:       "status.in",
			Usage:      "Filter IntraFi Account Enrollments for those with the specified status or statuses. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

var intrafiAccountEnrollmentsUnenroll = cli.Command{
	Name:    "unenroll",
	Usage:   "Unenroll an account from IntraFi",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "intrafi-account-enrollment-id",
			Usage:    "The Identifier of the IntraFi Account Enrollment to remove from IntraFi.",
			Required: true,
		},
	},
	Action:          handleIntrafiAccountEnrollmentsUnenroll,
	HideHelpCommand: true,
}

func handleIntrafiAccountEnrollmentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.IntrafiAccountEnrollmentNewParams{}

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
	_, err = client.IntrafiAccountEnrollments.New(ctx, params, options...)
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
		Title:          "intrafi-account-enrollments create",
		Transform:      transform,
	})
}

func handleIntrafiAccountEnrollmentsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("intrafi-account-enrollment-id") && len(unusedArgs) > 0 {
		cmd.Set("intrafi-account-enrollment-id", unusedArgs[0])
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
	_, err = client.IntrafiAccountEnrollments.Get(ctx, cmd.Value("intrafi-account-enrollment-id").(string), options...)
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
		Title:          "intrafi-account-enrollments retrieve",
		Transform:      transform,
	})
}

func handleIntrafiAccountEnrollmentsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.IntrafiAccountEnrollmentListParams{}

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
		_, err = client.IntrafiAccountEnrollments.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "intrafi-account-enrollments list",
			Transform:      transform,
		})
	} else {
		iter := client.IntrafiAccountEnrollments.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "intrafi-account-enrollments list",
			Transform:      transform,
		})
	}
}

func handleIntrafiAccountEnrollmentsUnenroll(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("intrafi-account-enrollment-id") && len(unusedArgs) > 0 {
		cmd.Set("intrafi-account-enrollment-id", unusedArgs[0])
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
	_, err = client.IntrafiAccountEnrollments.Unenroll(ctx, cmd.Value("intrafi-account-enrollment-id").(string), options...)
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
		Title:          "intrafi-account-enrollments unenroll",
		Transform:      transform,
	})
}
