package performances_controller

import (
	performances_application "github.com/juanfer2/whorship_band/src/performances/application"
	performances_repositories "github.com/juanfer2/whorship_band/src/performances/infrastructure/repositories"
)

func PerformanceControllerFactory() *PerformanceController {
	performanceRepository := performances_repositories.NewPerformancePGRepository()
	performanceUseCase := performances_application.NewPerformanceUseCase(performanceRepository)

	return NewPerformanceController(performanceUseCase)
}
