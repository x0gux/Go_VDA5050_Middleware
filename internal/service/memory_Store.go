package service

import (
	"sync"

	"example.com/m/v2/internal/domain"
)

type MemoryRobotStore struct {
	mu     sync.RWMutex
	states map[string]domain.RobotTypeDTO
}

func NewMemoryRobotStore() RobotStore {
	return &MemoryRobotStore{
		states: make(map[string]domain.RobotTypeDTO),
	}
}

func (s *MemoryRobotStore) SaveState(dto domain.RobotTypeDTO) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.states[dto.SerialNumber] = dto
}

func (s *MemoryRobotStore) GetState(serial string) (domain.RobotTypeDTO, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dto, exists := s.states[serial]
	return dto, exists
}

func (s *MemoryRobotStore) GetAllStates() []domain.RobotTypeDTO {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]domain.RobotTypeDTO, 0, len(s.states))
	for _, dto := range s.states {
		list = append(list, dto)
	}
	return list
}
