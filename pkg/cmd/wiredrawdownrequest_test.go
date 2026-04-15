// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestWireDrawdownRequestsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wire-drawdown-requests", "create",
			"--account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--amount", "10000",
			"--creditor-address", "{city: New York, country: US, line1: 33 Liberty Street, line2: x, postal_code: '10045', state: NY}",
			"--creditor-name", "National Phonograph Company",
			"--debtor-address", "{city: New York, country: US, line1: 33 Liberty Street, line2: x, postal_code: '10045', state: NY}",
			"--debtor-name", "Ian Crease",
			"--unstructured-remittance-information", "Invoice 29582",
			"--charge-bearer", "shared",
			"--debtor-account-number", "987654321",
			"--debtor-external-account-id", "debtor_external_account_id",
			"--debtor-routing-number", "101050001",
			"--end-to-end-identification", "end_to_end_identification",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(wireDrawdownRequestsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wire-drawdown-requests", "create",
			"--account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--amount", "10000",
			"--creditor-address.city", "New York",
			"--creditor-address.country", "US",
			"--creditor-address.line1", "33 Liberty Street",
			"--creditor-address.line2", "x",
			"--creditor-address.postal-code", "10045",
			"--creditor-address.state", "NY",
			"--creditor-name", "National Phonograph Company",
			"--debtor-address.city", "New York",
			"--debtor-address.country", "US",
			"--debtor-address.line1", "33 Liberty Street",
			"--debtor-address.line2", "x",
			"--debtor-address.postal-code", "10045",
			"--debtor-address.state", "NY",
			"--debtor-name", "Ian Crease",
			"--unstructured-remittance-information", "Invoice 29582",
			"--charge-bearer", "shared",
			"--debtor-account-number", "987654321",
			"--debtor-external-account-id", "debtor_external_account_id",
			"--debtor-routing-number", "101050001",
			"--end-to-end-identification", "end_to_end_identification",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_number_id: account_number_v18nkfqm6afpsrvy82b2\n" +
			"amount: 10000\n" +
			"creditor_address:\n" +
			"  city: New York\n" +
			"  country: US\n" +
			"  line1: 33 Liberty Street\n" +
			"  line2: x\n" +
			"  postal_code: '10045'\n" +
			"  state: NY\n" +
			"creditor_name: National Phonograph Company\n" +
			"debtor_address:\n" +
			"  city: New York\n" +
			"  country: US\n" +
			"  line1: 33 Liberty Street\n" +
			"  line2: x\n" +
			"  postal_code: '10045'\n" +
			"  state: NY\n" +
			"debtor_name: Ian Crease\n" +
			"unstructured_remittance_information: Invoice 29582\n" +
			"charge_bearer: shared\n" +
			"debtor_account_number: '987654321'\n" +
			"debtor_external_account_id: debtor_external_account_id\n" +
			"debtor_routing_number: '101050001'\n" +
			"end_to_end_identification: end_to_end_identification\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"wire-drawdown-requests", "create",
		)
	})
}

func TestWireDrawdownRequestsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wire-drawdown-requests", "retrieve",
			"--wire-drawdown-request-id", "wire_drawdown_request_q6lmocus3glo0lr2bfv3",
		)
	})
}

func TestWireDrawdownRequestsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wire-drawdown-requests", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [pending_submission]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(wireDrawdownRequestsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"wire-drawdown-requests", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[pending_submission]",
		)
	})
}
