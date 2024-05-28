package models

import "kpi_drive_buffer/pkg/misc"

type EnvConfig struct {
	Production    bool   `env:"PRODUCTION"`
	LogLevel      string `env:"LOG_LEVEL"`
	KpiDriveUrl   string `env:"KPI_DRIVE_URL"`
	KpiDriveToken string `env:"KPI_DRIVE_TOKEN"`
}

func NewEnvConfig() (*EnvConfig, error) {
	var conf EnvConfig
	err := misc.ParseEnv(&conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
