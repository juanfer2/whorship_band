package performances_controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofrs/uuid"

	"github.com/google/uuid"
	performances_domain "github.com/juanfer2/whorship_band/src/performances/domain"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
)

type PeformanceParams struct {
	Date                  string                       `json:"date" validate:"date"`
	PerformanceGroupMates []PerformanceGroupMateParams `json:"performanceGroupMates"`
	PerformanceSongs      []PerformanceSongParams      `json:"performanceSongs"`
}

type PerformanceSongParams struct {
	PerformanceID uuid.UUID `json:"performanceId"`
	SongID        uuid.UUID `json:"songId"`
}

type PerformanceGroupMateParams struct {
	PerformanceID uuid.UUID `json:"performanceId"`
	GroupMateID   uuid.UUID `json:"groupMateId"`
}

func getParamsPerformance(c *fiber.Ctx) (performances_domain.Performance, error) {
	var peformanceParams PeformanceParams

	if err := c.BodyParser(&peformanceParams); err != nil {
		return performances_domain.Performance{}, err
	}

	if err := validateStruct(peformanceParams); err != nil {
		return performances_domain.Performance{}, err
	}

	groupMates, err := GetGroupmatesParamsFromPerformance(peformanceParams.PerformanceGroupMates)
	if err != nil {
		return performances_domain.Performance{}, err
	}

	songs, err := GetSongsParamsFromPerformance(peformanceParams.PerformanceSongs)
	if err != nil {
		return performances_domain.Performance{}, err
	}

	date, err := utilities.StringToDate(peformanceParams.Date)
	if err != nil {
		return performances_domain.Performance{}, err
	}

	utilities.PrintJson(performances_domain.Performance{
		Date:                  date,
		PerformanceGroupMates: groupMates,
		PerformanceSongs:      songs,
	})

	return performances_domain.Performance{
		Date:                  date,
		PerformanceGroupMates: groupMates,
		PerformanceSongs:      songs,
	}, nil
}

func GetGroupmatesParamsFromPerformance(
	groupMateParams []PerformanceGroupMateParams,
) ([]performances_domain.PerformanceGroupMate, error) {
	var groupMateParamList []performances_domain.PerformanceGroupMate

	for _, groupMateParam := range groupMateParams {
		groupMateParamList = append(groupMateParamList, performances_domain.PerformanceGroupMate{
			GroupMateID: groupMateParam.GroupMateID,
		})
	}

	return groupMateParamList, nil
}

func GetSongsParamsFromPerformance(
	groupMateParams []PerformanceSongParams,
) ([]performances_domain.PerformanceSong, error) {
	var performanceSongParamslist []performances_domain.PerformanceSong

	for _, groupMateParam := range groupMateParams {
		performanceSongParamslist = append(
			performanceSongParamslist,
			performances_domain.PerformanceSong{
				SongID: groupMateParam.SongID,
			},
		)
	}

	return performanceSongParamslist, nil
}

func validateStruct[T any](params T) error {
	var errMessage string
	for _, errValidation := range utilities.ValidatorStruct(params) {
		errMessage = fmt.Sprintf("%s %s\n", errMessage, errValidation.Error())
	}

	if len(errMessage) > 0 {
		return fmt.Errorf(errMessage)
	}

	return nil
}
