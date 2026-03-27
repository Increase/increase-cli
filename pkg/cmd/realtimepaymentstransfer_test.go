// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestRealTimePaymentsTransfersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"real-time-payments-transfers", "create",
			"--amount", "100",
			"--creditor-name", "Ian Crease",
			"--source-account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--unstructured-remittance-information", "Invoice 29582",
			"--account-number", "987654321",
			"--debtor-name", "x",
			"--external-account-id", "external_account_id",
			"--require-approval=true",
			"--routing-number", "101050001",
			"--ultimate-creditor-name", "x",
			"--ultimate-debtor-name", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 100\n" +
			"creditor_name: Ian Crease\n" +
			"source_account_number_id: account_number_v18nkfqm6afpsrvy82b2\n" +
			"unstructured_remittance_information: Invoice 29582\n" +
			"account_number: '987654321'\n" +
			"debtor_name: x\n" +
			"external_account_id: external_account_id\n" +
			"require_approval: true\n" +
			"routing_number: '101050001'\n" +
			"ultimate_creditor_name: x\n" +
			"ultimate_debtor_name: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"real-time-payments-transfers", "create",
		)
	})
}

func TestRealTimePaymentsTransfersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"real-time-payments-transfers", "retrieve",
			"--real-time-payments-transfer-id", "real_time_payments_transfer_iyuhl5kdn7ssmup83mvq",
		)
	})
}

func TestRealTimePaymentsTransfersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"real-time-payments-transfers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--external-account-id", "external_account_id",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [pending_approval]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(realTimePaymentsTransfersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"real-time-payments-transfers", "list",
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
			"--status.in", "[pending_approval]",
		)
	})
}

func TestRealTimePaymentsTransfersApprove(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"real-time-payments-transfers", "approve",
			"--real-time-payments-transfer-id", "real_time_payments_transfer_iyuhl5kdn7ssmup83mvq",
		)
	})
}

func TestRealTimePaymentsTransfersCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"real-time-payments-transfers", "cancel",
			"--real-time-payments-transfer-id", "real_time_payments_transfer_iyuhl5kdn7ssmup83mvq",
		)
	})
}
