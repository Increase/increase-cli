// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestExternalAccountsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-accounts", "create",
			"--account-number", "987654321",
			"--description", "Landlord",
			"--routing-number", "101050001",
			"--account-holder", "business",
			"--funding", "checking",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_number: '987654321'\n" +
			"description: Landlord\n" +
			"routing_number: '101050001'\n" +
			"account_holder: business\n" +
			"funding: checking\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"external-accounts", "create",
		)
	})
}

func TestExternalAccountsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-accounts", "retrieve",
			"--external-account-id", "external_account_ukk55lr923a3ac0pp7iv",
		)
	})
}

func TestExternalAccountsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-accounts", "update",
			"--external-account-id", "external_account_ukk55lr923a3ac0pp7iv",
			"--account-holder", "business",
			"--description", "New description",
			"--funding", "checking",
			"--status", "active",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_holder: business\n" +
			"description: New description\n" +
			"funding: checking\n" +
			"status: active\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"external-accounts", "update",
			"--external-account-id", "external_account_ukk55lr923a3ac0pp7iv",
		)
	})
}

func TestExternalAccountsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-accounts", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--routing-number", "xxxxxxxxx",
			"--status", "{in: [active]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(externalAccountsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"external-accounts", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--routing-number", "xxxxxxxxx",
			"--status.in", "[active]",
		)
	})
}
