package group_mate_controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	group_mate_domain "github.com/juanfer2/whorship_band/src/group_mates/domain"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
)

type GroupMateParams struct {
	Name        string                                  `json:"name"`
	URLImage    string                                  `json:"urlImage"`
	Instruments []group_mate_domain.GroupMateInstrument `json:"instruments"`
}

func getParamsGroupMate(c *fiber.Ctx) (group_mate_domain.GroupMate, error) {
	var GroupMateParams GroupMateParams
	if err := c.BodyParser(&GroupMateParams); err != nil {
		return group_mate_domain.GroupMate{}, err
	}

	utilities.PrintJson(GroupMateParams)

	fmt.Println(GroupMateParams)

	return group_mate_domain.GroupMate{
		Name:        GroupMateParams.Name,
		URLImage:    GroupMateParams.URLImage,
		Instruments: GroupMateParams.Instruments,
	}, nil
}
