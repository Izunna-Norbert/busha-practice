package controllers

import (
	"net/http"

	"github.com/Izunna-Norbert/busha-practice/repository"
	"github.com/Izunna-Norbert/busha-practice/utils"
	"github.com/gin-gonic/gin"
)

type CharactersController struct {
	repository.CharactersStore
}

func (ctrl CharactersController) GetCharacters(ctx *gin.Context) {

	//get page and limit from query params
	page := ctx.DefaultQuery("page", "1")

	response, err := repository.FetchCharacters(page)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot fetch characters"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctrl CharactersController) ListCharacters(ctx *gin.Context) {
	p := utils.FormulatePagination(ctx)

	// filter := ctx.Query("filter")
	// filterby := ctx.Query("filterBy")

	// if filter != "" && filterby != "" {
	// 	response, err := repository.ListFilteredCharacters(p, filter, filterby)

	// 	if err != nil {
	// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot fetch characters"})
	// 		return
	// 	}

	// 	ctx.JSON(http.StatusOK, gin.H{"data": response})
	// 	return
	// }

	// get characters from repository
	characters, err := repository.ListCharacters(p)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot fetch characters"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": characters})

}
