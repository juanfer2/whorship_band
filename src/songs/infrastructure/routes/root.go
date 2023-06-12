package songs_routes

import (
	"github.com/gofiber/fiber/v2"
	songs_controller "github.com/juanfer2/whorship_band/src/songs/infrastructure/controller"
)

func SetupRoutes(app *fiber.App) {
	songController := songs_controller.SongControllerFactory()

	app.Get("/v1/songs", songController.All)
	app.Post("/v1/songs", songController.Create)
	app.Get("/v1/songs/:id", songController.Show)
	app.Put("/v1/songs/:id", songController.Update)
	app.Delete("/v1/songs/:id", songController.Delete)
}
