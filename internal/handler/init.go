package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "kpi_drive_buffer/docs"
	"kpi_drive_buffer/internal/handler/api"
	"kpi_drive_buffer/internal/handler/api/baseapi"
	"kpi_drive_buffer/internal/service"
	"kpi_drive_buffer/pkg/misc"
)

func InitRoutes(service *service.Service, maxCons int, production bool) *gin.Engine {
	router := gin.New()

	if production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validDate", misc.ValidDate)
	}

	if !production {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	baseAPIRouter := baseapi.NewRouter(maxCons)

	apiGroup := router.Group("/api/", baseAPIRouter.Middleware.LimitConnections)
	apiRouter := api.NewRouter(baseAPIRouter, service)
	apiRouter.RegisterHandlers(apiGroup)

	return router
}
