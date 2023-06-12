package cmd

import (
	"time"

	"github.com/fatih/color"
	wompi_client "github.com/juanfer2/whorship_band/src/shared/infrastructure/clients/wompi"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
	"github.com/spf13/cobra"
)

var merchantCmd = &cobra.Command{
	Use:   "merchant",
	Short: "merchant service",
	Long:  `merchant service`,
	Run:   merchant,
}

func merchant(cmd *cobra.Command, args []string) {
	/*
			categoryPromptContent := promptContent{
				"Please provide a category.",
				fmt.Sprintf("What category does word belong to?"),
			}
			category := promptGetSelect(categoryPromptContent)

		color.Green(category)
	*/
	wompiClient := wompi_client.WompiClientFactory()
	color.Green("Search...")
	merchantObj := wompiClient.GetMerchant()
	utilities.PrintJson(merchantObj)
}

var pyamentTestPSECmd = &cobra.Command{
	Use:   "payment:pse:test",
	Short: "pay with pse test",
	Long:  `pay with pse test`,
	Run:   pyamentTestPSE,
}

func pyamentTestPSE(cmd *cobra.Command, args []string) {
	wompiClient := wompi_client.WompiClientFactory()
	payment := wompi_client.NewPaymentBody(
		PaymentBodyTestData(),
		wompi_client.PaymentWithPSE(wompi_client.PaymentMethodPSE{
			Type:                     wompi_client.PSE.String(),
			UserType:                 0,
			UserLegalID:              "1999888777",
			UserLegalIDType:          "CC",
			FinancialInstitutionCode: "1",
			PaymentDescription:       "Pago a Tienda Wompi from CLI",
		}),
	)
	response := wompiClient.CreateTransaction(payment)

	utilities.PrintJson(response)
}

var pyamentTestCreditCardCmd = &cobra.Command{
	Use:   "payment:credit_card:test",
	Short: "pay with credit card test",
	Long:  `pay with credit card test`,
	Run:   pyamentTestCreditCard,
}

func pyamentTestCreditCard(cmd *cobra.Command, args []string) {
	wompiClient := wompi_client.WompiClientFactory()
	payment := wompi_client.NewPaymentBody(
		PaymentBodyTestData(),
		wompi_client.PaymentWithCreditCard(wompi_client.PaymentMethodCreditCard{
			Type:         wompi_client.CREDIT_CARD.String(),
			Installments: 1,
			Token:        "tok_stagint_2112_14bFa2c93aba2E9cA5ccD1ba5719d70a",
		}),
	)
	response := wompiClient.CreateTransaction(payment)

	utilities.PrintJson(response)
}

func PaymentBodyTestData() wompi_client.ParamsPaymentBody {
	return wompi_client.ParamsPaymentBody{
		AmountInCents: 1500000,
		Currency:      "COP",
		CustomerEmail: "juan.villadiego@wompi.com",
		Reference:     "test_reference_" + string(time.Now().UnixNano()),
		CustomerData: wompi_client.CustomerData{
			PhoneNumber: "+573005138128",
			FullName:    "Juan Fernando Villadiego",
		},
	}
}
