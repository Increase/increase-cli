// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestEventSubscriptionsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"event-subscriptions", "create",
			"--url", "https://website.com/webhooks",
			"--oauth-connection-id", "x",
			"--selected-event-category", "{event_category: account.created}",
			"--shared-secret", "x",
			"--status", "active",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(eventSubscriptionsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"event-subscriptions", "create",
			"--url", "https://website.com/webhooks",
			"--oauth-connection-id", "x",
			"--selected-event-category.event-category", "account.created",
			"--shared-secret", "x",
			"--status", "active",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: https://website.com/webhooks\n" +
			"oauth_connection_id: x\n" +
			"selected_event_categories:\n" +
			"  - event_category: account.created\n" +
			"shared_secret: x\n" +
			"status: active\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"event-subscriptions", "create",
		)
	})
}

func TestEventSubscriptionsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"event-subscriptions", "retrieve",
			"--event-subscription-id", "event_subscription_001dzz0r20rcdxgb013zqb8m04g",
		)
	})
}

func TestEventSubscriptionsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"event-subscriptions", "update",
			"--event-subscription-id", "event_subscription_001dzz0r20rcdxgb013zqb8m04g",
			"--status", "active",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("status: active")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"event-subscriptions", "update",
			"--event-subscription-id", "event_subscription_001dzz0r20rcdxgb013zqb8m04g",
		)
	})
}

func TestEventSubscriptionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"event-subscriptions", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
		)
	})
}
