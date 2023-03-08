package routers

import (
	"github.com/Izunna-Norbert/busha-practice/controllers"
	"github.com/gin-gonic/gin"
)

func CharactersRoutes(ro *gin.Engine) {

	var controller controllers.CharactersController
	v1 := ro.Group("/v1/characters")
	v1.GET("/", controller.GetCharacters)
	v1.GET("/list", controller.ListCharacters)
}
