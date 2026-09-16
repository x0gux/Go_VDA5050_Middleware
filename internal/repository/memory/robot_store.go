package memory

import (
	"sync"

	"example.com/m/v2/internal/domain/robot"
	"example.com/m/v2/internal/repository"
)

type RobotStore struct {
	mu     sync.RWMutex
	states map[string]robot.State
}

func NewRobotStore() repository.RobotStore {
	return &RobotStore{
		states: make(map[string]robot.State),
	}
}

func (s *RobotStore) SaveState(state robot.State) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.states[state.SerialNumber] = state
}

func (s *RobotStore) GetState(serial string) (robot.State, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	state, exists := s.states[serial]

	return state, exists
}

func (s *RobotStore) GetAllStates() []robot.State {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]robot.State, 0, len(s.states))

	for _, state := range s.states {
		list = append(list, state)
	}

	return list
}
