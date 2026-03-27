// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsInterestPaymentsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:interest-payments", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--amount", "1000",
			"--accrued-on-account-id", "accrued_on_account_id",
			"--period-end", "'2019-12-27T18:11:19.117Z'",
			"--period-start", "'2019-12-27T18:11:19.117Z'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_in71c4amph0vgo2qllky\n" +
			"amount: 1000\n" +
			"accrued_on_account_id: accrued_on_account_id\n" +
			"period_end: '2019-12-27T18:11:19.117Z'\n" +
			"period_start: '2019-12-27T18:11:19.117Z'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:interest-payments", "create",
		)
	})
}
