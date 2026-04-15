// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestSimulationsInboundACHTransfersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:inbound-ach-transfers", "create",
			"--account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--amount", "1000",
			"--addenda", "{category: freeform, freeform: {entries: [{payment_related_information: payment_related_information}]}}",
			"--company-descriptive-date", "J!",
			"--company-discretionary-data", "J!",
			"--company-entry-description", "J!",
			"--company-id", "company_id",
			"--company-name", "company_name",
			"--receiver-id-number", "x",
			"--receiver-name", "x",
			"--resolve-at", "'2019-12-27T18:11:19.117Z'",
			"--standard-entry-class-code", "corporate_credit_or_debit",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simulationsInboundACHTransfersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:inbound-ach-transfers", "create",
			"--account-number-id", "account_number_v18nkfqm6afpsrvy82b2",
			"--amount", "1000",
			"--addenda.category", "freeform",
			"--addenda.freeform", "{entries: [{payment_related_information: payment_related_information}]}",
			"--company-descriptive-date", "J!",
			"--company-discretionary-data", "J!",
			"--company-entry-description", "J!",
			"--company-id", "company_id",
			"--company-name", "company_name",
			"--receiver-id-number", "x",
			"--receiver-name", "x",
			"--resolve-at", "'2019-12-27T18:11:19.117Z'",
			"--standard-entry-class-code", "corporate_credit_or_debit",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_number_id: account_number_v18nkfqm6afpsrvy82b2\n" +
			"amount: 1000\n" +
			"addenda:\n" +
			"  category: freeform\n" +
			"  freeform:\n" +
			"    entries:\n" +
			"      - payment_related_information: payment_related_information\n" +
			"company_descriptive_date: J!\n" +
			"company_discretionary_data: J!\n" +
			"company_entry_description: J!\n" +
			"company_id: company_id\n" +
			"company_name: company_name\n" +
			"receiver_id_number: x\n" +
			"receiver_name: x\n" +
			"resolve_at: '2019-12-27T18:11:19.117Z'\n" +
			"standard_entry_class_code: corporate_credit_or_debit\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:inbound-ach-transfers", "create",
		)
	})
}
