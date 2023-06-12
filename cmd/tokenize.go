package cmd

import (
	"fmt"

	wompi_client "github.com/juanfer2/whorship_band/src/shared/infrastructure/clients/wompi"
	wompi_products "github.com/juanfer2/whorship_band/src/shared/infrastructure/clients/wompi/products"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
	"github.com/spf13/cobra"
)

var tokenizeCreditCardCmd = &cobra.Command{
	Use:   "tokenize:credit_card:test",
	Short: "tokenize credit card service",
	Long:  `tokenize credit card service`,
	Run:   tokenizeCreditCard,
}

func tokenizeCreditCard(cmd *cobra.Command, args []string) {
	wompiClient := wompi_client.WompiClientFactory()
	creditCard := wompi_client.TokenizeCreditCardOption(wompi_products.CreditCard{
		Number:     "4242424242424242",
		ExpMonth:   "08",
		ExpYear:    "25",
		CVC:        "777",
		CardHolder: "APPROVED",
	})
	tokenize := wompiClient.Tokenize(wompi_client.NewTokenizeParams(creditCard))
	fmt.Println("Token")
	utilities.PrintJson(tokenize)
}

var tokenizeNequiCmd = &cobra.Command{
	Use:   "tokenize:nequi:test",
	Short: "tokenize credit card service",
	Long:  `tokenize credit card service`,
	Run:   tokenizeNequi,
}

func tokenizeNequi(cmd *cobra.Command, args []string) {
	wompiClient := wompi_client.WompiClientFactory()
	creditCard := wompi_client.TokenizeNequiOption(wompi_products.Nequi{PhoneNumber: "3991111111"})
	tokenize := wompiClient.Tokenize(wompi_client.NewTokenizeParams(creditCard))
	utilities.PrintJson(tokenize)
}
