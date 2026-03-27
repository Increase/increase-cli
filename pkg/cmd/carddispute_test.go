// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestCardDisputesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-disputes", "create",
			"--disputed-transaction-id", "transaction_uyrp7fld2ium70oa7oi",
			"--network", "visa",
			"--amount", "100",
			"--attachment-file", "{file_id: file_id}",
			"--explanation", "x",
			"--visa", "{category: fraud, authorization: {account_status: account_closed}, consumer_canceled_merchandise: {merchant_resolution_attempted: attempted, purchase_explanation: x, received_or_expected_at: '2019-12-27', return_outcome: not_returned, cardholder_cancellation: {canceled_at: '2019-12-27', canceled_prior_to_ship_date: canceled_prior_to_ship_date, cancellation_policy_provided: not_provided, reason: x}, not_returned: {}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_canceled_recurring_transaction: {cancellation_target: account, merchant_contact_methods: {application_name: x, call_center_phone_number: x, email_address: x, in_person_address: x, mailing_address: x, text_phone_number: x}, transaction_or_account_canceled_at: '2019-12-27', other_form_of_payment_explanation: x}, consumer_canceled_services: {cardholder_cancellation: {canceled_at: '2019-12-27', cancellation_policy_provided: not_provided, reason: x}, contracted_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_explanation: x, service_type: guaranteed_reservation, guaranteed_reservation: {explanation: cardholder_canceled_prior_to_service}, other: {}, timeshare: {}}, consumer_counterfeit_merchandise: {counterfeit_explanation: x, disposition_explanation: x, order_explanation: x, received_at: '2019-12-27'}, consumer_credit_not_processed: {canceled_or_returned_at: '2019-12-27', credit_expected_at: '2019-12-27'}, consumer_damaged_or_defective_merchandise: {merchant_resolution_attempted: attempted, order_and_issue_explanation: x, received_at: '2019-12-27', return_outcome: not_returned, not_returned: {}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_merchandise_misrepresentation: {merchant_resolution_attempted: attempted, misrepresentation_explanation: x, purchase_explanation: x, received_at: '2019-12-27', return_outcome: not_returned, not_returned: {}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_merchandise_not_as_described: {merchant_resolution_attempted: attempted, received_at: '2019-12-27', return_outcome: returned, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_merchandise_not_received: {cancellation_outcome: cardholder_cancellation_prior_to_expected_receipt, delivery_issue: delayed, last_expected_receipt_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_info_and_explanation: x, cardholder_cancellation_prior_to_expected_receipt: {canceled_at: '2019-12-27', reason: x}, delayed: {explanation: x, return_outcome: not_returned, not_returned: {}, return_attempted: {attempted_at: '2019-12-27'}, returned: {merchant_received_return_at: '2019-12-27', returned_at: '2019-12-27'}}, delivered_to_wrong_location: {agreed_location: x}, merchant_cancellation: {canceled_at: '2019-12-27'}, no_cancellation: {}}, consumer_non_receipt_of_cash: {}, consumer_original_credit_transaction_not_accepted: {explanation: x, reason: prohibited_by_local_laws_or_regulation}, consumer_quality_merchandise: {expected_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_info_and_quality_issue: x, received_at: '2019-12-27', return_outcome: not_returned, not_returned: {}, ongoing_negotiations: {explanation: x, issuer_first_notified_at: '2019-12-27', started_at: '2019-12-27'}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_quality_services: {cardholder_cancellation: {accepted_by_merchant: accepted, canceled_at: '2019-12-27', reason: x}, non_fiat_currency_or_non_fungible_token_related_and_not_matching_description: not_related, purchase_info_and_quality_issue: x, services_received_at: '2019-12-27', cardholder_paid_to_have_work_redone: did_not_pay_to_have_work_redone, ongoing_negotiations: {explanation: x, issuer_first_notified_at: '2019-12-27', started_at: '2019-12-27'}, restaurant_food_related: not_related}, consumer_services_misrepresentation: {cardholder_cancellation: {accepted_by_merchant: accepted, canceled_at: '2019-12-27', reason: x}, merchant_resolution_attempted: attempted, misrepresentation_explanation: x, purchase_explanation: x, received_at: '2019-12-27'}, consumer_services_not_as_described: {cardholder_cancellation: {accepted_by_merchant: accepted, canceled_at: '2019-12-27', reason: x}, explanation: x, merchant_resolution_attempted: attempted, received_at: '2019-12-27'}, consumer_services_not_received: {cancellation_outcome: cardholder_cancellation_prior_to_expected_receipt, last_expected_receipt_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_info_and_explanation: x, cardholder_cancellation_prior_to_expected_receipt: {canceled_at: '2019-12-27', reason: x}, merchant_cancellation: {canceled_at: '2019-12-27'}, no_cancellation: {}}, fraud: {fraud_type: account_or_credentials_takeover}, processing_error: {error_reason: duplicate_transaction, merchant_resolution_attempted: attempted, duplicate_transaction: {other_transaction_id: x}, incorrect_amount: {expected_amount: 0}, paid_by_other_means: {other_form_of_payment_evidence: canceled_check, other_transaction_id: x}}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(cardDisputesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-disputes", "create",
			"--disputed-transaction-id", "transaction_uyrp7fld2ium70oa7oi",
			"--network", "visa",
			"--amount", "100",
			"--attachment-file.file-id", "file_id",
			"--explanation", "x",
			"--visa.category", "fraud",
			"--visa.authorization", "{account_status: account_closed}",
			"--visa.consumer-canceled-merchandise", "{merchant_resolution_attempted: attempted, purchase_explanation: x, received_or_expected_at: '2019-12-27', return_outcome: not_returned, cardholder_cancellation: {canceled_at: '2019-12-27', canceled_prior_to_ship_date: canceled_prior_to_ship_date, cancellation_policy_provided: not_provided, reason: x}, not_returned: {}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}",
			"--visa.consumer-canceled-recurring-transaction", "{cancellation_target: account, merchant_contact_methods: {application_name: x, call_center_phone_number: x, email_address: x, in_person_address: x, mailing_address: x, text_phone_number: x}, transaction_or_account_canceled_at: '2019-12-27', other_form_of_payment_explanation: x}",
			"--visa.consumer-canceled-services", "{cardholder_cancellation: {canceled_at: '2019-12-27', cancellation_policy_provided: not_provided, reason: x}, contracted_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_explanation: x, service_type: guaranteed_reservation, guaranteed_reservation: {explanation: cardholder_canceled_prior_to_service}, other: {}, timeshare: {}}",
			"--visa.consumer-counterfeit-merchandise", "{counterfeit_explanation: x, disposition_explanation: x, order_explanation: x, received_at: '2019-12-27'}",
			"--visa.consumer-credit-not-processed", "{canceled_or_returned_at: '2019-12-27', credit_expected_at: '2019-12-27'}",
			"--visa.consumer-damaged-or-defective-merchandise", "{merchant_resolution_attempted: attempted, order_and_issue_explanation: x, received_at: '2019-12-27', return_outcome: not_returned, not_returned: {}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}",
			"--visa.consumer-merchandise-misrepresentation", "{merchant_resolution_attempted: attempted, misrepresentation_explanation: x, purchase_explanation: x, received_at: '2019-12-27', return_outcome: not_returned, not_returned: {}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}",
			"--visa.consumer-merchandise-not-as-described", "{merchant_resolution_attempted: attempted, received_at: '2019-12-27', return_outcome: returned, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}",
			"--visa.consumer-merchandise-not-received", "{cancellation_outcome: cardholder_cancellation_prior_to_expected_receipt, delivery_issue: delayed, last_expected_receipt_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_info_and_explanation: x, cardholder_cancellation_prior_to_expected_receipt: {canceled_at: '2019-12-27', reason: x}, delayed: {explanation: x, return_outcome: not_returned, not_returned: {}, return_attempted: {attempted_at: '2019-12-27'}, returned: {merchant_received_return_at: '2019-12-27', returned_at: '2019-12-27'}}, delivered_to_wrong_location: {agreed_location: x}, merchant_cancellation: {canceled_at: '2019-12-27'}, no_cancellation: {}}",
			"--visa.consumer-non-receipt-of-cash", "{}",
			"--visa.consumer-original-credit-transaction-not-accepted", "{explanation: x, reason: prohibited_by_local_laws_or_regulation}",
			"--visa.consumer-quality-merchandise", "{expected_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_info_and_quality_issue: x, received_at: '2019-12-27', return_outcome: not_returned, not_returned: {}, ongoing_negotiations: {explanation: x, issuer_first_notified_at: '2019-12-27', started_at: '2019-12-27'}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}",
			"--visa.consumer-quality-services", "{cardholder_cancellation: {accepted_by_merchant: accepted, canceled_at: '2019-12-27', reason: x}, non_fiat_currency_or_non_fungible_token_related_and_not_matching_description: not_related, purchase_info_and_quality_issue: x, services_received_at: '2019-12-27', cardholder_paid_to_have_work_redone: did_not_pay_to_have_work_redone, ongoing_negotiations: {explanation: x, issuer_first_notified_at: '2019-12-27', started_at: '2019-12-27'}, restaurant_food_related: not_related}",
			"--visa.consumer-services-misrepresentation", "{cardholder_cancellation: {accepted_by_merchant: accepted, canceled_at: '2019-12-27', reason: x}, merchant_resolution_attempted: attempted, misrepresentation_explanation: x, purchase_explanation: x, received_at: '2019-12-27'}",
			"--visa.consumer-services-not-as-described", "{cardholder_cancellation: {accepted_by_merchant: accepted, canceled_at: '2019-12-27', reason: x}, explanation: x, merchant_resolution_attempted: attempted, received_at: '2019-12-27'}",
			"--visa.consumer-services-not-received", "{cancellation_outcome: cardholder_cancellation_prior_to_expected_receipt, last_expected_receipt_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_info_and_explanation: x, cardholder_cancellation_prior_to_expected_receipt: {canceled_at: '2019-12-27', reason: x}, merchant_cancellation: {canceled_at: '2019-12-27'}, no_cancellation: {}}",
			"--visa.fraud", "{fraud_type: account_or_credentials_takeover}",
			"--visa.processing-error", "{error_reason: duplicate_transaction, merchant_resolution_attempted: attempted, duplicate_transaction: {other_transaction_id: x}, incorrect_amount: {expected_amount: 0}, paid_by_other_means: {other_form_of_payment_evidence: canceled_check, other_transaction_id: x}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"disputed_transaction_id: transaction_uyrp7fld2ium70oa7oi\n" +
			"network: visa\n" +
			"amount: 100\n" +
			"attachment_files:\n" +
			"  - file_id: file_id\n" +
			"explanation: x\n" +
			"visa:\n" +
			"  category: fraud\n" +
			"  authorization:\n" +
			"    account_status: account_closed\n" +
			"  consumer_canceled_merchandise:\n" +
			"    merchant_resolution_attempted: attempted\n" +
			"    purchase_explanation: x\n" +
			"    received_or_expected_at: '2019-12-27'\n" +
			"    return_outcome: not_returned\n" +
			"    cardholder_cancellation:\n" +
			"      canceled_at: '2019-12-27'\n" +
			"      canceled_prior_to_ship_date: canceled_prior_to_ship_date\n" +
			"      cancellation_policy_provided: not_provided\n" +
			"      reason: x\n" +
			"    not_returned: {}\n" +
			"    return_attempted:\n" +
			"      attempt_explanation: x\n" +
			"      attempt_reason: merchant_not_responding\n" +
			"      attempted_at: '2019-12-27'\n" +
			"      merchandise_disposition: x\n" +
			"    returned:\n" +
			"      return_method: dhl\n" +
			"      returned_at: '2019-12-27'\n" +
			"      merchant_received_return_at: '2019-12-27'\n" +
			"      other_explanation: x\n" +
			"      tracking_number: x\n" +
			"  consumer_canceled_recurring_transaction:\n" +
			"    cancellation_target: account\n" +
			"    merchant_contact_methods:\n" +
			"      application_name: x\n" +
			"      call_center_phone_number: x\n" +
			"      email_address: x\n" +
			"      in_person_address: x\n" +
			"      mailing_address: x\n" +
			"      text_phone_number: x\n" +
			"    transaction_or_account_canceled_at: '2019-12-27'\n" +
			"    other_form_of_payment_explanation: x\n" +
			"  consumer_canceled_services:\n" +
			"    cardholder_cancellation:\n" +
			"      canceled_at: '2019-12-27'\n" +
			"      cancellation_policy_provided: not_provided\n" +
			"      reason: x\n" +
			"    contracted_at: '2019-12-27'\n" +
			"    merchant_resolution_attempted: attempted\n" +
			"    purchase_explanation: x\n" +
			"    service_type: guaranteed_reservation\n" +
			"    guaranteed_reservation:\n" +
			"      explanation: cardholder_canceled_prior_to_service\n" +
			"    other: {}\n" +
			"    timeshare: {}\n" +
			"  consumer_counterfeit_merchandise:\n" +
			"    counterfeit_explanation: x\n" +
			"    disposition_explanation: x\n" +
			"    order_explanation: x\n" +
			"    received_at: '2019-12-27'\n" +
			"  consumer_credit_not_processed:\n" +
			"    canceled_or_returned_at: '2019-12-27'\n" +
			"    credit_expected_at: '2019-12-27'\n" +
			"  consumer_damaged_or_defective_merchandise:\n" +
			"    merchant_resolution_attempted: attempted\n" +
			"    order_and_issue_explanation: x\n" +
			"    received_at: '2019-12-27'\n" +
			"    return_outcome: not_returned\n" +
			"    not_returned: {}\n" +
			"    return_attempted:\n" +
			"      attempt_explanation: x\n" +
			"      attempt_reason: merchant_not_responding\n" +
			"      attempted_at: '2019-12-27'\n" +
			"      merchandise_disposition: x\n" +
			"    returned:\n" +
			"      return_method: dhl\n" +
			"      returned_at: '2019-12-27'\n" +
			"      merchant_received_return_at: '2019-12-27'\n" +
			"      other_explanation: x\n" +
			"      tracking_number: x\n" +
			"  consumer_merchandise_misrepresentation:\n" +
			"    merchant_resolution_attempted: attempted\n" +
			"    misrepresentation_explanation: x\n" +
			"    purchase_explanation: x\n" +
			"    received_at: '2019-12-27'\n" +
			"    return_outcome: not_returned\n" +
			"    not_returned: {}\n" +
			"    return_attempted:\n" +
			"      attempt_explanation: x\n" +
			"      attempt_reason: merchant_not_responding\n" +
			"      attempted_at: '2019-12-27'\n" +
			"      merchandise_disposition: x\n" +
			"    returned:\n" +
			"      return_method: dhl\n" +
			"      returned_at: '2019-12-27'\n" +
			"      merchant_received_return_at: '2019-12-27'\n" +
			"      other_explanation: x\n" +
			"      tracking_number: x\n" +
			"  consumer_merchandise_not_as_described:\n" +
			"    merchant_resolution_attempted: attempted\n" +
			"    received_at: '2019-12-27'\n" +
			"    return_outcome: returned\n" +
			"    return_attempted:\n" +
			"      attempt_explanation: x\n" +
			"      attempt_reason: merchant_not_responding\n" +
			"      attempted_at: '2019-12-27'\n" +
			"      merchandise_disposition: x\n" +
			"    returned:\n" +
			"      return_method: dhl\n" +
			"      returned_at: '2019-12-27'\n" +
			"      merchant_received_return_at: '2019-12-27'\n" +
			"      other_explanation: x\n" +
			"      tracking_number: x\n" +
			"  consumer_merchandise_not_received:\n" +
			"    cancellation_outcome: cardholder_cancellation_prior_to_expected_receipt\n" +
			"    delivery_issue: delayed\n" +
			"    last_expected_receipt_at: '2019-12-27'\n" +
			"    merchant_resolution_attempted: attempted\n" +
			"    purchase_info_and_explanation: x\n" +
			"    cardholder_cancellation_prior_to_expected_receipt:\n" +
			"      canceled_at: '2019-12-27'\n" +
			"      reason: x\n" +
			"    delayed:\n" +
			"      explanation: x\n" +
			"      return_outcome: not_returned\n" +
			"      not_returned: {}\n" +
			"      return_attempted:\n" +
			"        attempted_at: '2019-12-27'\n" +
			"      returned:\n" +
			"        merchant_received_return_at: '2019-12-27'\n" +
			"        returned_at: '2019-12-27'\n" +
			"    delivered_to_wrong_location:\n" +
			"      agreed_location: x\n" +
			"    merchant_cancellation:\n" +
			"      canceled_at: '2019-12-27'\n" +
			"    no_cancellation: {}\n" +
			"  consumer_non_receipt_of_cash: {}\n" +
			"  consumer_original_credit_transaction_not_accepted:\n" +
			"    explanation: x\n" +
			"    reason: prohibited_by_local_laws_or_regulation\n" +
			"  consumer_quality_merchandise:\n" +
			"    expected_at: '2019-12-27'\n" +
			"    merchant_resolution_attempted: attempted\n" +
			"    purchase_info_and_quality_issue: x\n" +
			"    received_at: '2019-12-27'\n" +
			"    return_outcome: not_returned\n" +
			"    not_returned: {}\n" +
			"    ongoing_negotiations:\n" +
			"      explanation: x\n" +
			"      issuer_first_notified_at: '2019-12-27'\n" +
			"      started_at: '2019-12-27'\n" +
			"    return_attempted:\n" +
			"      attempt_explanation: x\n" +
			"      attempt_reason: merchant_not_responding\n" +
			"      attempted_at: '2019-12-27'\n" +
			"      merchandise_disposition: x\n" +
			"    returned:\n" +
			"      return_method: dhl\n" +
			"      returned_at: '2019-12-27'\n" +
			"      merchant_received_return_at: '2019-12-27'\n" +
			"      other_explanation: x\n" +
			"      tracking_number: x\n" +
			"  consumer_quality_services:\n" +
			"    cardholder_cancellation:\n" +
			"      accepted_by_merchant: accepted\n" +
			"      canceled_at: '2019-12-27'\n" +
			"      reason: x\n" +
			"    non_fiat_currency_or_non_fungible_token_related_and_not_matching_description: not_related\n" +
			"    purchase_info_and_quality_issue: x\n" +
			"    services_received_at: '2019-12-27'\n" +
			"    cardholder_paid_to_have_work_redone: did_not_pay_to_have_work_redone\n" +
			"    ongoing_negotiations:\n" +
			"      explanation: x\n" +
			"      issuer_first_notified_at: '2019-12-27'\n" +
			"      started_at: '2019-12-27'\n" +
			"    restaurant_food_related: not_related\n" +
			"  consumer_services_misrepresentation:\n" +
			"    cardholder_cancellation:\n" +
			"      accepted_by_merchant: accepted\n" +
			"      canceled_at: '2019-12-27'\n" +
			"      reason: x\n" +
			"    merchant_resolution_attempted: attempted\n" +
			"    misrepresentation_explanation: x\n" +
			"    purchase_explanation: x\n" +
			"    received_at: '2019-12-27'\n" +
			"  consumer_services_not_as_described:\n" +
			"    cardholder_cancellation:\n" +
			"      accepted_by_merchant: accepted\n" +
			"      canceled_at: '2019-12-27'\n" +
			"      reason: x\n" +
			"    explanation: x\n" +
			"    merchant_resolution_attempted: attempted\n" +
			"    received_at: '2019-12-27'\n" +
			"  consumer_services_not_received:\n" +
			"    cancellation_outcome: cardholder_cancellation_prior_to_expected_receipt\n" +
			"    last_expected_receipt_at: '2019-12-27'\n" +
			"    merchant_resolution_attempted: attempted\n" +
			"    purchase_info_and_explanation: x\n" +
			"    cardholder_cancellation_prior_to_expected_receipt:\n" +
			"      canceled_at: '2019-12-27'\n" +
			"      reason: x\n" +
			"    merchant_cancellation:\n" +
			"      canceled_at: '2019-12-27'\n" +
			"    no_cancellation: {}\n" +
			"  fraud:\n" +
			"    fraud_type: account_or_credentials_takeover\n" +
			"  processing_error:\n" +
			"    error_reason: duplicate_transaction\n" +
			"    merchant_resolution_attempted: attempted\n" +
			"    duplicate_transaction:\n" +
			"      other_transaction_id: x\n" +
			"    incorrect_amount:\n" +
			"      expected_amount: 0\n" +
			"    paid_by_other_means:\n" +
			"      other_form_of_payment_evidence: canceled_check\n" +
			"      other_transaction_id: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"card-disputes", "create",
		)
	})
}

func TestCardDisputesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-disputes", "retrieve",
			"--card-dispute-id", "card_dispute_h9sc95nbl1cgltpp7men",
		)
	})
}

func TestCardDisputesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-disputes", "list",
			"--max-items", "10",
			"--created-at", "{after: '2019-12-27T18:11:19.117Z', before: '2019-12-27T18:11:19.117Z', on_or_after: '2019-12-27T18:11:19.117Z', on_or_before: '2019-12-27T18:11:19.117Z'}",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status", "{in: [user_submission_required]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(cardDisputesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-disputes", "list",
			"--max-items", "10",
			"--created-at.after", "2019-12-27T18:11:19.117Z",
			"--created-at.before", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-after", "2019-12-27T18:11:19.117Z",
			"--created-at.on-or-before", "2019-12-27T18:11:19.117Z",
			"--cursor", "cursor",
			"--idempotency-key", "x",
			"--limit", "1",
			"--status.in", "[user_submission_required]",
		)
	})
}

func TestCardDisputesSubmitUserSubmission(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-disputes", "submit-user-submission",
			"--card-dispute-id", "card_dispute_h9sc95nbl1cgltpp7men",
			"--network", "visa",
			"--amount", "1",
			"--attachment-file", "{file_id: file_id}",
			"--explanation", "x",
			"--visa", "{category: merchant_prearbitration_decline, chargeback: {category: authorization, authorization: {account_status: account_closed}, consumer_canceled_merchandise: {merchant_resolution_attempted: attempted, purchase_explanation: x, received_or_expected_at: '2019-12-27', return_outcome: not_returned, cardholder_cancellation: {canceled_at: '2019-12-27', canceled_prior_to_ship_date: canceled_prior_to_ship_date, cancellation_policy_provided: not_provided, reason: x}, not_returned: {}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_canceled_recurring_transaction: {cancellation_target: account, merchant_contact_methods: {application_name: x, call_center_phone_number: x, email_address: x, in_person_address: x, mailing_address: x, text_phone_number: x}, transaction_or_account_canceled_at: '2019-12-27', other_form_of_payment_explanation: x}, consumer_canceled_services: {cardholder_cancellation: {canceled_at: '2019-12-27', cancellation_policy_provided: not_provided, reason: x}, contracted_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_explanation: x, service_type: guaranteed_reservation, guaranteed_reservation: {explanation: cardholder_canceled_prior_to_service}, other: {}, timeshare: {}}, consumer_counterfeit_merchandise: {counterfeit_explanation: x, disposition_explanation: x, order_explanation: x, received_at: '2019-12-27'}, consumer_credit_not_processed: {canceled_or_returned_at: '2019-12-27', credit_expected_at: '2019-12-27'}, consumer_damaged_or_defective_merchandise: {merchant_resolution_attempted: attempted, order_and_issue_explanation: x, received_at: '2019-12-27', return_outcome: not_returned, not_returned: {}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_merchandise_misrepresentation: {merchant_resolution_attempted: attempted, misrepresentation_explanation: x, purchase_explanation: x, received_at: '2019-12-27', return_outcome: not_returned, not_returned: {}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_merchandise_not_as_described: {merchant_resolution_attempted: attempted, received_at: '2019-12-27', return_outcome: returned, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_merchandise_not_received: {cancellation_outcome: cardholder_cancellation_prior_to_expected_receipt, delivery_issue: delayed, last_expected_receipt_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_info_and_explanation: x, cardholder_cancellation_prior_to_expected_receipt: {canceled_at: '2019-12-27', reason: x}, delayed: {explanation: x, return_outcome: not_returned, not_returned: {}, return_attempted: {attempted_at: '2019-12-27'}, returned: {merchant_received_return_at: '2019-12-27', returned_at: '2019-12-27'}}, delivered_to_wrong_location: {agreed_location: x}, merchant_cancellation: {canceled_at: '2019-12-27'}, no_cancellation: {}}, consumer_non_receipt_of_cash: {}, consumer_original_credit_transaction_not_accepted: {explanation: x, reason: prohibited_by_local_laws_or_regulation}, consumer_quality_merchandise: {expected_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_info_and_quality_issue: x, received_at: '2019-12-27', return_outcome: not_returned, not_returned: {}, ongoing_negotiations: {explanation: x, issuer_first_notified_at: '2019-12-27', started_at: '2019-12-27'}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_quality_services: {cardholder_cancellation: {accepted_by_merchant: accepted, canceled_at: '2019-12-27', reason: x}, non_fiat_currency_or_non_fungible_token_related_and_not_matching_description: not_related, purchase_info_and_quality_issue: x, services_received_at: '2019-12-27', cardholder_paid_to_have_work_redone: did_not_pay_to_have_work_redone, ongoing_negotiations: {explanation: x, issuer_first_notified_at: '2019-12-27', started_at: '2019-12-27'}, restaurant_food_related: not_related}, consumer_services_misrepresentation: {cardholder_cancellation: {accepted_by_merchant: accepted, canceled_at: '2019-12-27', reason: x}, merchant_resolution_attempted: attempted, misrepresentation_explanation: x, purchase_explanation: x, received_at: '2019-12-27'}, consumer_services_not_as_described: {cardholder_cancellation: {accepted_by_merchant: accepted, canceled_at: '2019-12-27', reason: x}, explanation: x, merchant_resolution_attempted: attempted, received_at: '2019-12-27'}, consumer_services_not_received: {cancellation_outcome: cardholder_cancellation_prior_to_expected_receipt, last_expected_receipt_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_info_and_explanation: x, cardholder_cancellation_prior_to_expected_receipt: {canceled_at: '2019-12-27', reason: x}, merchant_cancellation: {canceled_at: '2019-12-27'}, no_cancellation: {}}, fraud: {fraud_type: account_or_credentials_takeover}, processing_error: {error_reason: duplicate_transaction, merchant_resolution_attempted: attempted, duplicate_transaction: {other_transaction_id: x}, incorrect_amount: {expected_amount: 0}, paid_by_other_means: {other_form_of_payment_evidence: canceled_check, other_transaction_id: x}}}, merchant_prearbitration_decline: {reason: The pre-arbitration received from the merchantdoes not explain how they obtained permission to charge the card.}, user_prearbitration: {reason: x, category_change: {category: authorization, reason: x}}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(cardDisputesSubmitUserSubmission)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-disputes", "submit-user-submission",
			"--card-dispute-id", "card_dispute_h9sc95nbl1cgltpp7men",
			"--network", "visa",
			"--amount", "1",
			"--attachment-file.file-id", "file_id",
			"--explanation", "x",
			"--visa.category", "merchant_prearbitration_decline",
			"--visa.chargeback", "{category: authorization, authorization: {account_status: account_closed}, consumer_canceled_merchandise: {merchant_resolution_attempted: attempted, purchase_explanation: x, received_or_expected_at: '2019-12-27', return_outcome: not_returned, cardholder_cancellation: {canceled_at: '2019-12-27', canceled_prior_to_ship_date: canceled_prior_to_ship_date, cancellation_policy_provided: not_provided, reason: x}, not_returned: {}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_canceled_recurring_transaction: {cancellation_target: account, merchant_contact_methods: {application_name: x, call_center_phone_number: x, email_address: x, in_person_address: x, mailing_address: x, text_phone_number: x}, transaction_or_account_canceled_at: '2019-12-27', other_form_of_payment_explanation: x}, consumer_canceled_services: {cardholder_cancellation: {canceled_at: '2019-12-27', cancellation_policy_provided: not_provided, reason: x}, contracted_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_explanation: x, service_type: guaranteed_reservation, guaranteed_reservation: {explanation: cardholder_canceled_prior_to_service}, other: {}, timeshare: {}}, consumer_counterfeit_merchandise: {counterfeit_explanation: x, disposition_explanation: x, order_explanation: x, received_at: '2019-12-27'}, consumer_credit_not_processed: {canceled_or_returned_at: '2019-12-27', credit_expected_at: '2019-12-27'}, consumer_damaged_or_defective_merchandise: {merchant_resolution_attempted: attempted, order_and_issue_explanation: x, received_at: '2019-12-27', return_outcome: not_returned, not_returned: {}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_merchandise_misrepresentation: {merchant_resolution_attempted: attempted, misrepresentation_explanation: x, purchase_explanation: x, received_at: '2019-12-27', return_outcome: not_returned, not_returned: {}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_merchandise_not_as_described: {merchant_resolution_attempted: attempted, received_at: '2019-12-27', return_outcome: returned, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_merchandise_not_received: {cancellation_outcome: cardholder_cancellation_prior_to_expected_receipt, delivery_issue: delayed, last_expected_receipt_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_info_and_explanation: x, cardholder_cancellation_prior_to_expected_receipt: {canceled_at: '2019-12-27', reason: x}, delayed: {explanation: x, return_outcome: not_returned, not_returned: {}, return_attempted: {attempted_at: '2019-12-27'}, returned: {merchant_received_return_at: '2019-12-27', returned_at: '2019-12-27'}}, delivered_to_wrong_location: {agreed_location: x}, merchant_cancellation: {canceled_at: '2019-12-27'}, no_cancellation: {}}, consumer_non_receipt_of_cash: {}, consumer_original_credit_transaction_not_accepted: {explanation: x, reason: prohibited_by_local_laws_or_regulation}, consumer_quality_merchandise: {expected_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_info_and_quality_issue: x, received_at: '2019-12-27', return_outcome: not_returned, not_returned: {}, ongoing_negotiations: {explanation: x, issuer_first_notified_at: '2019-12-27', started_at: '2019-12-27'}, return_attempted: {attempt_explanation: x, attempt_reason: merchant_not_responding, attempted_at: '2019-12-27', merchandise_disposition: x}, returned: {return_method: dhl, returned_at: '2019-12-27', merchant_received_return_at: '2019-12-27', other_explanation: x, tracking_number: x}}, consumer_quality_services: {cardholder_cancellation: {accepted_by_merchant: accepted, canceled_at: '2019-12-27', reason: x}, non_fiat_currency_or_non_fungible_token_related_and_not_matching_description: not_related, purchase_info_and_quality_issue: x, services_received_at: '2019-12-27', cardholder_paid_to_have_work_redone: did_not_pay_to_have_work_redone, ongoing_negotiations: {explanation: x, issuer_first_notified_at: '2019-12-27', started_at: '2019-12-27'}, restaurant_food_related: not_related}, consumer_services_misrepresentation: {cardholder_cancellation: {accepted_by_merchant: accepted, canceled_at: '2019-12-27', reason: x}, merchant_resolution_attempted: attempted, misrepresentation_explanation: x, purchase_explanation: x, received_at: '2019-12-27'}, consumer_services_not_as_described: {cardholder_cancellation: {accepted_by_merchant: accepted, canceled_at: '2019-12-27', reason: x}, explanation: x, merchant_resolution_attempted: attempted, received_at: '2019-12-27'}, consumer_services_not_received: {cancellation_outcome: cardholder_cancellation_prior_to_expected_receipt, last_expected_receipt_at: '2019-12-27', merchant_resolution_attempted: attempted, purchase_info_and_explanation: x, cardholder_cancellation_prior_to_expected_receipt: {canceled_at: '2019-12-27', reason: x}, merchant_cancellation: {canceled_at: '2019-12-27'}, no_cancellation: {}}, fraud: {fraud_type: account_or_credentials_takeover}, processing_error: {error_reason: duplicate_transaction, merchant_resolution_attempted: attempted, duplicate_transaction: {other_transaction_id: x}, incorrect_amount: {expected_amount: 0}, paid_by_other_means: {other_form_of_payment_evidence: canceled_check, other_transaction_id: x}}}",
			"--visa.merchant-prearbitration-decline", "{reason: The pre-arbitration received from the merchantdoes not explain how they obtained permission to charge the card.}",
			"--visa.user-prearbitration", "{reason: x, category_change: {category: authorization, reason: x}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"network: visa\n" +
			"amount: 1\n" +
			"attachment_files:\n" +
			"  - file_id: file_id\n" +
			"explanation: x\n" +
			"visa:\n" +
			"  category: merchant_prearbitration_decline\n" +
			"  chargeback:\n" +
			"    category: authorization\n" +
			"    authorization:\n" +
			"      account_status: account_closed\n" +
			"    consumer_canceled_merchandise:\n" +
			"      merchant_resolution_attempted: attempted\n" +
			"      purchase_explanation: x\n" +
			"      received_or_expected_at: '2019-12-27'\n" +
			"      return_outcome: not_returned\n" +
			"      cardholder_cancellation:\n" +
			"        canceled_at: '2019-12-27'\n" +
			"        canceled_prior_to_ship_date: canceled_prior_to_ship_date\n" +
			"        cancellation_policy_provided: not_provided\n" +
			"        reason: x\n" +
			"      not_returned: {}\n" +
			"      return_attempted:\n" +
			"        attempt_explanation: x\n" +
			"        attempt_reason: merchant_not_responding\n" +
			"        attempted_at: '2019-12-27'\n" +
			"        merchandise_disposition: x\n" +
			"      returned:\n" +
			"        return_method: dhl\n" +
			"        returned_at: '2019-12-27'\n" +
			"        merchant_received_return_at: '2019-12-27'\n" +
			"        other_explanation: x\n" +
			"        tracking_number: x\n" +
			"    consumer_canceled_recurring_transaction:\n" +
			"      cancellation_target: account\n" +
			"      merchant_contact_methods:\n" +
			"        application_name: x\n" +
			"        call_center_phone_number: x\n" +
			"        email_address: x\n" +
			"        in_person_address: x\n" +
			"        mailing_address: x\n" +
			"        text_phone_number: x\n" +
			"      transaction_or_account_canceled_at: '2019-12-27'\n" +
			"      other_form_of_payment_explanation: x\n" +
			"    consumer_canceled_services:\n" +
			"      cardholder_cancellation:\n" +
			"        canceled_at: '2019-12-27'\n" +
			"        cancellation_policy_provided: not_provided\n" +
			"        reason: x\n" +
			"      contracted_at: '2019-12-27'\n" +
			"      merchant_resolution_attempted: attempted\n" +
			"      purchase_explanation: x\n" +
			"      service_type: guaranteed_reservation\n" +
			"      guaranteed_reservation:\n" +
			"        explanation: cardholder_canceled_prior_to_service\n" +
			"      other: {}\n" +
			"      timeshare: {}\n" +
			"    consumer_counterfeit_merchandise:\n" +
			"      counterfeit_explanation: x\n" +
			"      disposition_explanation: x\n" +
			"      order_explanation: x\n" +
			"      received_at: '2019-12-27'\n" +
			"    consumer_credit_not_processed:\n" +
			"      canceled_or_returned_at: '2019-12-27'\n" +
			"      credit_expected_at: '2019-12-27'\n" +
			"    consumer_damaged_or_defective_merchandise:\n" +
			"      merchant_resolution_attempted: attempted\n" +
			"      order_and_issue_explanation: x\n" +
			"      received_at: '2019-12-27'\n" +
			"      return_outcome: not_returned\n" +
			"      not_returned: {}\n" +
			"      return_attempted:\n" +
			"        attempt_explanation: x\n" +
			"        attempt_reason: merchant_not_responding\n" +
			"        attempted_at: '2019-12-27'\n" +
			"        merchandise_disposition: x\n" +
			"      returned:\n" +
			"        return_method: dhl\n" +
			"        returned_at: '2019-12-27'\n" +
			"        merchant_received_return_at: '2019-12-27'\n" +
			"        other_explanation: x\n" +
			"        tracking_number: x\n" +
			"    consumer_merchandise_misrepresentation:\n" +
			"      merchant_resolution_attempted: attempted\n" +
			"      misrepresentation_explanation: x\n" +
			"      purchase_explanation: x\n" +
			"      received_at: '2019-12-27'\n" +
			"      return_outcome: not_returned\n" +
			"      not_returned: {}\n" +
			"      return_attempted:\n" +
			"        attempt_explanation: x\n" +
			"        attempt_reason: merchant_not_responding\n" +
			"        attempted_at: '2019-12-27'\n" +
			"        merchandise_disposition: x\n" +
			"      returned:\n" +
			"        return_method: dhl\n" +
			"        returned_at: '2019-12-27'\n" +
			"        merchant_received_return_at: '2019-12-27'\n" +
			"        other_explanation: x\n" +
			"        tracking_number: x\n" +
			"    consumer_merchandise_not_as_described:\n" +
			"      merchant_resolution_attempted: attempted\n" +
			"      received_at: '2019-12-27'\n" +
			"      return_outcome: returned\n" +
			"      return_attempted:\n" +
			"        attempt_explanation: x\n" +
			"        attempt_reason: merchant_not_responding\n" +
			"        attempted_at: '2019-12-27'\n" +
			"        merchandise_disposition: x\n" +
			"      returned:\n" +
			"        return_method: dhl\n" +
			"        returned_at: '2019-12-27'\n" +
			"        merchant_received_return_at: '2019-12-27'\n" +
			"        other_explanation: x\n" +
			"        tracking_number: x\n" +
			"    consumer_merchandise_not_received:\n" +
			"      cancellation_outcome: cardholder_cancellation_prior_to_expected_receipt\n" +
			"      delivery_issue: delayed\n" +
			"      last_expected_receipt_at: '2019-12-27'\n" +
			"      merchant_resolution_attempted: attempted\n" +
			"      purchase_info_and_explanation: x\n" +
			"      cardholder_cancellation_prior_to_expected_receipt:\n" +
			"        canceled_at: '2019-12-27'\n" +
			"        reason: x\n" +
			"      delayed:\n" +
			"        explanation: x\n" +
			"        return_outcome: not_returned\n" +
			"        not_returned: {}\n" +
			"        return_attempted:\n" +
			"          attempted_at: '2019-12-27'\n" +
			"        returned:\n" +
			"          merchant_received_return_at: '2019-12-27'\n" +
			"          returned_at: '2019-12-27'\n" +
			"      delivered_to_wrong_location:\n" +
			"        agreed_location: x\n" +
			"      merchant_cancellation:\n" +
			"        canceled_at: '2019-12-27'\n" +
			"      no_cancellation: {}\n" +
			"    consumer_non_receipt_of_cash: {}\n" +
			"    consumer_original_credit_transaction_not_accepted:\n" +
			"      explanation: x\n" +
			"      reason: prohibited_by_local_laws_or_regulation\n" +
			"    consumer_quality_merchandise:\n" +
			"      expected_at: '2019-12-27'\n" +
			"      merchant_resolution_attempted: attempted\n" +
			"      purchase_info_and_quality_issue: x\n" +
			"      received_at: '2019-12-27'\n" +
			"      return_outcome: not_returned\n" +
			"      not_returned: {}\n" +
			"      ongoing_negotiations:\n" +
			"        explanation: x\n" +
			"        issuer_first_notified_at: '2019-12-27'\n" +
			"        started_at: '2019-12-27'\n" +
			"      return_attempted:\n" +
			"        attempt_explanation: x\n" +
			"        attempt_reason: merchant_not_responding\n" +
			"        attempted_at: '2019-12-27'\n" +
			"        merchandise_disposition: x\n" +
			"      returned:\n" +
			"        return_method: dhl\n" +
			"        returned_at: '2019-12-27'\n" +
			"        merchant_received_return_at: '2019-12-27'\n" +
			"        other_explanation: x\n" +
			"        tracking_number: x\n" +
			"    consumer_quality_services:\n" +
			"      cardholder_cancellation:\n" +
			"        accepted_by_merchant: accepted\n" +
			"        canceled_at: '2019-12-27'\n" +
			"        reason: x\n" +
			"      non_fiat_currency_or_non_fungible_token_related_and_not_matching_description: not_related\n" +
			"      purchase_info_and_quality_issue: x\n" +
			"      services_received_at: '2019-12-27'\n" +
			"      cardholder_paid_to_have_work_redone: did_not_pay_to_have_work_redone\n" +
			"      ongoing_negotiations:\n" +
			"        explanation: x\n" +
			"        issuer_first_notified_at: '2019-12-27'\n" +
			"        started_at: '2019-12-27'\n" +
			"      restaurant_food_related: not_related\n" +
			"    consumer_services_misrepresentation:\n" +
			"      cardholder_cancellation:\n" +
			"        accepted_by_merchant: accepted\n" +
			"        canceled_at: '2019-12-27'\n" +
			"        reason: x\n" +
			"      merchant_resolution_attempted: attempted\n" +
			"      misrepresentation_explanation: x\n" +
			"      purchase_explanation: x\n" +
			"      received_at: '2019-12-27'\n" +
			"    consumer_services_not_as_described:\n" +
			"      cardholder_cancellation:\n" +
			"        accepted_by_merchant: accepted\n" +
			"        canceled_at: '2019-12-27'\n" +
			"        reason: x\n" +
			"      explanation: x\n" +
			"      merchant_resolution_attempted: attempted\n" +
			"      received_at: '2019-12-27'\n" +
			"    consumer_services_not_received:\n" +
			"      cancellation_outcome: cardholder_cancellation_prior_to_expected_receipt\n" +
			"      last_expected_receipt_at: '2019-12-27'\n" +
			"      merchant_resolution_attempted: attempted\n" +
			"      purchase_info_and_explanation: x\n" +
			"      cardholder_cancellation_prior_to_expected_receipt:\n" +
			"        canceled_at: '2019-12-27'\n" +
			"        reason: x\n" +
			"      merchant_cancellation:\n" +
			"        canceled_at: '2019-12-27'\n" +
			"      no_cancellation: {}\n" +
			"    fraud:\n" +
			"      fraud_type: account_or_credentials_takeover\n" +
			"    processing_error:\n" +
			"      error_reason: duplicate_transaction\n" +
			"      merchant_resolution_attempted: attempted\n" +
			"      duplicate_transaction:\n" +
			"        other_transaction_id: x\n" +
			"      incorrect_amount:\n" +
			"        expected_amount: 0\n" +
			"      paid_by_other_means:\n" +
			"        other_form_of_payment_evidence: canceled_check\n" +
			"        other_transaction_id: x\n" +
			"  merchant_prearbitration_decline:\n" +
			"    reason: >-\n" +
			"      The pre-arbitration received from the merchantdoes not explain how they\n" +
			"      obtained permission to charge the card.\n" +
			"  user_prearbitration:\n" +
			"    reason: x\n" +
			"    category_change:\n" +
			"      category: authorization\n" +
			"      reason: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"card-disputes", "submit-user-submission",
			"--card-dispute-id", "card_dispute_h9sc95nbl1cgltpp7men",
		)
	})
}

func TestCardDisputesWithdraw(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"card-disputes", "withdraw",
			"--card-dispute-id", "card_dispute_h9sc95nbl1cgltpp7men",
			"--explanation", "The explanation for withdrawing the Card Dispute.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("explanation: The explanation for withdrawing the Card Dispute.")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"card-disputes", "withdraw",
			"--card-dispute-id", "card_dispute_h9sc95nbl1cgltpp7men",
		)
	})
}
