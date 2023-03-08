package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(ro *gin.Engine) {
	ro.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not Found"})
	})

	ro.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	MoviesRoutes(ro)
	CharactersRoutes(ro)
}
