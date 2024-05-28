package main

import (
	"kpi_drive_buffer/internal/app"
	"kpi_drive_buffer/internal/models"
	"kpi_drive_buffer/pkg/misc"
)

// @title KPI DRIVE BUFFER API
// @version 1.0
// @description API Server for KPI DRIVE's test task
// @host localhost:3000
// @BasePath /api/v1
func main() {
	envConf, err := models.NewEnvConfig()
	if err != nil {
		panic(err.Error())
		return
	}

	logger := misc.GetLogger()
	logger.SetParams(envConf.LogLevel)

	appObj := app.NewApp(envConf)
	appObj.Init()
}
