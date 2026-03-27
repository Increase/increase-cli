// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsCardRefundsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-refunds", "create",
			"--amount", "1",
			"--pending-transaction-id", "pending_transaction_id",
			"--transaction-id", "transaction_uyrp7fld2ium70oa7oi",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 1\n" +
			"pending_transaction_id: pending_transaction_id\n" +
			"transaction_id: transaction_uyrp7fld2ium70oa7oi\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:card-refunds", "create",
		)
	})
}
