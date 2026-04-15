// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestSwiftTransfersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"swift-transfers", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--account-number", "987654321",
			"--bank-identification-code", "ECBFDEFFTPP",
			"--creditor-address", "{city: Frankfurt, country: DE, line1: Sonnemannstrasse 20, line2: line2, postal_code: '60314', state: x}",
			"--creditor-name", "Ian Crease",
			"--debtor-address", "{city: New York, country: US, line1: 33 Liberty Street, line2: line2, postal_code: '10045', state: NY}",
			"--debtor-name", "National Phonograph Company",
			"--instructed-amount", "100",
			"--instructed-currency", "USD",
			"--source-account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--unstructured-remittance-information", "New Swift transfer",
			"--require-approval=true",
			"--routing-number", "sq",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(swiftTransfersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"swift-transfers", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--account-number", "987654321",
			"--bank-identification-code", "ECBFDEFFTPP",
			"--creditor-address.city", "Frankfurt",
			"--creditor-address.country", "DE",
			"--creditor-address.line1", "Sonnemannstrasse 20",
			"--creditor-address.line2", "line2",
			"--creditor-address.postal-code", "60314",
			"--creditor-address.state", "x",
			"--creditor-name", "Ian Crease",
			"--debtor-address.city", "New York",
			"--debtor-address.country", "US",
			"--debtor-address.line1", "33 Liberty Street",
			"--debtor-address.line2", "line2",
			"--debtor-address.postal-code", "10045",
			"--debtor-address.state", "NY",
			"--debtor-name", "National Phonograph Company",
			"--instructed-amount", "100",
			"--instructed-currency", "USD",
			"--source-account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--unstructured-remittance-information", "New Swift transfer",
			"--require-approval=true",
			"--routing-number", "sq",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_in71c4amph0vgo2qllky\n" +
			"account_number: '987654321'\n" +
			"bank_identification_code: ECBFDEFFTPP\n" +
			"creditor_address:\n" +
			"  city: Frankfurt\n" +
			"  country: DE\n" +
			"  line1: Sonnemannstrasse 20\n" +
			"  line2: line2\n" +
			"  postal_code: '60314'\n" +
			"  state: x\n" +
			"creditor_name: Ian Crease\n" +
			"debtor_address:\n" +
			"  city: New York\n" +
			"  country: US\n" +
			"  line1: 33 Liberty Street\n" +
			"  line2: line2\n" +
			"  postal_code: '10045'\n" +
			"  state: NY\n" +
			"debtor_name: National Phonograph Company\n" +
			"instructed_amount: 100\n" +
			"instructed_currency: USD\n" +
			"source_account_number_id: account_number_v18nkfqm6afpsrvy82b2\n" +
			"unstructured_remittance_information: New Swift transfer\n" +
			"require_approval: true\n" +
			"routing_number: sq\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"swift-transfers", "create",
		)
	})
}

func TestSwiftTransfersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"swift-transfers", "retrieve",
			"--swift-transfer-id", "swift_transfer_29h21xkng03788zwd3fh",
		)
	})
}

func TestSwiftTransfersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"swift-transfers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [pending_approval]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(swiftTransfersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"swift-transfers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[pending_approval]",
		)
	})
}

func TestSwiftTransfersApprove(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"swift-transfers", "approve",
			"--swift-transfer-id", "swift_transfer_29h21xkng03788zwd3fh",
		)
	})
}

func TestSwiftTransfersCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"swift-transfers", "cancel",
			"--swift-transfer-id", "swift_transfer_29h21xkng03788zwd3fh",
		)
	})
}
