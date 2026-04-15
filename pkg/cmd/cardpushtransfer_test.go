// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestCardPushTransfersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-push-transfers", "create",
			"--business-application-identifier", "funds_disbursement",
			"--card-token-id", "outbound_card_token_zlt0ml6youq3q7vcdlg0",
			"--merchant-category-code", "1234",
			"--merchant-city-name", "New York",
			"--merchant-name", "Acme Corp",
			"--merchant-name-prefix", "Acme",
			"--merchant-postal-code", "10045",
			"--merchant-state", "NY",
			"--presentment-amount", "{currency: USD, value: '1234.56'}",
			"--recipient-name", "Ian Crease",
			"--sender-address-city", "New York",
			"--sender-address-line1", "33 Liberty Street",
			"--sender-address-postal-code", "10045",
			"--sender-address-state", "NY",
			"--sender-name", "Ian Crease",
			"--source-account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--merchant-legal-business-name", "x6",
			"--merchant-street-address", "merchant_street_address",
			"--recipient-address-city", "recipient_address_city",
			"--recipient-address-line1", "recipient_address_line1",
			"--recipient-address-postal-code", "x6",
			"--recipient-address-state", "x6",
			"--require-approval=true",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(cardPushTransfersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-push-transfers", "create",
			"--business-application-identifier", "funds_disbursement",
			"--card-token-id", "outbound_card_token_zlt0ml6youq3q7vcdlg0",
			"--merchant-category-code", "1234",
			"--merchant-city-name", "New York",
			"--merchant-name", "Acme Corp",
			"--merchant-name-prefix", "Acme",
			"--merchant-postal-code", "10045",
			"--merchant-state", "NY",
			"--presentment-amount.currency", "USD",
			"--presentment-amount.value", "1234.56",
			"--recipient-name", "Ian Crease",
			"--sender-address-city", "New York",
			"--sender-address-line1", "33 Liberty Street",
			"--sender-address-postal-code", "10045",
			"--sender-address-state", "NY",
			"--sender-name", "Ian Crease",
			"--source-account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--merchant-legal-business-name", "x6",
			"--merchant-street-address", "merchant_street_address",
			"--recipient-address-city", "recipient_address_city",
			"--recipient-address-line1", "recipient_address_line1",
			"--recipient-address-postal-code", "x6",
			"--recipient-address-state", "x6",
			"--require-approval=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"business_application_identifier: funds_disbursement\n" +
			"card_token_id: outbound_card_token_zlt0ml6youq3q7vcdlg0\n" +
			"merchant_category_code: '1234'\n" +
			"merchant_city_name: New York\n" +
			"merchant_name: Acme Corp\n" +
			"merchant_name_prefix: Acme\n" +
			"merchant_postal_code: '10045'\n" +
			"merchant_state: NY\n" +
			"presentment_amount:\n" +
			"  currency: USD\n" +
			"  value: '1234.56'\n" +
			"recipient_name: Ian Crease\n" +
			"sender_address_city: New York\n" +
			"sender_address_line1: 33 Liberty Street\n" +
			"sender_address_postal_code: '10045'\n" +
			"sender_address_state: NY\n" +
			"sender_name: Ian Crease\n" +
			"source_account_number_id: account_number_v18nkfqm6afpsrvy82b2\n" +
			"merchant_legal_business_name: x6\n" +
			"merchant_street_address: merchant_street_address\n" +
			"recipient_address_city: recipient_address_city\n" +
			"recipient_address_line1: recipient_address_line1\n" +
			"recipient_address_postal_code: x6\n" +
			"recipient_address_state: x6\n" +
			"require_approval: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"card-push-transfers", "create",
		)
	})
}

func TestCardPushTransfersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-push-transfers", "retrieve",
			"--card-push-transfer-id", "outbound_card_push_transfer_e0z9rdpamraczh4tvwye",
		)
	})
}

func TestCardPushTransfersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-push-transfers", "list",
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
		requestflag.CheckInnerFlags(cardPushTransfersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-push-transfers", "list",
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

func TestCardPushTransfersApprove(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-push-transfers", "approve",
			"--card-push-transfer-id", "outbound_card_push_transfer_e0z9rdpamraczh4tvwye",
		)
	})
}

func TestCardPushTransfersCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-push-transfers", "cancel",
			"--card-push-transfer-id", "outbound_card_push_transfer_e0z9rdpamraczh4tvwye",
		)
	})
}
