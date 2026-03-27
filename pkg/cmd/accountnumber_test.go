// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestAccountNumbersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-numbers", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--name", "Rent payments",
			"--inbound-ach", "{debit_status: allowed}",
			"--inbound-checks", "{status: allowed}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountNumbersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-numbers", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--name", "Rent payments",
			"--inbound-ach.debit-status", "allowed",
			"--inbound-checks.status", "allowed",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_in71c4amph0vgo2qllky\n" +
			"name: Rent payments\n" +
			"inbound_ach:\n" +
			"  debit_status: allowed\n" +
			"inbound_checks:\n" +
			"  status: allowed\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"account-numbers", "create",
		)
	})
}

func TestAccountNumbersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-numbers", "retrieve",
			"--account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
		)
	})
}

func TestAccountNumbersUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-numbers", "update",
			"--account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--inbound-ach", "{debit_status: blocked}",
			"--inbound-checks", "{status: allowed}",
			"--name", "x",
			"--status", "disabled",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountNumbersUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-numbers", "update",
			"--account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--inbound-ach.debit-status", "blocked",
			"--inbound-checks.status", "allowed",
			"--name", "x",
			"--status", "disabled",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"inbound_ach:\n" +
			"  debit_status: blocked\n" +
			"inbound_checks:\n" +
			"  status: allowed\n" +
			"name: x\n" +
			"status: disabled\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"account-numbers", "update",
			"--account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
		)
	})
}

func TestAccountNumbersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-numbers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--ach-debit-status", "{in: [allowed]}",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [active]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountNumbersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-numbers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--ach-debit-status.in", "[allowed]",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[active]",
		)
	})
}
