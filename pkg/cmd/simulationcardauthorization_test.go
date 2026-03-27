// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestSimulationsCardAuthorizationsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-authorizations", "create",
			"--amount", "1000",
			"--authenticated-card-payment-id", "authenticated_card_payment_id",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
			"--decline-reason", "account_closed",
			"--digital-wallet-token-id", "digital_wallet_token_id",
			"--event-subscription-id", "event_subscription_001dzz0r20rcdxgb013zqb8m04g",
			"--merchant-acceptor-id", "5665270011000168",
			"--merchant-category-code", "5734",
			"--merchant-city", "New York",
			"--merchant-country", "US",
			"--merchant-descriptor", "AMAZON.COM",
			"--merchant-state", "NY",
			"--network-details", "{visa: {stand_in_processing_reason: issuer_error}}",
			"--network-risk-score", "0",
			"--physical-card-id", "physical_card_id",
			"--processing-category", "{category: account_funding, refund: {original_card_payment_id: original_card_payment_id}}",
			"--terminal-id", "x",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simulationsCardAuthorizationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-authorizations", "create",
			"--amount", "1000",
			"--authenticated-card-payment-id", "authenticated_card_payment_id",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
			"--decline-reason", "account_closed",
			"--digital-wallet-token-id", "digital_wallet_token_id",
			"--event-subscription-id", "event_subscription_001dzz0r20rcdxgb013zqb8m04g",
			"--merchant-acceptor-id", "5665270011000168",
			"--merchant-category-code", "5734",
			"--merchant-city", "New York",
			"--merchant-country", "US",
			"--merchant-descriptor", "AMAZON.COM",
			"--merchant-state", "NY",
			"--network-details.visa", "{stand_in_processing_reason: issuer_error}",
			"--network-risk-score", "0",
			"--physical-card-id", "physical_card_id",
			"--processing-category.category", "account_funding",
			"--processing-category.refund", "{original_card_payment_id: original_card_payment_id}",
			"--terminal-id", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 1000\n" +
			"authenticated_card_payment_id: authenticated_card_payment_id\n" +
			"card_id: card_oubs0hwk5rn6knuecxg2\n" +
			"decline_reason: account_closed\n" +
			"digital_wallet_token_id: digital_wallet_token_id\n" +
			"event_subscription_id: event_subscription_001dzz0r20rcdxgb013zqb8m04g\n" +
			"merchant_acceptor_id: '5665270011000168'\n" +
			"merchant_category_code: '5734'\n" +
			"merchant_city: New York\n" +
			"merchant_country: US\n" +
			"merchant_descriptor: AMAZON.COM\n" +
			"merchant_state: NY\n" +
			"network_details:\n" +
			"  visa:\n" +
			"    stand_in_processing_reason: issuer_error\n" +
			"network_risk_score: 0\n" +
			"physical_card_id: physical_card_id\n" +
			"processing_category:\n" +
			"  category: account_funding\n" +
			"  refund:\n" +
			"    original_card_payment_id: original_card_payment_id\n" +
			"terminal_id: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:card-authorizations", "create",
		)
	})
}
