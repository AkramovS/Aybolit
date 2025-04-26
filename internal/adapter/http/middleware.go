package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Передаем управление следующему middleware или обработчику
		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		// Логируем информацию о запросе
		log.Printf("Method: %s, Path: %s, Status: %d, Duration: %v\n", c.Request.Method, c.Request.URL.Path, statusCode, duration)
	}
}
