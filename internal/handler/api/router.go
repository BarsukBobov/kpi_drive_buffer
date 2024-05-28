package api

import (
	"github.com/gin-gonic/gin"
	"kpi_drive_buffer/internal/handler/api/baseapi"
	"kpi_drive_buffer/internal/handler/api/v1"
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
	v1Group := router.Group("/v1")
	v1Router := v1.NewRouter(h.Router, h.service)
	v1Router.RegisterHandlers(v1Group)
}
