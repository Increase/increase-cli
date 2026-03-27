// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsCardAuthenticationsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-authentications", "create",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
			"--category", "payment_authentication",
			"--device-channel", "app",
			"--merchant-acceptor-id", "5665270011000168",
			"--merchant-category-code", "5734",
			"--merchant-country", "US",
			"--merchant-name", "x",
			"--purchase-amount", "1000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"card_id: card_oubs0hwk5rn6knuecxg2\n" +
			"category: payment_authentication\n" +
			"device_channel: app\n" +
			"merchant_acceptor_id: '5665270011000168'\n" +
			"merchant_category_code: '5734'\n" +
			"merchant_country: US\n" +
			"merchant_name: x\n" +
			"purchase_amount: 1000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:card-authentications", "create",
		)
	})
}

func TestSimulationsCardAuthenticationsChallengeAttempts(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-authentications", "challenge-attempts",
			"--card-payment-id", "card_payment_nd3k2kacrqjli8482ave",
			"--one-time-code", "123456",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("one_time_code: '123456'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:card-authentications", "challenge-attempts",
			"--card-payment-id", "card_payment_nd3k2kacrqjli8482ave",
		)
	})
}

func TestSimulationsCardAuthenticationsChallenges(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-authentications", "challenges",
			"--card-payment-id", "card_payment_nd3k2kacrqjli8482ave",
		)
	})
}
