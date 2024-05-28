package baseapi

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/semaphore"
	"net/http"
	"time"
)

type Middleware struct {
	sem *semaphore.Weighted
}

func NewMiddleware(maxCons int) *Middleware {
	return &Middleware{
		sem: semaphore.NewWeighted(int64(maxCons)),
	}
}

func (m *Middleware) LimitConnections(c *gin.Context) {
	// Создаем контекст с таймаутом ожидания
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	// Попробуем занять семафор
	if err := m.sem.Acquire(ctx, 1); err != nil {
		// Если не удалось занять семафор в течение времени таймаута, возвращаем ошибку
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error": "Too many requests, please try again later.",
		})
		c.Abort()
		return
	}
	c.Next()
	m.sem.Release(1)
}

func (m *Middleware) SessionRequired(c *gin.Context) {
}

func (m *Middleware) AdminRequired(c *gin.Context) {
}

func (m *Middleware) AuthRequired(c *gin.Context) {
}
