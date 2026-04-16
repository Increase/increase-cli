// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/Increase/increase-cli/internal/apiquery"
	"github.com/Increase/increase-cli/internal/requestflag"
	"github.com/Increase/increase-go"
	"github.com/Increase/increase-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var simulationsCardPurchaseSupplementsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Simulates the creation of a Card Purchase Supplement (Level 3 data) for a card\nsettlement. This happens asynchronously in production when Visa sends enhanced\ntransaction data about a purchase.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "transaction-id",
			Usage:    "The identifier of the Transaction to create a Card Purchase Supplement for. The Transaction must have a source of type `card_settlement`.",
			Required: true,
			BodyPath: "transaction_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "invoice",
			Usage:    "Invoice-level information about the payment.",
			BodyPath: "invoice",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "line-item",
			Usage:    "Line item information, such as individual products purchased.",
			BodyPath: "line_items",
		},
	},
	Action:          handleSimulationsCardPurchaseSupplementsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"invoice": {
		&requestflag.InnerFlag[int64]{
			Name:       "invoice.discount-amount",
			Usage:      "Discount given to cardholder.",
			InnerField: "discount_amount",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "invoice.duty-tax-amount",
			Usage:      "Amount of duty taxes.",
			InnerField: "duty_tax_amount",
		},
		&requestflag.InnerFlag[any]{
			Name:       "invoice.order-date",
			Usage:      "Date the order was taken.",
			InnerField: "order_date",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "invoice.shipping-amount",
			Usage:      "The shipping cost.",
			InnerField: "shipping_amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "invoice.shipping-destination-country-code",
			Usage:      "Country code of the shipping destination.",
			InnerField: "shipping_destination_country_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "invoice.shipping-destination-postal-code",
			Usage:      "Postal code of the shipping destination.",
			InnerField: "shipping_destination_postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "invoice.shipping-source-postal-code",
			Usage:      "Postal code of the location being shipped from.",
			InnerField: "shipping_source_postal_code",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "invoice.shipping-tax-amount",
			Usage:      "Taxes paid for freight and shipping.",
			InnerField: "shipping_tax_amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "invoice.shipping-tax-rate",
			Usage:      "Tax rate for freight and shipping.",
			InnerField: "shipping_tax_rate",
		},
		&requestflag.InnerFlag[string]{
			Name:       "invoice.unique-value-added-tax-invoice-reference",
			Usage:      "Value added tax invoice reference number.",
			InnerField: "unique_value_added_tax_invoice_reference",
		},
	},
	"line-item": {
		&requestflag.InnerFlag[int64]{
			Name:       "line-item.discount-amount",
			Usage:      "Discount amount for this specific line item.",
			InnerField: "discount_amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "line-item.item-commodity-code",
			Usage:      "Code used to categorize the purchase item.",
			InnerField: "item_commodity_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "line-item.item-descriptor",
			Usage:      "Description of the purchase item.",
			InnerField: "item_descriptor",
		},
		&requestflag.InnerFlag[string]{
			Name:       "line-item.item-quantity",
			Usage:      "The number of units of the product being purchased.",
			InnerField: "item_quantity",
		},
		&requestflag.InnerFlag[string]{
			Name:       "line-item.product-code",
			Usage:      "Code used to categorize the product being purchased.",
			InnerField: "product_code",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "line-item.sales-tax-amount",
			Usage:      "Sales tax amount for this line item.",
			InnerField: "sales_tax_amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "line-item.sales-tax-rate",
			Usage:      "Sales tax rate for this line item.",
			InnerField: "sales_tax_rate",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "line-item.total-amount",
			Usage:      "Total amount of all line items.",
			InnerField: "total_amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "line-item.unit-cost",
			Usage:      "Cost of line item per unit of measure, in major units.",
			InnerField: "unit_cost",
		},
		&requestflag.InnerFlag[string]{
			Name:       "line-item.unit-of-measure-code",
			Usage:      "Code indicating unit of measure (gallons, etc.).",
			InnerField: "unit_of_measure_code",
		},
	},
})

func handleSimulationsCardPurchaseSupplementsCreate(ctx context.Context, cmd *cli.Command) error {
	client := increase.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := increase.SimulationCardPurchaseSupplementNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatDots,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simulations.CardPurchaseSupplements.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "simulations:card-purchase-supplements create",
		Transform:      transform,
	})
}
