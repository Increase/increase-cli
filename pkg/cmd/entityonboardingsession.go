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

var entityOnboardingSessionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an Entity Onboarding Session",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "program-id",
			Usage:    "The identifier of the Program the Entity will be onboarded to.",
			Required: true,
			BodyPath: "program_id",
		},
		&requestflag.Flag[string]{
			Name:     "redirect-url",
			Usage:    "The URL to redirect the customer to after they complete the onboarding form. The redirect will include `entity_onboarding_session_id` and `entity_id` query parameters.",
			Required: true,
			BodyPath: "redirect_url",
		},
		&requestflag.Flag[string]{
			Name:     "entity-id",
			Usage:    "The identifier of an existing Entity to associate with the onboarding session. If provided, the onboarding form will display any outstanding tasks required to complete the Entity's onboarding.",
			BodyPath: "entity_id",
		},
	},
	Action:          handleEntityOnboardingSessionsCreate,
	HideHelpCommand: true,
}

var entityOnboardingSessionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an Entity Onboarding Session",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-onboarding-session-id",
			Usage:    "The identifier of the Entity Onboarding Session.",
			Required: true,
		},
	},
	Action:          handleEntityOnboardingSessionsRetrieve,
	HideHelpCommand: true,
}

var entityOnboardingSessionsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Entity Onboarding Session",
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
	Action:          handleEntityOnboardingSessionsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"status": {
		&requestflag.InnerFlag[[]string]{
			Name:       "status.in",
			Usage:      "Filter Entity Onboarding Session for those with the specified status or statuses. For GET requests, this should be encoded as a comma-delimited string, such as `?in=one,two,three`.",
			InnerField: "in",
		},
	},
})

var entityOnboardingSessionsExpire = cli.Command{
	Name:    "expire",
	Usage:   "Expire an Entity Onboarding Session",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-onboarding-session-id",
			Usage:    "The identifier of the Entity Onboarding Session to expire.",
			Required: true,
		},
	},
	Action:          handleEntityOnboardingSessionsExpire,
	HideHelpCommand: true,
}

func handleEntityOnboardingSessionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.EntityOnboardingSessionNewParams{}

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
	_, err = client.EntityOnboardingSessions.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "entity-onboarding-sessions create", obj, format, transform)
}

func handleEntityOnboardingSessionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-onboarding-session-id") && len(unusedArgs) > 0 {
		cmd.Set("entity-onboarding-session-id", unusedArgs[0])
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
	_, err = client.EntityOnboardingSessions.Get(ctx, cmd.Value("entity-onboarding-session-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "entity-onboarding-sessions retrieve", obj, format, transform)
}

func handleEntityOnboardingSessionsList(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.EntityOnboardingSessionListParams{}

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
		_, err = client.EntityOnboardingSessions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "entity-onboarding-sessions list", obj, format, transform)
	} else {
		iter := client.EntityOnboardingSessions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "entity-onboarding-sessions list", iter, format, transform, maxItems)
	}
}

func handleEntityOnboardingSessionsExpire(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-onboarding-session-id") && len(unusedArgs) > 0 {
		cmd.Set("entity-onboarding-session-id", unusedArgs[0])
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
	_, err = client.EntityOnboardingSessions.Expire(ctx, cmd.Value("entity-onboarding-session-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "entity-onboarding-sessions expire", obj, format, transform)
}
