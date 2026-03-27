// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestBookkeepingAccountsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bookkeeping-accounts", "create",
			"--name", "New Account!",
			"--account-id", "account_id",
			"--compliance-category", "commingled_cash",
			"--entity-id", "entity_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: New Account!\n" +
			"account_id: account_id\n" +
			"compliance_category: commingled_cash\n" +
			"entity_id: entity_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"bookkeeping-accounts", "create",
		)
	})
}

func TestBookkeepingAccountsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bookkeeping-accounts", "update",
			"--bookkeeping-account-id", "bookkeeping_account_e37p1f1iuocw5intf35v",
			"--name", "Deprecated Account",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: Deprecated Account")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"bookkeeping-accounts", "update",
			"--bookkeeping-account-id", "bookkeeping_account_e37p1f1iuocw5intf35v",
		)
	})
}

func TestBookkeepingAccountsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bookkeeping-accounts", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
		)
	})
}

func TestBookkeepingAccountsBalance(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bookkeeping-accounts", "balance",
			"--bookkeeping-account-id", "bookkeeping_account_e37p1f1iuocw5intf35v",
			"--at-time", "'2019-12-27T18:11:19.117Z'",
		)
	})
}
