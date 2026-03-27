// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestRealTimeDecisionsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"real-time-decisions", "retrieve",
			"--real-time-decision-id", "real_time_decision_j76n2e810ezcg3zh5qtn",
		)
	})
}

func TestRealTimeDecisionsAction(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"real-time-decisions", "action",
			"--real-time-decision-id", "real_time_decision_j76n2e810ezcg3zh5qtn",
			"--card-authentication", "{decision: approve}",
			"--card-authentication-challenge", "{result: success, success: {email: dev@stainless.com, phone: x}}",
			"--card-authorization", "{decision: approve, approval: {cardholder_address_verification_result: {line1: match, postal_code: no_match}, partial_amount: 1}, decline: {reason: insufficient_funds}}",
			"--card-balance-inquiry", "{decision: approve, approval: {balance: 0}}",
			"--digital-wallet-authentication", "{result: success, success: {email: dev@stainless.com, phone: x}}",
			"--digital-wallet-token", "{approval: {email: dev@stainless.com, phone: x}, decline: {reason: x}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(realTimeDecisionsAction)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"real-time-decisions", "action",
			"--real-time-decision-id", "real_time_decision_j76n2e810ezcg3zh5qtn",
			"--card-authentication.decision", "approve",
			"--card-authentication-challenge.result", "success",
			"--card-authentication-challenge.success", "{email: dev@stainless.com, phone: x}",
			"--card-authorization.decision", "approve",
			"--card-authorization.approval", "{cardholder_address_verification_result: {line1: match, postal_code: no_match}, partial_amount: 1}",
			"--card-authorization.decline", "{reason: insufficient_funds}",
			"--card-balance-inquiry.decision", "approve",
			"--card-balance-inquiry.approval", "{balance: 0}",
			"--digital-wallet-authentication.result", "success",
			"--digital-wallet-authentication.success", "{email: dev@stainless.com, phone: x}",
			"--digital-wallet-token.approval", "{email: dev@stainless.com, phone: x}",
			"--digital-wallet-token.decline", "{reason: x}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"card_authentication:\n" +
			"  decision: approve\n" +
			"card_authentication_challenge:\n" +
			"  result: success\n" +
			"  success:\n" +
			"    email: dev@stainless.com\n" +
			"    phone: x\n" +
			"card_authorization:\n" +
			"  decision: approve\n" +
			"  approval:\n" +
			"    cardholder_address_verification_result:\n" +
			"      line1: match\n" +
			"      postal_code: no_match\n" +
			"    partial_amount: 1\n" +
			"  decline:\n" +
			"    reason: insufficient_funds\n" +
			"card_balance_inquiry:\n" +
			"  decision: approve\n" +
			"  approval:\n" +
			"    balance: 0\n" +
			"digital_wallet_authentication:\n" +
			"  result: success\n" +
			"  success:\n" +
			"    email: dev@stainless.com\n" +
			"    phone: x\n" +
			"digital_wallet_token:\n" +
			"  approval:\n" +
			"    email: dev@stainless.com\n" +
			"    phone: x\n" +
			"  decline:\n" +
			"    reason: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"real-time-decisions", "action",
			"--real-time-decision-id", "real_time_decision_j76n2e810ezcg3zh5qtn",
		)
	})
}
