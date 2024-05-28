package service

import (
	"kpi_drive_buffer/internal/repository/kpidrive"
	"sync"
)

type BufferService struct {
	kdApi *kpidrive.Adapter
	mux   sync.Mutex
}

func newBufferService(kpiDriveAdapter *kpidrive.Adapter) *BufferService {
	return &BufferService{kdApi: kpiDriveAdapter}
}

func (s *BufferService) SaveFact(saveFactForm *kpidrive.SaveFactForm) (*kpidrive.SaveFactResponseWithStatus, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	resp, err := s.kdApi.SaveFact(saveFactForm)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}
