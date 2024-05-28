package buffer

import (
	"github.com/gin-gonic/gin"
	"kpi_drive_buffer/internal/handler/api/baseapi"
	"kpi_drive_buffer/internal/service"
)

type router struct {
	*baseapi.Router
	service *service.BufferService
}

func NewRouter(
	baseAPIRouter *baseapi.Router,
	service *service.BufferService,
) baseapi.ApiRouter {
	return &router{
		Router:  baseAPIRouter,
		service: service,
	}
}

func (h *router) RegisterHandlers(router *gin.RouterGroup) {
	router.POST("/save_fact", h.saveFact)
}
