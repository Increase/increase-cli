// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Increase/increase-cli/internal/mocktest"
	"github.com/Increase/increase-cli/internal/requestflag"
)

func TestSimulationsCardPurchaseSupplementsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-purchase-supplements", "create",
			"--transaction-id", "transaction_uyrp7fld2ium70oa7oi",
			"--invoice", "{discount_amount: 100, duty_tax_amount: 200, order_date: '2023-07-20', shipping_amount: 300, shipping_destination_country_code: US, shipping_destination_postal_code: '10045', shipping_source_postal_code: '10045', shipping_tax_amount: 400, shipping_tax_rate: '0.2', unique_value_added_tax_invoice_reference: '12302'}",
			"--line-item", "{discount_amount: 0, item_commodity_code: '001', item_descriptor: Coffee, item_quantity: '1', product_code: '101', sales_tax_amount: 0, sales_tax_rate: '-16699', total_amount: 500, unit_cost: '5', unit_of_measure_code: NMB}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simulationsCardPurchaseSupplementsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"simulations:card-purchase-supplements", "create",
			"--transaction-id", "transaction_uyrp7fld2ium70oa7oi",
			"--invoice.discount-amount", "100",
			"--invoice.duty-tax-amount", "200",
			"--invoice.order-date", "2023-07-20",
			"--invoice.shipping-amount", "300",
			"--invoice.shipping-destination-country-code", "US",
			"--invoice.shipping-destination-postal-code", "10045",
			"--invoice.shipping-source-postal-code", "10045",
			"--invoice.shipping-tax-amount", "400",
			"--invoice.shipping-tax-rate", "0.2",
			"--invoice.unique-value-added-tax-invoice-reference", "12302",
			"--line-item.discount-amount", "0",
			"--line-item.item-commodity-code", "001",
			"--line-item.item-descriptor", "Coffee",
			"--line-item.item-quantity", "1",
			"--line-item.product-code", "101",
			"--line-item.sales-tax-amount", "0",
			"--line-item.sales-tax-rate", "-16699",
			"--line-item.total-amount", "500",
			"--line-item.unit-cost", "5",
			"--line-item.unit-of-measure-code", "NMB",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"transaction_id: transaction_uyrp7fld2ium70oa7oi\n" +
			"invoice:\n" +
			"  discount_amount: 100\n" +
			"  duty_tax_amount: 200\n" +
			"  order_date: '2023-07-20'\n" +
			"  shipping_amount: 300\n" +
			"  shipping_destination_country_code: US\n" +
			"  shipping_destination_postal_code: '10045'\n" +
			"  shipping_source_postal_code: '10045'\n" +
			"  shipping_tax_amount: 400\n" +
			"  shipping_tax_rate: '0.2'\n" +
			"  unique_value_added_tax_invoice_reference: '12302'\n" +
			"line_items:\n" +
			"  - discount_amount: 0\n" +
			"    item_commodity_code: '001'\n" +
			"    item_descriptor: Coffee\n" +
			"    item_quantity: '1'\n" +
			"    product_code: '101'\n" +
			"    sales_tax_amount: 0\n" +
			"    sales_tax_rate: '-16699'\n" +
			"    total_amount: 500\n" +
			"    unit_cost: '5'\n" +
			"    unit_of_measure_code: NMB\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"simulations:card-purchase-supplements", "create",
		)
	})
}
