package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := uuid.New()
		ctx.Writer.Header().Set("request-id", uuid.String())
		ctx.Next()
	}
}
