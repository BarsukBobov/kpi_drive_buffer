package baseapi

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Status string `json:"status"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{Error: err.Error()}
}

func AccessDenied(c *gin.Context) {
	Response403(c, errors.New("access denied"))
}

func Response(c *gin.Context, statusCode int, obj interface{}) {
	c.JSON(statusCode, obj)
}

func Response200(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

func NewSuccessResponse(c *gin.Context) {
	Response200(c, SuccessResponse{Status: "success"})
}

func Response401(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, NewErrorResponse(err))
}

func Response403(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, NewErrorResponse(err))
}

func Response404(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, NewErrorResponse(err))
}
