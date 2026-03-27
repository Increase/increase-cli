// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
)

func TestSimulationsPhysicalCardsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:physical-cards", "create",
			"--physical-card-id", "physical_card_ode8duyq5v2ynhjoharl",
			"--category", "delivered",
			"--carrier-estimated-delivery-at", "'2019-12-27T18:11:19.117Z'",
			"--city", "New York",
			"--postal-code", "10045",
			"--state", "NY",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"category: delivered\n" +
			"carrier_estimated_delivery_at: '2019-12-27T18:11:19.117Z'\n" +
			"city: New York\n" +
			"postal_code: '10045'\n" +
			"state: NY\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:physical-cards", "create",
			"--physical-card-id", "physical_card_ode8duyq5v2ynhjoharl",
		)
	})
}

func TestSimulationsPhysicalCardsAdvanceShipment(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:physical-cards", "advance-shipment",
			"--physical-card-id", "physical_card_ode8duyq5v2ynhjoharl",
			"--shipment-status", "shipped",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("shipment_status: shipped")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:physical-cards", "advance-shipment",
			"--physical-card-id", "physical_card_ode8duyq5v2ynhjoharl",
		)
	})
}
