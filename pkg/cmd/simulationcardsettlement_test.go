// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsCardSettlementsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-settlements", "create",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
			"--pending-transaction-id", "pending_transaction_k1sfetcau2qbvjbzgju4",
			"--amount", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"card_id: card_oubs0hwk5rn6knuecxg2\n" +
			"pending_transaction_id: pending_transaction_k1sfetcau2qbvjbzgju4\n" +
			"amount: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:card-settlements", "create",
		)
	})
}
