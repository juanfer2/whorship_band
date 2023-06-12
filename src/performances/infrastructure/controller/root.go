package performances_controller

import (
	"github.com/gofiber/fiber/v2"

	performances_application "github.com/juanfer2/whorship_band/src/performances/application"
	performances_serializer "github.com/juanfer2/whorship_band/src/performances/infrastructure/serializers"
	middlewares "github.com/juanfer2/whorship_band/src/shared/infrastructure/controllers"
)

type PerformanceController struct {
	performanceUseCase *performances_application.PerformanceUseCase
}

func NewPerformanceController(
	performanceUseCase *performances_application.PerformanceUseCase,
) *PerformanceController {
	return &PerformanceController{performanceUseCase: performanceUseCase}
}

func (PerformanceController *PerformanceController) All(c *fiber.Ctx) error {
	performances, err := PerformanceController.performanceUseCase.GetAll()
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(performances_serializer.NewPerformanceSerializerInstrumentList(performances))
}

func (PerformanceController *PerformanceController) Create(c *fiber.Ctx) error {
	instrumentParams, err := getParamsPerformance(c)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(403).JSON(errorHandler)
	}

	instrument, err := PerformanceController.performanceUseCase.Create(instrumentParams)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(performances_serializer.NewPerformanceSerializerFromItem(instrument))
}

func (PerformanceController *PerformanceController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	instrumentParams, err := getParamsPerformance(c)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(403).JSON(errorHandler)
	}

	instrument, err := PerformanceController.performanceUseCase.Update(id, instrumentParams)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(performances_serializer.NewPerformanceSerializerFromItem(instrument))
}

func (PerformanceController *PerformanceController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	instrument, err := PerformanceController.performanceUseCase.Delete(id)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(performances_serializer.NewPerformanceSerializerFromItem(instrument))
}
