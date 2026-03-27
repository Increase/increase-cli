// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestBookkeepingEntriesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bookkeeping-entries", "retrieve",
			"--bookkeeping-entry-id", "bookkeeping_entry_ctjpajsj3ks2blx10375",
		)
	})
}

func TestBookkeepingEntriesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bookkeeping-entries", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}
