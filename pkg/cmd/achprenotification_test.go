// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestACHPrenotificationsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ach-prenotifications", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--account-number", "987654321",
			"--routing-number", "101050001",
			"--addendum", "x",
			"--company-descriptive-date", "x",
			"--company-discretionary-data", "x",
			"--company-entry-description", "x",
			"--company-name", "x",
			"--credit-debit-indicator", "credit",
			"--effective-date", "'2019-12-27'",
			"--individual-id", "x",
			"--individual-name", "x",
			"--standard-entry-class-code", "corporate_credit_or_debit",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_in71c4amph0vgo2qllky\n" +
			"account_number: '987654321'\n" +
			"routing_number: '101050001'\n" +
			"addendum: x\n" +
			"company_descriptive_date: x\n" +
			"company_discretionary_data: x\n" +
			"company_entry_description: x\n" +
			"company_name: x\n" +
			"credit_debit_indicator: credit\n" +
			"effective_date: '2019-12-27'\n" +
			"individual_id: x\n" +
			"individual_name: x\n" +
			"standard_entry_class_code: corporate_credit_or_debit\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ach-prenotifications", "create",
		)
	})
}

func TestACHPrenotificationsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ach-prenotifications", "retrieve",
			"--ach-prenotification-id", "ach_prenotification_ubjf9qqsxl3obbcn1u34",
		)
	})
}

func TestACHPrenotificationsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ach-prenotifications", "list",
			"--max-items", "10",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(achPrenotificationsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ach-prenotifications", "list",
			"--max-items", "10",
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
