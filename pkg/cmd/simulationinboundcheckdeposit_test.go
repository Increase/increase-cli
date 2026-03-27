// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsInboundCheckDepositsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:inbound-check-deposits", "create",
			"--account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--amount", "1000",
			"--check-number", "1234567890",
			"--payee-name-analysis", "name_matches",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_number_id: account_number_v18nkfqm6afpsrvy82b2\n" +
			"amount: 1000\n" +
			"check_number: '1234567890'\n" +
			"payee_name_analysis: name_matches\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:inbound-check-deposits", "create",
		)
	})
}

func TestSimulationsInboundCheckDepositsAdjustment(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:inbound-check-deposits", "adjustment",
			"--inbound-check-deposit-id", "inbound_check_deposit_zoshvqybq0cjjm31mra",
			"--amount", "1000",
			"--reason", "late_return",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 1000\n" +
			"reason: late_return\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:inbound-check-deposits", "adjustment",
			"--inbound-check-deposit-id", "inbound_check_deposit_zoshvqybq0cjjm31mra",
		)
	})
}
