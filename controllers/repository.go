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

// @Summary GetRepositories
// @Description GetRepositories
// @ID get-repositories
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 500 {object} string "Internal Server Error"
// @Failure 404 {object} string "Not found"
// @Router /repositories [get]
func (c RepositoryController) GetRepositories(context *gin.Context) {
	query := context.Query("query")

	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}

	limit, err := strconv.Atoi(context.DefaultQuery("limit", "20"))
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}

	skip := (page - 1) * limit
	repositories, err := c.store.FindPackageWithinLimit(query, skip, limit)
	if err != nil {
		if err.Error() == "not found" {
			context.JSON(http.StatusNotFound, err.Error())
			return
		}
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, repositories)
}

// @Summary GetRepository by ID
// @Description GetRepository by ID
// @ID get-repository-by-id
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Repository Id"
// @Success 200 {string} string	"ok"
// @Failure 500 {object} string "Internal Server Error"
// @Failure 404 {object} string "Not found"
// @Router /repositories/{id} [get]
func (c RepositoryController) GetRepositoryById(context *gin.Context) {
	id := context.Param("id")

	repository, err := c.store.FindId(id)

	if err != nil {
		if err.Error() == "not found" {
			context.JSON(http.StatusNotFound, err.Error())
			return
		}
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, repository)
}
