package controllers

import (
	"github.com/Golang-Coach/server/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RepositoryController struct {
	store interfaces.IRepositoryStore
}

func NewRepositoryController(store interfaces.IRepositoryStore) RepositoryController {
	return RepositoryController{
		store,
	}
}

// TODO replace ginContext with IContext
//noinspection ALL
// @Summary GetRepositories
// @Description GetRepositories
// @ID get-repositories
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /repositories [get]
func (c RepositoryController) GetRepositories(context *gin.Context) {
	query := context.Query("query")
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))    // TODO
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "20")) // TODO
	skip := (page - 1) * limit
	repositories, _ := c.store.FindPackageWithinLimit(query, skip , limit)

	context.JSON(http.StatusOK, repositories)
}
