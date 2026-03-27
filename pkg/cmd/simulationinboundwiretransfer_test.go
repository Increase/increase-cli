// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsInboundWireTransfersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:inbound-wire-transfers", "create",
			"--account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--amount", "1000",
			"--creditor-address-line1", "x",
			"--creditor-address-line2", "x",
			"--creditor-address-line3", "x",
			"--creditor-name", "x",
			"--debtor-address-line1", "x",
			"--debtor-address-line2", "x",
			"--debtor-address-line3", "x",
			"--debtor-name", "x",
			"--end-to-end-identification", "x",
			"--instructing-agent-routing-number", "x",
			"--instruction-identification", "x",
			"--unique-end-to-end-transaction-reference", "x",
			"--unstructured-remittance-information", "x",
			"--wire-drawdown-request-id", "wire_drawdown_request_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_number_id: account_number_v18nkfqm6afpsrvy82b2\n" +
			"amount: 1000\n" +
			"creditor_address_line1: x\n" +
			"creditor_address_line2: x\n" +
			"creditor_address_line3: x\n" +
			"creditor_name: x\n" +
			"debtor_address_line1: x\n" +
			"debtor_address_line2: x\n" +
			"debtor_address_line3: x\n" +
			"debtor_name: x\n" +
			"end_to_end_identification: x\n" +
			"instructing_agent_routing_number: x\n" +
			"instruction_identification: x\n" +
			"unique_end_to_end_transaction_reference: x\n" +
			"unstructured_remittance_information: x\n" +
			"wire_drawdown_request_id: wire_drawdown_request_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:inbound-wire-transfers", "create",
		)
	})
}
