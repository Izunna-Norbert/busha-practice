package routers

import (
	"github.com/Izunna-Norbert/busha-practice/controllers"
	"github.com/gin-gonic/gin"
)

func MoviesRoutes(ro *gin.Engine) {

	var controller controllers.MovieController
	v1 := ro.Group("/v1/movies")
	v1.GET("/", controller.GetMovies)
	v1.GET("/:id", controller.GetMovie)
	v1.GET("/:id/comments", controller.GetMovieComments)
	v1.POST("/:id/comments", controller.CreateMovieComment)
}
