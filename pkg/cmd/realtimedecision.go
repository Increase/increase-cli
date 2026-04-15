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

var realTimeDecisionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Real-Time Decision",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "real-time-decision-id",
			Usage:    "The identifier of the Real-Time Decision.",
			Required: true,
		},
	},
	Action:          handleRealTimeDecisionsRetrieve,
	HideHelpCommand: true,
}

var realTimeDecisionsAction = requestflag.WithInnerFlags(cli.Command{
	Name:    "action",
	Usage:   "Action a Real-Time Decision",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "real-time-decision-id",
			Usage:    "The identifier of the Real-Time Decision.",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "card-authentication",
			Usage:    "If the Real-Time Decision relates to a 3DS card authentication attempt, this object contains your response to the authentication.",
			BodyPath: "card_authentication",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "card-authentication-challenge",
			Usage:    "If the Real-Time Decision relates to 3DS card authentication challenge delivery, this object contains your response.",
			BodyPath: "card_authentication_challenge",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "card-authorization",
			Usage:    "If the Real-Time Decision relates to a card authorization attempt, this object contains your response to the authorization.",
			BodyPath: "card_authorization",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "card-balance-inquiry",
			Usage:    "If the Real-Time Decision relates to a card balance inquiry attempt, this object contains your response to the inquiry.",
			BodyPath: "card_balance_inquiry",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "digital-wallet-authentication",
			Usage:    "If the Real-Time Decision relates to a digital wallet authentication attempt, this object contains your response to the authentication.",
			BodyPath: "digital_wallet_authentication",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "digital-wallet-token",
			Usage:    "If the Real-Time Decision relates to a digital wallet token provisioning attempt, this object contains your response to the attempt.",
			BodyPath: "digital_wallet_token",
		},
	},
	Action:          handleRealTimeDecisionsAction,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"card-authentication": {
		&requestflag.InnerFlag[string]{
			Name:       "card-authentication.decision",
			Usage:      "Whether the card authentication attempt should be approved or declined.",
			InnerField: "decision",
		},
	},
	"card-authentication-challenge": {
		&requestflag.InnerFlag[string]{
			Name:       "card-authentication-challenge.result",
			Usage:      "Whether the card authentication challenge was successfully delivered to the cardholder.",
			InnerField: "result",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "card-authentication-challenge.success",
			Usage:      "If your application was able to deliver the one-time code, this contains metadata about the delivery.",
			InnerField: "success",
		},
	},
	"card-authorization": {
		&requestflag.InnerFlag[string]{
			Name:       "card-authorization.decision",
			Usage:      "Whether the card authorization should be approved or declined.",
			InnerField: "decision",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "card-authorization.approval",
			Usage:      "If your application approves the authorization, this contains metadata about your decision to approve. Your response here is advisory to the acquiring bank. The bank may choose to reverse the authorization if you approve the transaction but indicate the address does not match.",
			InnerField: "approval",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "card-authorization.decline",
			Usage:      "If your application declines the authorization, this contains details about the decline.",
			InnerField: "decline",
		},
	},
	"card-balance-inquiry": {
		&requestflag.InnerFlag[string]{
			Name:       "card-balance-inquiry.decision",
			Usage:      "Whether the card balance inquiry should be approved or declined.",
			InnerField: "decision",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "card-balance-inquiry.approval",
			Usage:      "If your application approves the balance inquiry, this contains metadata about your decision to approve.",
			InnerField: "approval",
		},
	},
	"digital-wallet-authentication": {
		&requestflag.InnerFlag[string]{
			Name:       "digital-wallet-authentication.result",
			Usage:      "Whether your application was able to deliver the one-time passcode.",
			InnerField: "result",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "digital-wallet-authentication.success",
			Usage:      "If your application was able to deliver the one-time passcode, this contains metadata about the delivery. Exactly one of `phone` or `email` must be provided.",
			InnerField: "success",
		},
	},
	"digital-wallet-token": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "digital-wallet-token.approval",
			Usage:      "If your application approves the provisioning attempt, this contains metadata about the digital wallet token that will be generated.",
			InnerField: "approval",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "digital-wallet-token.decline",
			Usage:      "If your application declines the provisioning attempt, this contains details about the decline.",
			InnerField: "decline",
		},
	},
})

func handleRealTimeDecisionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("real-time-decision-id") && len(unusedArgs) > 0 {
		cmd.Set("real-time-decision-id", unusedArgs[0])
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
	_, err = client.RealTimeDecisions.Get(ctx, cmd.Value("real-time-decision-id").(string), options...)
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
		Title:          "real-time-decisions retrieve",
		Transform:      transform,
	})
}

func handleRealTimeDecisionsAction(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("real-time-decision-id") && len(unusedArgs) > 0 {
		cmd.Set("real-time-decision-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.RealTimeDecisionActionParams{}

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
	_, err = client.RealTimeDecisions.Action(
		ctx,
		cmd.Value("real-time-decision-id").(string),
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
		Title:          "real-time-decisions action",
		Transform:      transform,
	})
}
