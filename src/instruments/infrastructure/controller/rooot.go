package instrument_controller

import (
	"github.com/gofiber/fiber/v2"

	instrument_application "github.com/juanfer2/whorship_band/src/instruments/application"
	instrument_serializer "github.com/juanfer2/whorship_band/src/instruments/infrastructure/serializer"
	middlewares "github.com/juanfer2/whorship_band/src/shared/infrastructure/controllers"
)

type InstrumentController struct {
	instrumentUseCase *instrument_application.InstrumentUseCase
}

func NewInstrumentController(
	instrumentUseCase *instrument_application.InstrumentUseCase,
) *InstrumentController {
	return &InstrumentController{instrumentUseCase: instrumentUseCase}
}

func (InstrumentController *InstrumentController) AllInstrument(c *fiber.Ctx) error {
	instruments := InstrumentController.instrumentUseCase.GetAll()
	return c.JSON(instrument_serializer.NewInstrumentSerializerInstrumentList(instruments))
}

func (InstrumentController *InstrumentController) CreateInstrument(c *fiber.Ctx) error {
	instrumentParams, err := getParamsInstrument(c)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(403).JSON(errorHandler)
	}

	instrument, err := InstrumentController.instrumentUseCase.Create(instrumentParams)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(instrument_serializer.NewInstrumentSerializerFromtInstrument(instrument))
}

func (InstrumentController *InstrumentController) UpdateInstrument(c *fiber.Ctx) error {
	id := c.Params("id")
	instrumentParams, err := getParamsInstrument(c)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(403).JSON(errorHandler)
	}

	instrument, err := InstrumentController.instrumentUseCase.Update(id, instrumentParams)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(instrument_serializer.NewInstrumentSerializerFromtInstrument(instrument))
}

func (InstrumentController *InstrumentController) DeleteInstrument(c *fiber.Ctx) error {
	id := c.Params("id")
	instrument, err := InstrumentController.instrumentUseCase.Delete(id)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(instrument_serializer.NewInstrumentSerializerFromtInstrument(instrument))
}
