package middleware

import (
	"Rebuild_demo/utils/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO log
		logger.Log.Info("request", zap.Time("time", time.Now()),
			zap.String("url", c.Request.URL.Path),
		)
		c.Next()
		// TODO log
	}
}
