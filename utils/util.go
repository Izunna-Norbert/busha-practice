package utils

import (
	"strconv"

	"github.com/Izunna-Norbert/busha-practice/models"
	"github.com/gin-gonic/gin"
)


func FormulatePagination(ctx *gin.Context) models.Pagination {
	p := &models.Pagination{}
	p.Page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))
	p.Limit, _ = strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	sort := ctx.DefaultQuery("sort", "desc")
	sortBy := ctx.DefaultQuery("sortBy", "created")

	filter := ctx.Query("filter")
	filterby := ctx.Query("filterBy")

	if filter != "" && filterby != "" {
		p.Filter = map[string]string{filterby: filter}
	}

	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}
	p.Sort = sortBy + " " + sort

	return *p
}