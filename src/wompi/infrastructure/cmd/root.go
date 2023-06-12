package wompi_cmd

import (
	"github.com/google/uuid"
	wompi_client "github.com/juanfer2/whorship_band/src/shared/infrastructure/clients/wompi"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
	wompi_application "github.com/juanfer2/whorship_band/src/wompi/application"
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
	wompi_repositories "github.com/juanfer2/whorship_band/src/wompi/infrastructure/repositories/wompi_repository"
	"github.com/spf13/cobra"
)

var PaymentCmd = &cobra.Command{
	Use:     "pay",
	Aliases: []string{"p"},
	Short:   "Pay with differents payments ways",
}

func init() {
	PaymentCmd.AddCommand(CreditCardSourceCmd, PSECmd)
}

var CreditCardSourceCmd = &cobra.Command{
	Use:     "credit_card:test",
	Aliases: []string{"ct"},
	Short:   "pay with credit card",
	Run:     creditCardSource,
}

func creditCardSource(cmd *cobra.Command, args []string) {
	wompiRepository := wompi_repositories.NewWompiRepository()
	wompiPayment := wompi_application.NewExecuteTransaction(wompiRepository)
	payCreditCard := wompi_entities.NewCreditCardSourcePayment(
		wompi_entities.Payment{
			AmountInCents: 1500000,
			Currency:      "COP",
			CustomerEmail: "juan.villadiego@wompi.com",
			Reference:     "test_reference_" + uuid.New().String(),
			CustomerData: wompi_entities.CustomerData{
				PhoneNumber: "+573005138128",
				FullName:    "Juan Fernando Villadiego",
			},
		},
		wompi_entities.CreditCardSource{
			Installments: 1,
			Token:        "tok_stagint_2112_14bFa2c93aba2E9cA5ccD1ba5719d70a",
		},
	)
	response := wompiPayment.PayWithSourceCreditCard(*payCreditCard)
	utilities.PrintJson(response)
}

var PSECmd = &cobra.Command{
	Use:     "pse:test",
	Aliases: []string{"pt"},
	Short:   "pay with pse",
	Run:     pse,
}

func pse(cmd *cobra.Command, args []string) {
	wompiRepository := wompi_repositories.NewWompiRepository()
	wompiPayment := wompi_application.NewExecuteTransaction(wompiRepository)
	payPSE := wompi_entities.NewPSEPayment(
		wompi_entities.Payment{
			AmountInCents: 1500000,
			Currency:      "COP",
			CustomerEmail: "juan.villadiego@wompi.com",
			Reference:     "test_reference_" + uuid.New().String(),
			CustomerData: wompi_entities.CustomerData{
				PhoneNumber: "+573005138128",
				FullName:    "Juan Fernando Villadiego",
			},
		},
		wompi_entities.PSEProduct{
			Type:                     wompi_client.PSE.String(),
			UserType:                 0,
			UserLegalID:              "1999888777",
			UserLegalIDType:          "CC",
			FinancialInstitutionCode: "1",
			PaymentDescription:       "Pago a Tienda Wompi from CLI",
		},
	)

	response := wompiPayment.PayWithPSE(*payPSE)
	utilities.PrintJson(response)
}
