// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestPhysicalCardsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-cards", "create",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
			"--cardholder", "{first_name: Ian, last_name: Crease}",
			"--shipment", "{address: {city: New York, line1: 33 Liberty Street, name: Ian Crease, postal_code: '10045', state: NY, country: x, line2: Unit 2, line3: x, phone_number: x}, method: usps, schedule: next_day}",
			"--physical-card-profile-id", "physical_card_profile_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(physicalCardsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-cards", "create",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
			"--cardholder.first-name", "Ian",
			"--cardholder.last-name", "Crease",
			"--shipment.address", "{city: New York, line1: 33 Liberty Street, name: Ian Crease, postal_code: '10045', state: NY, country: x, line2: Unit 2, line3: x, phone_number: x}",
			"--shipment.method", "usps",
			"--shipment.schedule", "next_day",
			"--physical-card-profile-id", "physical_card_profile_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"card_id: card_oubs0hwk5rn6knuecxg2\n" +
			"cardholder:\n" +
			"  first_name: Ian\n" +
			"  last_name: Crease\n" +
			"shipment:\n" +
			"  address:\n" +
			"    city: New York\n" +
			"    line1: 33 Liberty Street\n" +
			"    name: Ian Crease\n" +
			"    postal_code: '10045'\n" +
			"    state: NY\n" +
			"    country: x\n" +
			"    line2: Unit 2\n" +
			"    line3: x\n" +
			"    phone_number: x\n" +
			"  method: usps\n" +
			"  schedule: next_day\n" +
			"physical_card_profile_id: physical_card_profile_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"physical-cards", "create",
		)
	})
}

func TestPhysicalCardsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-cards", "retrieve",
			"--physical-card-id", "physical_card_ode8duyq5v2ynhjoharl",
		)
	})
}

func TestPhysicalCardsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-cards", "update",
			"--physical-card-id", "physical_card_ode8duyq5v2ynhjoharl",
			"--status", "disabled",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("status: disabled")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"physical-cards", "update",
			"--physical-card-id", "physical_card_ode8duyq5v2ynhjoharl",
		)
	})
}

func TestPhysicalCardsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-cards", "list",
			"--max-items", "10",
			"--card-id", "card_id",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(physicalCardsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"physical-cards", "list",
			"--max-items", "10",
			"--card-id", "card_id",
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
