package controllers

import (
	"log"
	"net/http"

	"github.com/Izunna-Norbert/busha-practice/forms"
	"github.com/Izunna-Norbert/busha-practice/models"
	"github.com/Izunna-Norbert/busha-practice/repository"
	"github.com/gin-gonic/gin"
)

type MovieController struct {
	repository.MoviesStore
}

var commentForm = new(forms.CommentForm)

func (ctrl MovieController) GetMovies(ctx *gin.Context) {
	response, err := repository.FetchMovies()

	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot fetch movies"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctrl MovieController) GetMovie(ctx *gin.Context) {
	response, err := repository.FetchMovie(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot fetch movie"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctrl MovieController) GetMovieComments(ctx *gin.Context) {
	response, err := repository.FetchMovieComments(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot fetch movie comments"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctrl MovieController) CreateMovieComment(ctx *gin.Context) {

	var form forms.CreateCommentForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		message := commentForm.Create(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}
	var comment models.Comment

	// save IP address of the user
	comment.ClientIP = ctx.ClientIP()

	comment.IdentifierID = ctx.Param("id")
	comment.Comment = form.Comment
	comment.Name = form.Name

	response, err := repository.CreateMovieComment(&comment)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot create movie comment"})
		return
	}

	//get only the JSON response don't return the Capitalized keys
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}
