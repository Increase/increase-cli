// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsCardReversalsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-reversals", "create",
			"--card-payment-id", "card_payment_nd3k2kacrqjli8482ave",
			"--amount", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"card_payment_id: card_payment_nd3k2kacrqjli8482ave\n" +
			"amount: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:card-reversals", "create",
		)
	})
}
