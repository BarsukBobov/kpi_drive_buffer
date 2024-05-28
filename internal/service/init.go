package service

import (
	"kpi_drive_buffer/internal/repository/kpidrive"
	"kpi_drive_buffer/pkg/misc"
)

var logger = misc.GetLogger()

type Service struct {
	Buffer *BufferService
}

func NewService(kpiDriveAdapter *kpidrive.Adapter) *Service {
	return &Service{
		Buffer: newBufferService(kpiDriveAdapter),
	}
}
