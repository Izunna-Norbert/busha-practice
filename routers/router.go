package routers

import (
	"os"

	"github.com/Izunna-Norbert/busha-practice/routers/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	env := os.Getenv("env")

	if env == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.RequestIDMiddleware())

	//REGISTER ROUTES
	RegisterRoutes(router)

	return router

}
