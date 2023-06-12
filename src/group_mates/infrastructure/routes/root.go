package group_mate_routes

import (
	"github.com/gofiber/fiber/v2"
	group_mate_controller "github.com/juanfer2/whorship_band/src/group_mates/infrastructure/controllers"
)

func SetupRoutes(app *fiber.App) {
	groupMateController := group_mate_controller.GroupMateControllerFactory()

	app.Get("/v1/group_mates", groupMateController.AllGroupMate)
	app.Post("/v1/group_mates", groupMateController.CreateGroupMate)
	app.Get("/v1/group_mates/:id", groupMateController.ShowGroupMate)
	app.Put("/v1/group_mates/:id", groupMateController.UpdateGroupMate)
	app.Delete("/v1/group_mates/:id", groupMateController.DeleteGroupMate)
}
