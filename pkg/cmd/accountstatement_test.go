// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestAccountStatementsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-statements", "retrieve",
			"--account-statement-id", "account_statement_lkc03a4skm2k7f38vj15",
		)
	})
}

func TestAccountStatementsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-statements", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--cursor", "cursor",
			"--limit", "1",
			"--statement-period-start", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountStatementsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account-statements", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--cursor", "cursor",
			"--limit", "1",
			"--statement-period-start.after", "2019-12-27T18:11:19.117Z",
			"--statement-period-start.before", "2019-12-27T18:11:19.117Z",
			"--statement-period-start.on-or-after", "2019-12-27T18:11:19.117Z",
			"--statement-period-start.on-or-before", "2019-12-27T18:11:19.117Z",
		)
	})
}
