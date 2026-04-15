// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestCardsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--authorization-controls", "{merchant_acceptor_identifier: {allowed: [{identifier: x}], blocked: [{identifier: x}]}, merchant_category_code: {allowed: [{code: '0253'}], blocked: [{code: '0253'}]}, merchant_country: {allowed: [{country: xx}], blocked: [{country: xx}]}, usage: {category: single_use, multi_use: {spending_limits: [{interval: all_time, settlement_amount: 0, merchant_category_codes: [{code: x}]}]}, single_use: {settlement_amount: {comparison: equals, value: 0}}}}",
			"--billing-address", "{city: x, line1: x, postal_code: x, state: x, line2: x}",
			"--description", "Card for Ian Crease",
			"--digital-wallet", "{digital_card_profile_id: digital_card_profile_id, email: dev@stainless.com, phone: x}",
			"--entity-id", "entity_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(cardsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--authorization-controls.merchant-acceptor-identifier", "{allowed: [{identifier: x}], blocked: [{identifier: x}]}",
			"--authorization-controls.merchant-category-code", "{allowed: [{code: '0253'}], blocked: [{code: '0253'}]}",
			"--authorization-controls.merchant-country", "{allowed: [{country: xx}], blocked: [{country: xx}]}",
			"--authorization-controls.usage", "{category: single_use, multi_use: {spending_limits: [{interval: all_time, settlement_amount: 0, merchant_category_codes: [{code: x}]}]}, single_use: {settlement_amount: {comparison: equals, value: 0}}}",
			"--billing-address.city", "x",
			"--billing-address.line1", "x",
			"--billing-address.postal-code", "x",
			"--billing-address.state", "x",
			"--billing-address.line2", "x",
			"--description", "Card for Ian Crease",
			"--digital-wallet.digital-card-profile-id", "digital_card_profile_id",
			"--digital-wallet.email", "dev@stainless.com",
			"--digital-wallet.phone", "x",
			"--entity-id", "entity_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_in71c4amph0vgo2qllky\n" +
			"authorization_controls:\n" +
			"  merchant_acceptor_identifier:\n" +
			"    allowed:\n" +
			"      - identifier: x\n" +
			"    blocked:\n" +
			"      - identifier: x\n" +
			"  merchant_category_code:\n" +
			"    allowed:\n" +
			"      - code: '0253'\n" +
			"    blocked:\n" +
			"      - code: '0253'\n" +
			"  merchant_country:\n" +
			"    allowed:\n" +
			"      - country: xx\n" +
			"    blocked:\n" +
			"      - country: xx\n" +
			"  usage:\n" +
			"    category: single_use\n" +
			"    multi_use:\n" +
			"      spending_limits:\n" +
			"        - interval: all_time\n" +
			"          settlement_amount: 0\n" +
			"          merchant_category_codes:\n" +
			"            - code: x\n" +
			"    single_use:\n" +
			"      settlement_amount:\n" +
			"        comparison: equals\n" +
			"        value: 0\n" +
			"billing_address:\n" +
			"  city: x\n" +
			"  line1: x\n" +
			"  postal_code: x\n" +
			"  state: x\n" +
			"  line2: x\n" +
			"description: Card for Ian Crease\n" +
			"digital_wallet:\n" +
			"  digital_card_profile_id: digital_card_profile_id\n" +
			"  email: dev@stainless.com\n" +
			"  phone: x\n" +
			"entity_id: entity_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cards", "create",
		)
	})
}

func TestCardsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "retrieve",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
		)
	})
}

func TestCardsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "update",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
			"--authorization-controls", "{merchant_acceptor_identifier: {allowed: [{identifier: x}], blocked: [{identifier: x}]}, merchant_category_code: {allowed: [{code: '0253'}], blocked: [{code: '0253'}]}, merchant_country: {allowed: [{country: xx}], blocked: [{country: xx}]}, usage: {category: single_use, multi_use: {spending_limits: [{interval: all_time, settlement_amount: 0, merchant_category_codes: [{code: x}]}]}, single_use: {settlement_amount: {comparison: equals, value: 0}}}}",
			"--billing-address", "{city: x, line1: x, postal_code: x, state: x, line2: x}",
			"--description", "New description",
			"--digital-wallet", "{digital_card_profile_id: digital_card_profile_id, email: dev@stainless.com, phone: x}",
			"--entity-id", "entity_id",
			"--status", "active",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(cardsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "update",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
			"--authorization-controls.merchant-acceptor-identifier", "{allowed: [{identifier: x}], blocked: [{identifier: x}]}",
			"--authorization-controls.merchant-category-code", "{allowed: [{code: '0253'}], blocked: [{code: '0253'}]}",
			"--authorization-controls.merchant-country", "{allowed: [{country: xx}], blocked: [{country: xx}]}",
			"--authorization-controls.usage", "{category: single_use, multi_use: {spending_limits: [{interval: all_time, settlement_amount: 0, merchant_category_codes: [{code: x}]}]}, single_use: {settlement_amount: {comparison: equals, value: 0}}}",
			"--billing-address.city", "x",
			"--billing-address.line1", "x",
			"--billing-address.postal-code", "x",
			"--billing-address.state", "x",
			"--billing-address.line2", "x",
			"--description", "New description",
			"--digital-wallet.digital-card-profile-id", "digital_card_profile_id",
			"--digital-wallet.email", "dev@stainless.com",
			"--digital-wallet.phone", "x",
			"--entity-id", "entity_id",
			"--status", "active",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"authorization_controls:\n" +
			"  merchant_acceptor_identifier:\n" +
			"    allowed:\n" +
			"      - identifier: x\n" +
			"    blocked:\n" +
			"      - identifier: x\n" +
			"  merchant_category_code:\n" +
			"    allowed:\n" +
			"      - code: '0253'\n" +
			"    blocked:\n" +
			"      - code: '0253'\n" +
			"  merchant_country:\n" +
			"    allowed:\n" +
			"      - country: xx\n" +
			"    blocked:\n" +
			"      - country: xx\n" +
			"  usage:\n" +
			"    category: single_use\n" +
			"    multi_use:\n" +
			"      spending_limits:\n" +
			"        - interval: all_time\n" +
			"          settlement_amount: 0\n" +
			"          merchant_category_codes:\n" +
			"            - code: x\n" +
			"    single_use:\n" +
			"      settlement_amount:\n" +
			"        comparison: equals\n" +
			"        value: 0\n" +
			"billing_address:\n" +
			"  city: x\n" +
			"  line1: x\n" +
			"  postal_code: x\n" +
			"  state: x\n" +
			"  line2: x\n" +
			"description: New description\n" +
			"digital_wallet:\n" +
			"  digital_card_profile_id: digital_card_profile_id\n" +
			"  email: dev@stainless.com\n" +
			"  phone: x\n" +
			"entity_id: entity_id\n" +
			"status: active\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cards", "update",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
		)
	})
}

func TestCardsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [active]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(cardsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[active]",
		)
	})
}

func TestCardsCreateDetailsIframe(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "create-details-iframe",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
			"--physical-card-id", "physical_card_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("physical_card_id: physical_card_id")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cards", "create-details-iframe",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
		)
	})
}

func TestCardsDetails(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "details",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
		)
	})
}

func TestCardsUpdatePin(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "update-pin",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
			"--pin", "1234",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("pin: '1234'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cards", "update-pin",
			"--card-id", "card_oubs0hwk5rn6knuecxg2",
		)
	})
}
