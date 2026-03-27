// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsInboundWireDrawdownRequestsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:inbound-wire-drawdown-requests", "create",
			"--amount", "10000",
			"--creditor-account-number", "987654321",
			"--creditor-routing-number", "101050001",
			"--currency", "USD",
			"--recipient-account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--creditor-address-line1", "33 Liberty Street",
			"--creditor-address-line2", "New York, NY, 10045",
			"--creditor-address-line3", "x",
			"--creditor-name", "Ian Crease",
			"--debtor-account-number", "987654321",
			"--debtor-address-line1", "33 Liberty Street",
			"--debtor-address-line2", "New York, NY, 10045",
			"--debtor-address-line3", "x",
			"--debtor-name", "Ian Crease",
			"--debtor-routing-number", "101050001",
			"--end-to-end-identification", "x",
			"--instruction-identification", "x",
			"--unique-end-to-end-transaction-reference", "x",
			"--unstructured-remittance-information", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 10000\n" +
			"creditor_account_number: '987654321'\n" +
			"creditor_routing_number: '101050001'\n" +
			"currency: USD\n" +
			"recipient_account_number_id: account_number_v18nkfqm6afpsrvy82b2\n" +
			"creditor_address_line1: 33 Liberty Street\n" +
			"creditor_address_line2: New York, NY, 10045\n" +
			"creditor_address_line3: x\n" +
			"creditor_name: Ian Crease\n" +
			"debtor_account_number: '987654321'\n" +
			"debtor_address_line1: 33 Liberty Street\n" +
			"debtor_address_line2: New York, NY, 10045\n" +
			"debtor_address_line3: x\n" +
			"debtor_name: Ian Crease\n" +
			"debtor_routing_number: '101050001'\n" +
			"end_to_end_identification: x\n" +
			"instruction_identification: x\n" +
			"unique_end_to_end_transaction_reference: x\n" +
			"unstructured_remittance_information: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:inbound-wire-drawdown-requests", "create",
		)
	})
}
