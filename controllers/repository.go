package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/Golang-Coach/server/interfaces"
)

type RepositoryController struct {
	store interfaces.IDataStore
}


func NewRepositoryController(store interfaces.IDataStore) RepositoryController {
	return RepositoryController{
		store: store,
	}
}

// TODO replace ginContext with IContext
//noinspection ALL
func (this RepositoryController) GetRespositories(context *gin.Context){
	query := context.Query("query")
	page := context.DefaultQuery("page", "1")
	limit := context.DefaultQuery("limit", "20")

	this.store.FindPackageWithinLimit()
}
