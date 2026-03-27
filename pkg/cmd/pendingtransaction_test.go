// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestPendingTransactionsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pending-transactions", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--amount", "-1000",
			"--description", "Hold for pending transaction",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_in71c4amph0vgo2qllky\n" +
			"amount: -1000\n" +
			"description: Hold for pending transaction\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pending-transactions", "create",
		)
	})
}

func TestPendingTransactionsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pending-transactions", "retrieve",
			"--pending-transaction-id", "pending_transaction_k1sfetcau2qbvjbzgju4",
		)
	})
}

func TestPendingTransactionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pending-transactions", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--category", "{in: [account_transfer_instruction]}",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--limit", "1",
			"--route-id", "route_id",
			"--status", "{in: [pending]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(pendingTransactionsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pending-transactions", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--category.in", "[account_transfer_instruction]",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--limit", "1",
			"--route-id", "route_id",
			"--status.in", "[pending]",
		)
	})
}

func TestPendingTransactionsRelease(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pending-transactions", "release",
			"--pending-transaction-id", "pending_transaction_k1sfetcau2qbvjbzgju4",
		)
	})
}
