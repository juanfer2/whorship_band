package wompi_routes

import (
	"github.com/gofiber/fiber/v2"
	wompi_controller "github.com/juanfer2/whorship_band/src/wompi/infrastructure/controller"
)

func SetupRoutes(app *fiber.App) {
	wompiController := wompi_controller.FactoryWompiController()

	app.Post("/token/credit_card", wompiController.TokenizeCreditCard)
	app.Get("/products", wompiController.AllProducts)
	app.Post("/extract_credit_card", wompiController.GetCreditCardByImage)
}
