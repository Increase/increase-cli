// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestFednowTransfersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fednow-transfers", "create",
			"--amount", "100",
			"--creditor-name", "Ian Crease",
			"--debtor-name", "National Phonograph Company",
			"--source-account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--unstructured-remittance-information", "Invoice 29582",
			"--account-number", "987654321",
			"--creditor-address", "{city: New York, postal_code: '10045', state: NY, line1: 33 Liberty Street}",
			"--debtor-address", "{city: x, postal_code: 21029-9469, state: xx, line1: x}",
			"--external-account-id", "external_account_id",
			"--require-approval=true",
			"--routing-number", "101050001",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(fednowTransfersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fednow-transfers", "create",
			"--amount", "100",
			"--creditor-name", "Ian Crease",
			"--debtor-name", "National Phonograph Company",
			"--source-account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--unstructured-remittance-information", "Invoice 29582",
			"--account-number", "987654321",
			"--creditor-address.city", "New York",
			"--creditor-address.postal-code", "10045",
			"--creditor-address.state", "NY",
			"--creditor-address.line1", "33 Liberty Street",
			"--debtor-address.city", "x",
			"--debtor-address.postal-code", "21029-9469",
			"--debtor-address.state", "xx",
			"--debtor-address.line1", "x",
			"--external-account-id", "external_account_id",
			"--require-approval=true",
			"--routing-number", "101050001",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 100\n" +
			"creditor_name: Ian Crease\n" +
			"debtor_name: National Phonograph Company\n" +
			"source_account_number_id: account_number_v18nkfqm6afpsrvy82b2\n" +
			"unstructured_remittance_information: Invoice 29582\n" +
			"account_number: '987654321'\n" +
			"creditor_address:\n" +
			"  city: New York\n" +
			"  postal_code: '10045'\n" +
			"  state: NY\n" +
			"  line1: 33 Liberty Street\n" +
			"debtor_address:\n" +
			"  city: x\n" +
			"  postal_code: 21029-9469\n" +
			"  state: xx\n" +
			"  line1: x\n" +
			"external_account_id: external_account_id\n" +
			"require_approval: true\n" +
			"routing_number: '101050001'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"fednow-transfers", "create",
		)
	})
}

func TestFednowTransfersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fednow-transfers", "retrieve",
			"--fednow-transfer-id", "fednow_transfer_4i0mptrdu1mueg1196bg",
		)
	})
}

func TestFednowTransfersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fednow-transfers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--external-account-id", "external_account_id",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [pending_reviewing]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(fednowTransfersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fednow-transfers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--external-account-id", "external_account_id",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[pending_reviewing]",
		)
	})
}

func TestFednowTransfersApprove(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fednow-transfers", "approve",
			"--fednow-transfer-id", "fednow_transfer_4i0mptrdu1mueg1196bg",
		)
	})
}

func TestFednowTransfersCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fednow-transfers", "cancel",
			"--fednow-transfer-id", "fednow_transfer_4i0mptrdu1mueg1196bg",
		)
	})
}
