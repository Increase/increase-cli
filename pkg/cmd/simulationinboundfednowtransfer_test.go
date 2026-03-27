// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsInboundFednowTransfersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:inbound-fednow-transfers", "create",
			"--account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--amount", "1000",
			"--debtor-account-number", "x",
			"--debtor-name", "x",
			"--debtor-routing-number", "xxxxxxxxx",
			"--unstructured-remittance-information", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_number_id: account_number_v18nkfqm6afpsrvy82b2\n" +
			"amount: 1000\n" +
			"debtor_account_number: x\n" +
			"debtor_name: x\n" +
			"debtor_routing_number: xxxxxxxxx\n" +
			"unstructured_remittance_information: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:inbound-fednow-transfers", "create",
		)
	})
}
