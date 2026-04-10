// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestExportsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"exports", "create",
			"--category", "transaction_csv",
			"--account-statement-bai2", "{account_id: account_id, effective_date: '2019-12-27', program_id: program_id}",
			"--account-statement-ofx", "{account_id: account_id, created_at: {before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z'}}",
			"--account-verification-letter", "{account_number_id: account_number_id, balance_date: '2019-12-27'}",
			"--balance-csv", "{account_id: account_id, created_at: {after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}}",
			"--bookkeeping-account-balance-csv", "{bookkeeping_account_id: bookkeeping_account_id, on_or_after_date: '2019-12-27', on_or_before_date: '2019-12-27'}",
			"--daily-account-balance-csv", "{account_id: account_id, on_or_after_date: '2019-12-27', on_or_before_date: '2019-12-27'}",
			"--entity-csv", "{}",
			"--funding-instructions", "{account_number_id: account_number_id}",
			"--transaction-csv", "{account_id: account_in71c4amph0vgo2qllky, created_at: {after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}}",
			"--vendor-csv", "{}",
			"--voided-check", "{account_number_id: account_number_id, payer: [{line: x}]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(exportsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"exports", "create",
			"--category", "transaction_csv",
			"--account-statement-bai2.account-id", "account_id",
			"--account-statement-bai2.effective-date", "2019-12-27",
			"--account-statement-bai2.program-id", "program_id",
			"--account-statement-ofx.account-id", "account_id",
			"--account-statement-ofx.created-at", "{before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z'}",
			"--account-verification-letter.account-number-id", "account_number_id",
			"--account-verification-letter.balance-date", "2019-12-27",
			"--balance-csv.account-id", "account_id",
			"--balance-csv.created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--bookkeeping-account-balance-csv.bookkeeping-account-id", "bookkeeping_account_id",
			"--bookkeeping-account-balance-csv.on-or-after-date", "2019-12-27",
			"--bookkeeping-account-balance-csv.on-or-before-date", "2019-12-27",
			"--daily-account-balance-csv.account-id", "account_id",
			"--daily-account-balance-csv.on-or-after-date", "2019-12-27",
			"--daily-account-balance-csv.on-or-before-date", "2019-12-27",
			"--funding-instructions.account-number-id", "account_number_id",
			"--transaction-csv.account-id", "account_in71c4amph0vgo2qllky",
			"--transaction-csv.created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--voided-check.account-number-id", "account_number_id",
			"--voided-check.payer", "[{line: x}]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"category: transaction_csv\n" +
			"account_statement_bai2:\n" +
			"  account_id: account_id\n" +
			"  effective_date: '2019-12-27'\n" +
			"  program_id: program_id\n" +
			"account_statement_ofx:\n" +
			"  account_id: account_id\n" +
			"  created_at:\n" +
			"    before: '2019-12-27T18:11:19.117Z'\n" +
			"    on_or_after: '2019-12-27T18:11:19.117Z'\n" +
			"account_verification_letter:\n" +
			"  account_number_id: account_number_id\n" +
			"  balance_date: '2019-12-27'\n" +
			"balance_csv:\n" +
			"  account_id: account_id\n" +
			"  created_at:\n" +
			"    after: '2019-12-27T18:11:19.117Z'\n" +
			"    before: '2019-12-27T18:11:19.117Z'\n" +
			"    on_or_after: '2019-12-27T18:11:19.117Z'\n" +
			"    on_or_before: '2019-12-27T18:11:19.117Z'\n" +
			"bookkeeping_account_balance_csv:\n" +
			"  bookkeeping_account_id: bookkeeping_account_id\n" +
			"  on_or_after_date: '2019-12-27'\n" +
			"  on_or_before_date: '2019-12-27'\n" +
			"daily_account_balance_csv:\n" +
			"  account_id: account_id\n" +
			"  on_or_after_date: '2019-12-27'\n" +
			"  on_or_before_date: '2019-12-27'\n" +
			"entity_csv: {}\n" +
			"funding_instructions:\n" +
			"  account_number_id: account_number_id\n" +
			"transaction_csv:\n" +
			"  account_id: account_in71c4amph0vgo2qllky\n" +
			"  created_at:\n" +
			"    after: '2019-12-27T18:11:19.117Z'\n" +
			"    before: '2019-12-27T18:11:19.117Z'\n" +
			"    on_or_after: '2019-12-27T18:11:19.117Z'\n" +
			"    on_or_before: '2019-12-27T18:11:19.117Z'\n" +
			"vendor_csv: {}\n" +
			"voided_check:\n" +
			"  account_number_id: account_number_id\n" +
			"  payer:\n" +
			"    - line: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"exports", "create",
		)
	})
}

func TestExportsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"exports", "retrieve",
			"--export-id", "export_8s4m48qz3bclzje0zwh9",
		)
	})
}

func TestExportsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"exports", "list",
			"--max-items", "10",
			"--category", "account_statement_ofx",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--form-1099-int", "{account_id: account_id}",
			"--form-1099-misc", "{account_id: account_id}",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [pending]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(exportsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"exports", "list",
			"--max-items", "10",
			"--category", "account_statement_ofx",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--form-1099-int.account-id", "account_id",
			"--form-1099-misc.account-id", "account_id",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[pending]",
		)
	})
}
