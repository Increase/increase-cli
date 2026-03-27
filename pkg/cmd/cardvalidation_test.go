// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestCardValidationsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-validations", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--card-token-id", "outbound_card_token_zlt0ml6youq3q7vcdlg0",
			"--merchant-category-code", "1234",
			"--merchant-city-name", "New York",
			"--merchant-name", "Acme Corp",
			"--merchant-postal-code", "10045",
			"--merchant-state", "NY",
			"--cardholder-first-name", "Dee",
			"--cardholder-last-name", "Hock",
			"--cardholder-middle-name", "Ward",
			"--cardholder-postal-code", "10045",
			"--cardholder-street-address", "33 Liberty Street",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_in71c4amph0vgo2qllky\n" +
			"card_token_id: outbound_card_token_zlt0ml6youq3q7vcdlg0\n" +
			"merchant_category_code: '1234'\n" +
			"merchant_city_name: New York\n" +
			"merchant_name: Acme Corp\n" +
			"merchant_postal_code: '10045'\n" +
			"merchant_state: NY\n" +
			"cardholder_first_name: Dee\n" +
			"cardholder_last_name: Hock\n" +
			"cardholder_middle_name: Ward\n" +
			"cardholder_postal_code: '10045'\n" +
			"cardholder_street_address: 33 Liberty Street\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"card-validations", "create",
		)
	})
}

func TestCardValidationsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-validations", "retrieve",
			"--card-validation-id", "outbound_card_validation_qqlzagpc6v1x2gcdhe24",
		)
	})
}

func TestCardValidationsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-validations", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [requires_attention]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(cardValidationsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-validations", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[requires_attention]",
		)
	})
}
