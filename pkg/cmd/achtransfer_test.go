// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestACHTransfersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ach-transfers", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--amount", "100",
			"--statement-descriptor", "New ACH transfer",
			"--account-number", "987654321",
			"--addenda", "{category: freeform, freeform: {entries: [{payment_related_information: payment_related_information}]}, payment_order_remittance_advice: {invoices: [{invoice_number: sq, paid_amount: 0}]}}",
			"--company-descriptive-date", "J!",
			"--company-discretionary-data", "J!",
			"--company-entry-description", "J!",
			"--company-name", "company_name",
			"--destination-account-holder", "business",
			"--external-account-id", "external_account_id",
			"--funding", "checking",
			"--individual-id", "individual_id",
			"--individual-name", "individual_name",
			"--preferred-effective-date", "{date: '2019-12-27', settlement_schedule: same_day}",
			"--require-approval=true",
			"--routing-number", "101050001",
			"--standard-entry-class-code", "corporate_credit_or_debit",
			"--transaction-timing", "synchronous",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(achTransfersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ach-transfers", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--amount", "100",
			"--statement-descriptor", "New ACH transfer",
			"--account-number", "987654321",
			"--addenda.category", "freeform",
			"--addenda.freeform", "{entries: [{payment_related_information: payment_related_information}]}",
			"--addenda.payment-order-remittance-advice", "{invoices: [{invoice_number: sq, paid_amount: 0}]}",
			"--company-descriptive-date", "J!",
			"--company-discretionary-data", "J!",
			"--company-entry-description", "J!",
			"--company-name", "company_name",
			"--destination-account-holder", "business",
			"--external-account-id", "external_account_id",
			"--funding", "checking",
			"--individual-id", "individual_id",
			"--individual-name", "individual_name",
			"--preferred-effective-date.date", "2019-12-27",
			"--preferred-effective-date.settlement-schedule", "same_day",
			"--require-approval=true",
			"--routing-number", "101050001",
			"--standard-entry-class-code", "corporate_credit_or_debit",
			"--transaction-timing", "synchronous",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_in71c4amph0vgo2qllky\n" +
			"amount: 100\n" +
			"statement_descriptor: New ACH transfer\n" +
			"account_number: '987654321'\n" +
			"addenda:\n" +
			"  category: freeform\n" +
			"  freeform:\n" +
			"    entries:\n" +
			"      - payment_related_information: payment_related_information\n" +
			"  payment_order_remittance_advice:\n" +
			"    invoices:\n" +
			"      - invoice_number: sq\n" +
			"        paid_amount: 0\n" +
			"company_descriptive_date: J!\n" +
			"company_discretionary_data: J!\n" +
			"company_entry_description: J!\n" +
			"company_name: company_name\n" +
			"destination_account_holder: business\n" +
			"external_account_id: external_account_id\n" +
			"funding: checking\n" +
			"individual_id: individual_id\n" +
			"individual_name: individual_name\n" +
			"preferred_effective_date:\n" +
			"  date: '2019-12-27'\n" +
			"  settlement_schedule: same_day\n" +
			"require_approval: true\n" +
			"routing_number: '101050001'\n" +
			"standard_entry_class_code: corporate_credit_or_debit\n" +
			"transaction_timing: synchronous\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ach-transfers", "create",
		)
	})
}

func TestACHTransfersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ach-transfers", "retrieve",
			"--ach-transfer-id", "ach_transfer_uoxatyh3lt5evrsdvo7q",
		)
	})
}

func TestACHTransfersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ach-transfers", "list",
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
		requestflag.CheckInnerFlags(achTransfersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ach-transfers", "list",
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

func TestACHTransfersApprove(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ach-transfers", "approve",
			"--ach-transfer-id", "ach_transfer_uoxatyh3lt5evrsdvo7q",
		)
	})
}

func TestACHTransfersCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ach-transfers", "cancel",
			"--ach-transfer-id", "ach_transfer_uoxatyh3lt5evrsdvo7q",
		)
	})
}
