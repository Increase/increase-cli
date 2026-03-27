// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestIntrafiBalancesIntrafiBalance(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"intrafi-balances", "intrafi-balance",
			"--account-id", "account_in71c4amph0vgo2qllky",
		)
	})
}
