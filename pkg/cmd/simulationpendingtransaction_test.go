// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsPendingTransactionsReleaseInboundFundsHold(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:pending-transactions", "release-inbound-funds-hold",
			"--pending-transaction-id", "pending_transaction_k1sfetcau2qbvjbzgju4",
		)
	})
}
