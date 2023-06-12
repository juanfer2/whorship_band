package instrument_routes

import (
	"github.com/gofiber/fiber/v2"

	instrument_controller "github.com/juanfer2/whorship_band/src/instruments/infrastructure/controller"
)

func SetupRoutes(app *fiber.App) {
	instrumentController := instrument_controller.InstrumentControllerFactory()

	app.Get("/v1/instruments", instrumentController.AllInstrument)
	app.Post("/v1/instruments", instrumentController.CreateInstrument)
	app.Put("/v1/instruments/:id", instrumentController.UpdateInstrument)
	app.Delete("/v1/instruments/:id", instrumentController.DeleteInstrument)
}
