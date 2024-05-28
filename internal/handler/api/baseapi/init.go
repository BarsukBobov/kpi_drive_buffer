package baseapi

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Middleware *Middleware
}

func NewRouter(maxCons int) *Router {
	return &Router{
		Middleware: NewMiddleware(maxCons),
	}
}

type ApiRouter interface {
	RegisterHandlers(router *gin.RouterGroup)
}
