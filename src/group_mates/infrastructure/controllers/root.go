package group_mate_controller

import (
	"github.com/gofiber/fiber/v2"

	group_mates_application "github.com/juanfer2/whorship_band/src/group_mates/application"
	group_mates_serializers "github.com/juanfer2/whorship_band/src/group_mates/infrastructure/serializers"
	middlewares "github.com/juanfer2/whorship_band/src/shared/infrastructure/controllers"
)

type GroupMateController struct {
	groupMateUseCase *group_mates_application.GroupMateUseCase
}

func NewGroupMateController(
	groupMateUseCase *group_mates_application.GroupMateUseCase,
) *GroupMateController {
	return &GroupMateController{groupMateUseCase: groupMateUseCase}
}

func (groupMateController *GroupMateController) AllGroupMate(c *fiber.Ctx) error {
	groupMates := groupMateController.groupMateUseCase.GetAll()

	return c.JSON(group_mates_serializers.NewGroupMateSerializerFromGroupMateList(groupMates))
}

func (groupMateController *GroupMateController) CreateGroupMate(c *fiber.Ctx) error {
	groupMateParams, err := getParamsGroupMate(c)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(403).JSON(errorHandler)
	}

	groupMate, err := groupMateController.groupMateUseCase.Create(groupMateParams)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	group_mates_serializers.NewGroupMateSerializerFromGroupMate(groupMate)

	return c.JSON(group_mates_serializers.NewGroupMateSerializerFromGroupMate(groupMate))
}

func (groupMateController *GroupMateController) ShowGroupMate(c *fiber.Ctx) error {
	id := c.Params("id")
	groupMate := groupMateController.groupMateUseCase.FindByUuid(id)

	return c.JSON(group_mates_serializers.NewGroupMateSerializerFromGroupMateList(groupMate))
}

func (groupMateController *GroupMateController) UpdateGroupMate(c *fiber.Ctx) error {
	id := c.Params("id")
	groupMateParams, err := getParamsGroupMate(c)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(403).JSON(errorHandler)
	}

	groupMate, err := groupMateController.groupMateUseCase.Update(id, groupMateParams)

	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(group_mates_serializers.NewGroupMateSerializerFromGroupMate(groupMate))
}

func (groupMateController *GroupMateController) DeleteGroupMate(c *fiber.Ctx) error {
	id := c.Params("id")
	groupMate, err := groupMateController.groupMateUseCase.Delete(id)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(group_mates_serializers.NewGroupMateSerializerFromGroupMate(groupMate))
}
