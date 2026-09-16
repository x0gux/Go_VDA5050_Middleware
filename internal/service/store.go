package service

import "example.com/m/v2/internal/domain"

type RobotStore interface {
	SaveState(dto domain.RobotTypeDTO)
	GetState(serial string) (domain.RobotTypeDTO, bool)

	GetAllStates() []domain.RobotTypeDTO
}
