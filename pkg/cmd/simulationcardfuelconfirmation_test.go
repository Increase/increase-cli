// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsCardFuelConfirmationsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-fuel-confirmations", "create",
			"--amount", "5000",
			"--card-payment-id", "card_payment_nd3k2kacrqjli8482ave",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 5000\n" +
			"card_payment_id: card_payment_nd3k2kacrqjli8482ave\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:card-fuel-confirmations", "create",
		)
	})
}
