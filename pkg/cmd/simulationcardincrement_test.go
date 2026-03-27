// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsCardIncrementsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-increments", "create",
			"--amount", "500",
			"--card-payment-id", "card_payment_nd3k2kacrqjli8482ave",
			"--event-subscription-id", "event_subscription_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 500\n" +
			"card_payment_id: card_payment_nd3k2kacrqjli8482ave\n" +
			"event_subscription_id: event_subscription_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:card-increments", "create",
		)
	})
}
