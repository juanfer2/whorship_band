package group_mate_controller

import (
	group_mate_application "github.com/juanfer2/whorship_band/src/group_mates/application"
	group_mate_repositories "github.com/juanfer2/whorship_band/src/group_mates/infrastructure/repositories"
)

func GroupMateControllerFactory() *GroupMateController {
	groupMateRepository := group_mate_repositories.NewGroupMatePGRepository()
	groupMateUseCase := group_mate_application.NewGroupMateUseCase(groupMateRepository)

	return NewGroupMateController(groupMateUseCase)
}
