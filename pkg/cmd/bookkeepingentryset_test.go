// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestBookkeepingEntrySetsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bookkeeping-entry-sets", "create",
			"--entry", "{account_id: bookkeeping_account_9husfpw68pzmve9dvvc7, amount: 100}",
			"--entry", "{account_id: bookkeeping_account_t2obldz1rcu15zr54umg, amount: -100}",
			"--date", "'2020-01-31T23:59:59Z'",
			"--transaction-id", "transaction_uyrp7fld2ium70oa7oi",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(bookkeepingEntrySetsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bookkeeping-entry-sets", "create",
			"--entry.account-id", "bookkeeping_account_9husfpw68pzmve9dvvc7",
			"--entry.amount", "100",
			"--entry.account-id", "bookkeeping_account_t2obldz1rcu15zr54umg",
			"--entry.amount", "-100",
			"--date", "'2020-01-31T23:59:59Z'",
			"--transaction-id", "transaction_uyrp7fld2ium70oa7oi",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"entries:\n" +
			"  - account_id: bookkeeping_account_9husfpw68pzmve9dvvc7\n" +
			"    amount: 100\n" +
			"  - account_id: bookkeeping_account_t2obldz1rcu15zr54umg\n" +
			"    amount: -100\n" +
			"date: '2020-01-31T23:59:59Z'\n" +
			"transaction_id: transaction_uyrp7fld2ium70oa7oi\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"bookkeeping-entry-sets", "create",
		)
	})
}

func TestBookkeepingEntrySetsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bookkeeping-entry-sets", "retrieve",
			"--bookkeeping-entry-set-id", "bookkeeping_entry_set_n80c6wr2p8gtc6p4ingf",
		)
	})
}

func TestBookkeepingEntrySetsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bookkeeping-entry-sets", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--transaction-id", "transaction_id",
		)
	})
}
