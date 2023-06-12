package performances_routes

import (
	"github.com/gofiber/fiber/v2"

	performances_controller "github.com/juanfer2/whorship_band/src/performances/infrastructure/controller"
)

func SetupRoutes(app *fiber.App) {
	performanceController := performances_controller.PerformanceControllerFactory()

	app.Get("/v1/performances", performanceController.All)
	app.Post("/v1/performances", performanceController.Create)
	app.Put("/v1/performances/:id", performanceController.Update)
	app.Delete("/v1/performances/:id", performanceController.Delete)
}
