package repository

import "example.com/m/v2/internal/domain/robot"

type RobotStore interface {
	SaveState(state robot.State)
	GetState(serial string) (robot.State, bool)
	GetAllStates() []robot.State
}
