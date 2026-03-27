// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestAccountsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "create",
			"--name", "New Account!",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
			"--funding", "loan",
			"--informational-entity-id", "informational_entity_id",
			"--loan", "{credit_limit: 0, grace_period_days: 0, statement_day_of_month: 1, statement_payment_type: balance, maturity_date: '2019-12-27'}",
			"--program-id", "program_i2v2os4mwza1oetokh9i",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "create",
			"--name", "New Account!",
			"--entity-id", "entity_n8y8tnk2p9339ti393yi",
			"--funding", "loan",
			"--informational-entity-id", "informational_entity_id",
			"--loan.credit-limit", "0",
			"--loan.grace-period-days", "0",
			"--loan.statement-day-of-month", "1",
			"--loan.statement-payment-type", "balance",
			"--loan.maturity-date", "2019-12-27",
			"--program-id", "program_i2v2os4mwza1oetokh9i",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: New Account!\n" +
			"entity_id: entity_n8y8tnk2p9339ti393yi\n" +
			"funding: loan\n" +
			"informational_entity_id: informational_entity_id\n" +
			"loan:\n" +
			"  credit_limit: 0\n" +
			"  grace_period_days: 0\n" +
			"  statement_day_of_month: 1\n" +
			"  statement_payment_type: balance\n" +
			"  maturity_date: '2019-12-27'\n" +
			"program_id: program_i2v2os4mwza1oetokh9i\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"accounts", "create",
		)
	})
}

func TestAccountsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "retrieve",
			"--account-id", "account_in71c4amph0vgo2qllky",
		)
	})
}

func TestAccountsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "update",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--loan", "{credit_limit: 0}",
			"--name", "My renamed account",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "update",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--loan.credit-limit", "0",
			"--name", "My renamed account",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"loan:\n" +
			"  credit_limit: 0\n" +
			"name: My renamed account\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"accounts", "update",
			"--account-id", "account_in71c4amph0vgo2qllky",
		)
	})
}

func TestAccountsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "list",
			"--max-items", "10",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--entity-id", "entity_id",
			"--idempotency-key", "x",
			"--informational-entity-id", "informational_entity_id",
			"--limit", "1",
			"--program-id", "program_id",
			"--status", "{in: [closed]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "list",
			"--max-items", "10",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--entity-id", "entity_id",
			"--idempotency-key", "x",
			"--informational-entity-id", "informational_entity_id",
			"--limit", "1",
			"--program-id", "program_id",
			"--status.in", "[closed]",
		)
	})
}

func TestAccountsBalance(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "balance",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--at-time", "'2019-12-27T18:11:19.117Z'",
		)
	})
}

func TestAccountsClose(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "close",
			"--account-id", "account_in71c4amph0vgo2qllky",
		)
	})
}
