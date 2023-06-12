package instrument_controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	intrument_domain "github.com/juanfer2/whorship_band/src/instruments/domain"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
)

// Struct: instrumentParams
type instrumentParams struct {
	Name string `json:"name"`
}

func getParamsInstrument(c *fiber.Ctx) (intrument_domain.Instrument, error) {
	var GroupMateParams instrumentParams
	if err := c.BodyParser(&GroupMateParams); err != nil {
		return intrument_domain.Instrument{}, err
	}

	utilities.PrintJson(GroupMateParams)

	fmt.Println(GroupMateParams)

	return intrument_domain.Instrument{
		Name: GroupMateParams.Name,
	}, nil
}
