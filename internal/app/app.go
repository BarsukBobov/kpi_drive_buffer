package app

import (
	"context"
	"errors"
	"fmt"
	"kpi_drive_buffer/internal/handler"
	"kpi_drive_buffer/internal/models"
	"kpi_drive_buffer/internal/repository/kpidrive"
	"kpi_drive_buffer/internal/service"
	"kpi_drive_buffer/pkg/misc"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var logger = misc.GetLogger()

const maxConcurrentConnections = 1000

type App struct {
	envConf *models.EnvConfig
	srv     *misc.Server
}

func NewApp(envConf *models.EnvConfig) *App {
	return &App{envConf: envConf}
}

func (a *App) Init() {
	err := a.start()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("Kpi Drive Buffer API started successfully")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	a.shutdown()

	logger.Info("Kpi Drive Buffer API shutting down successfully")
}

func (a *App) start() error {
	var err error
	var errMessage string

	kpiDriveAdapter, err := kpidrive.NewAdapter(a.envConf.KpiDriveUrl, a.envConf.KpiDriveToken)
	if err != nil {
		errMessage = fmt.Sprintf("failed to connect Kpi Drive API: %s", err.Error())
		return errors.New(errMessage)
	}

	logger.Info("Adapter created successfully")

	newService := service.NewService(kpiDriveAdapter)

	a.srv = new(misc.Server)

	go func() {
		if err = a.srv.Run(handler.InitRoutes(newService, maxConcurrentConnections, a.envConf.Production)); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Errorf("error occured while running http server: %s", err.Error())
			}
		}
	}()

	return nil
}

// shutdown function for graceful close. The function will not run on develop mode
func (a *App) shutdown() {
	if err := a.srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error occured on server shutting down: %s", err.Error())
	} else {
		logger.Info("HTTP server closed")
	}
}
