// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsProgramsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:programs", "create",
			"--name", "For Benefit Of",
			"--bank", "core_bank",
			"--lending-maximum-extendable-credit", "0",
			"--reserve-account-id", "reserve_account_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: For Benefit Of\n" +
			"bank: core_bank\n" +
			"lending_maximum_extendable_credit: 0\n" +
			"reserve_account_id: reserve_account_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:programs", "create",
		)
	})
}
