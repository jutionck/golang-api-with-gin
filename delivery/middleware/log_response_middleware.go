package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"strconv"
)

func LogResponseMiddleware(log *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedError := c.Errors.Last()
		if detectedError == nil {
			return
		}
		e := detectedError.Error()
		status := strconv.Itoa(c.Writer.Status())
		if status[0] != '2' {
			log.Error().Str("error", e).Str("host", c.Request.RequestURI).Str("method", c.Request.Method).Str("status", status).Msg("Response")
		}

	}
}
