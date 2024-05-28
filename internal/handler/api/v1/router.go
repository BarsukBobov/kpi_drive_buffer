package v1

import (
	"github.com/gin-gonic/gin"
	"kpi_drive_buffer/internal/handler/api/baseapi"
	"kpi_drive_buffer/internal/handler/api/v1/buffer"
	"kpi_drive_buffer/internal/service"
)

type router struct {
	*baseapi.Router
	service *service.Service
}

func NewRouter(
	baseAPIRouter *baseapi.Router,
	service *service.Service,
) baseapi.ApiRouter {
	return &router{
		Router:  baseAPIRouter,
		service: service,
	}
}

func (h *router) RegisterHandlers(router *gin.RouterGroup) {
	bufferGroup := router.Group("/buffer")
	bufferRouter := buffer.NewRouter(h.Router, h.service.Buffer)
	bufferRouter.RegisterHandlers(bufferGroup)
}
