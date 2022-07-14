package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func LogRequestMiddleware(log *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info().Str("host", c.Request.RequestURI).Str("method", c.Request.Method).Msg("Request")
		c.Next()
	}
}
