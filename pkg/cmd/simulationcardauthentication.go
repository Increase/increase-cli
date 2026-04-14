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

var simulationsCardAuthenticationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Simulates a Card Authentication attempt on a [Card](#cards). The attempt always\nresults in a [Card Payment](#card_payments) being created, either with a status\nthat allows further action or a terminal failed status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-id",
			Usage:    "The identifier of the Card to be authorized.",
			Required: true,
			BodyPath: "card_id",
		},
		&requestflag.Flag[string]{
			Name:     "category",
			Usage:    "The category of the card authentication attempt.",
			BodyPath: "category",
		},
		&requestflag.Flag[string]{
			Name:     "device-channel",
			Usage:    "The device channel of the card authentication attempt.",
			BodyPath: "device_channel",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-acceptor-id",
			Usage:    "The merchant identifier (commonly abbreviated as MID) of the merchant the card is transacting with.",
			BodyPath: "merchant_acceptor_id",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-category-code",
			Usage:    "The Merchant Category Code (commonly abbreviated as MCC) of the merchant the card is transacting with.",
			BodyPath: "merchant_category_code",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-country",
			Usage:    "The country the merchant resides in.",
			BodyPath: "merchant_country",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-name",
			Usage:    "The name of the merchant",
			BodyPath: "merchant_name",
		},
		&requestflag.Flag[int64]{
			Name:     "purchase-amount",
			Usage:    "The purchase amount in cents.",
			BodyPath: "purchase_amount",
		},
	},
	Action:          handleSimulationsCardAuthenticationsCreate,
	HideHelpCommand: true,
}

var simulationsCardAuthenticationsChallengeAttempts = cli.Command{
	Name:    "challenge-attempts",
	Usage:   "Simulates an attempt at a Card Authentication Challenge. This updates the\n`card_authentications` object under the [Card Payment](#card_payments). You can\nalso attempt the challenge by navigating to\nhttps://dashboard.increase.com/card_authentication_simulation/:card_payment_id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-payment-id",
			Usage:    "The identifier of the Card Payment to be challenged.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "one-time-code",
			Usage:    "The one-time code to be validated.",
			Required: true,
			BodyPath: "one_time_code",
		},
	},
	Action:          handleSimulationsCardAuthenticationsChallengeAttempts,
	HideHelpCommand: true,
}

var simulationsCardAuthenticationsChallenges = cli.Command{
	Name:    "challenges",
	Usage:   "Simulates starting a Card Authentication Challenge for an existing Card\nAuthentication. This updates the `card_authentications` object under the\n[Card Payment](#card_payments). To attempt the challenge, use the\n`/simulations/card_authentications/:card_payment_id/challenge_attempts` endpoint\nor navigate to\nhttps://dashboard.increase.com/card_authentication_simulation/:card_payment_id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "card-payment-id",
			Usage:    "The identifier of the Card Payment to be challenged.",
			Required: true,
		},
	},
	Action:          handleSimulationsCardAuthenticationsChallenges,
	HideHelpCommand: true,
}

func handleSimulationsCardAuthenticationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationCardAuthenticationNewParams{}

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
	_, err = client.Simulations.CardAuthentications.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "simulations:card-authentications create", obj, format, explicitFormat, transform)
}

func handleSimulationsCardAuthenticationsChallengeAttempts(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-payment-id") && len(unusedArgs) > 0 {
		cmd.Set("card-payment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationCardAuthenticationChallengeAttemptsParams{}

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
	_, err = client.Simulations.CardAuthentications.ChallengeAttempts(
		ctx,
		cmd.Value("card-payment-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "simulations:card-authentications challenge-attempts", obj, format, explicitFormat, transform)
}

func handleSimulationsCardAuthenticationsChallenges(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-payment-id") && len(unusedArgs) > 0 {
		cmd.Set("card-payment-id", unusedArgs[0])
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
	_, err = client.Simulations.CardAuthentications.Challenges(ctx, cmd.Value("card-payment-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "simulations:card-authentications challenges", obj, format, explicitFormat, transform)
}
