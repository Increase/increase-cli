// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestCheckTransfersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"check-transfers", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--amount", "1000",
			"--fulfillment-method", "physical_check",
			"--source-account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--balance-check", "full",
			"--check-number", "469",
			"--physical-check", "{mailing_address: {city: New York, line1: 33 Liberty Street, postal_code: '10045', state: NY, line2: line2, name: Ian Crease, phone: '+16505046304'}, memo: Check payment, recipient_name: Ian Crease, attachment_file_id: attachment_file_id, check_voucher_image_file_id: check_voucher_image_file_id, note: note, payer: [{contents: contents}], return_address: {city: city, line1: line1, name: name, postal_code: 21029-9469, state: x, line2: line2, phone: x}, shipping_method: usps_first_class, signature: {image_file_id: image_file_id, text: Ian Crease}}",
			"--require-approval=true",
			"--third-party", "{recipient_name: x}",
			"--valid-until-date", "'2025-12-31'",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(checkTransfersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"check-transfers", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--amount", "1000",
			"--fulfillment-method", "physical_check",
			"--source-account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--balance-check", "full",
			"--check-number", "469",
			"--physical-check.mailing-address", "{city: New York, line1: 33 Liberty Street, postal_code: '10045', state: NY, line2: line2, name: Ian Crease, phone: '+16505046304'}",
			"--physical-check.memo", "Check payment",
			"--physical-check.recipient-name", "Ian Crease",
			"--physical-check.attachment-file-id", "attachment_file_id",
			"--physical-check.check-voucher-image-file-id", "check_voucher_image_file_id",
			"--physical-check.note", "note",
			"--physical-check.payer", "[{contents: contents}]",
			"--physical-check.return-address", "{city: city, line1: line1, name: name, postal_code: 21029-9469, state: x, line2: line2, phone: x}",
			"--physical-check.shipping-method", "usps_first_class",
			"--physical-check.signature", "{image_file_id: image_file_id, text: Ian Crease}",
			"--require-approval=true",
			"--third-party.recipient-name", "x",
			"--valid-until-date", "'2025-12-31'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_in71c4amph0vgo2qllky\n" +
			"amount: 1000\n" +
			"fulfillment_method: physical_check\n" +
			"source_account_number_id: account_number_v18nkfqm6afpsrvy82b2\n" +
			"balance_check: full\n" +
			"check_number: '469'\n" +
			"physical_check:\n" +
			"  mailing_address:\n" +
			"    city: New York\n" +
			"    line1: 33 Liberty Street\n" +
			"    postal_code: '10045'\n" +
			"    state: NY\n" +
			"    line2: line2\n" +
			"    name: Ian Crease\n" +
			"    phone: '+16505046304'\n" +
			"  memo: Check payment\n" +
			"  recipient_name: Ian Crease\n" +
			"  attachment_file_id: attachment_file_id\n" +
			"  check_voucher_image_file_id: check_voucher_image_file_id\n" +
			"  note: note\n" +
			"  payer:\n" +
			"    - contents: contents\n" +
			"  return_address:\n" +
			"    city: city\n" +
			"    line1: line1\n" +
			"    name: name\n" +
			"    postal_code: 21029-9469\n" +
			"    state: x\n" +
			"    line2: line2\n" +
			"    phone: x\n" +
			"  shipping_method: usps_first_class\n" +
			"  signature:\n" +
			"    image_file_id: image_file_id\n" +
			"    text: Ian Crease\n" +
			"require_approval: true\n" +
			"third_party:\n" +
			"  recipient_name: x\n" +
			"valid_until_date: '2025-12-31'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"check-transfers", "create",
		)
	})
}

func TestCheckTransfersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"check-transfers", "retrieve",
			"--check-transfer-id", "check_transfer_30b43acfu9vw8fyc4f5",
		)
	})
}

func TestCheckTransfersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"check-transfers", "list",
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
		requestflag.CheckInnerFlags(checkTransfersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"check-transfers", "list",
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

func TestCheckTransfersApprove(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"check-transfers", "approve",
			"--check-transfer-id", "check_transfer_30b43acfu9vw8fyc4f5",
		)
	})
}

func TestCheckTransfersCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"check-transfers", "cancel",
			"--check-transfer-id", "check_transfer_30b43acfu9vw8fyc4f5",
		)
	})
}

func TestCheckTransfersStopPayment(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"check-transfers", "stop-payment",
			"--check-transfer-id", "check_transfer_30b43acfu9vw8fyc4f5",
			"--reason", "mail_delivery_failed",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("reason: mail_delivery_failed")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"check-transfers", "stop-payment",
			"--check-transfer-id", "check_transfer_30b43acfu9vw8fyc4f5",
		)
	})
}
