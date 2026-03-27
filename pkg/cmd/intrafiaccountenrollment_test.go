// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestIntrafiAccountEnrollmentsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"intrafi-account-enrollments", "create",
			"--account-id", "account_in71c4amph0vgo2qllky",
			"--email-address", "user@example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"account_id: account_in71c4amph0vgo2qllky\n" +
			"email_address: user@example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"intrafi-account-enrollments", "create",
		)
	})
}

func TestIntrafiAccountEnrollmentsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"intrafi-account-enrollments", "retrieve",
			"--intrafi-account-enrollment-id", "intrafi_account_enrollment_w8l97znzreopkwf2tg75",
		)
	})
}

func TestIntrafiAccountEnrollmentsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"intrafi-account-enrollments", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [pending_enrolling]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(intrafiAccountEnrollmentsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"intrafi-account-enrollments", "list",
			"--max-items", "10",
			"--account-id", "account_id",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[pending_enrolling]",
		)
	})
}

func TestIntrafiAccountEnrollmentsUnenroll(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"intrafi-account-enrollments", "unenroll",
			"--intrafi-account-enrollment-id", "intrafi_account_enrollment_w8l97znzreopkwf2tg75",
		)
	})
}
