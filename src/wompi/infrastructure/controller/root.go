package wompi_controller

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/juanfer2/whorship_band/src/shared/domain"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
	wompi_application "github.com/juanfer2/whorship_band/src/wompi/application"
	wompi_serializers "github.com/juanfer2/whorship_band/src/wompi/infrastructure/serializers"
)

type WompiController struct {
	tokenizeProduct          *wompi_application.TokenizeProduct
	productUseCase           *wompi_application.ProductUseCase
	extractCreditCardUseCase *wompi_application.ExtractCreditcCardFromImage
	executeTransaction       *wompi_application.ExecuteTransaction
}

func NewWompiController(
	tokenizeProduct *wompi_application.TokenizeProduct,
	productUseCase *wompi_application.ProductUseCase,
	extractCreditCardUseCase *wompi_application.ExtractCreditcCardFromImage,
	executeTransaction *wompi_application.ExecuteTransaction,
) *WompiController {
	return &WompiController{tokenizeProduct: tokenizeProduct, productUseCase: productUseCase,
		extractCreditCardUseCase: extractCreditCardUseCase, executeTransaction: executeTransaction}
}

func (wompiController *WompiController) AllProducts(c *fiber.Ctx) error {
	products := wompiController.productUseCase.All()

	return c.JSON(wompi_serializers.NewProductsSerializer(products))
}

func (wompiController *WompiController) Payment(c *fiber.Ctx) error {
	products := wompiController.productUseCase.All()

	return c.JSON(wompi_serializers.NewProductsSerializer(products))
}

func (wompiController *WompiController) TokenizeCreditCard(c *fiber.Ctx) error {
	creditCardInput, err := getParamsCreditCard(c)
	if err != nil {
		fmt.Println(err)
		return c.JSON(err)
	}

	product := wompiController.tokenizeProduct.TokenizeCreditCard(
		creditCardInput.CreditCard, creditCardInput.CustomerEmail,
	)

	return c.JSON(wompi_serializers.NewProductSerializer(product))
}

func (wompiController *WompiController) GetCreditCardByImage(c *fiber.Ctx) error {
	file, err := c.FormFile("creditCard")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Leemos el contenido del archivo
	content, err := file.Open()
	if err != nil {
		return c.SendString(fmt.Sprintf("Error al leer el archivo: %s", err))
	}
	defer content.Close()

	// Leemos el contenido del archivo y lo guardamos en un byte slice
	fileBytes, err := ioutil.ReadAll(content)
	if err != nil {
		return c.SendString(fmt.Sprintf("Error al leer el contenido del archivo: %s", err))
	}

	filename := utilities.GetRootFolder() + "/tmp/images/sources/" + file.Filename
	// Guardamos el archivo en el disco
	err = ioutil.WriteFile(filename, fileBytes, 0644)
	if err != nil {
		log.Fatal(err)
		return c.SendString(fmt.Sprintf("Error al guardar el archivo: %s", err))
	}

	creditCardImage := domain.File{
		Filename: filename,
		Header:   file.Header,
		Size:     int(file.Size),
	}

	creditCard := wompiController.extractCreditCardUseCase.GetCreditCard(creditCardImage)

	utilities.PrintJson(file)

	return c.JSON(wompi_serializers.CreditCardSerializer(creditCard))
}
