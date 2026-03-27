// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestAccountTransfersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-transfers", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--amount", "100",
			"--description", "Creating liquidity",
			"--destination-account-id", "account_uf16sut2ct5bevmq3eh",
			"--require-approval=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_in71c4amph0vgo2qllky\n" +
			"amount: 100\n" +
			"description: Creating liquidity\n" +
			"destination_account_id: account_uf16sut2ct5bevmq3eh\n" +
			"require_approval: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"account-transfers", "create",
		)
	})
}

func TestAccountTransfersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-transfers", "retrieve",
			"--account-transfer-id", "account_transfer_7k9qe1ysdgqztnt63l7n",
		)
	})
}

func TestAccountTransfersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-transfers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountTransfersList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-transfers", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
		)
	})
}

func TestAccountTransfersApprove(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-transfers", "approve",
			"--account-transfer-id", "account_transfer_7k9qe1ysdgqztnt63l7n",
		)
	})
}

func TestAccountTransfersCancel(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-transfers", "cancel",
			"--account-transfer-id", "account_transfer_7k9qe1ysdgqztnt63l7n",
		)
	})
}
