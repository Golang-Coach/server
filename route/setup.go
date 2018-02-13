package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/Golang-Coach/server/db"
	"github.com/gin-contrib/gzip"
	"github.com/Golang-Coach/server/controllers"
	"github.com/Golang-Coach/server/dal"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)


func SetupRouter(store *db.DataStore) *gin.Engine {
	r := gin.Default()

	//// use compression
	r.Use(gzip.Gzip(gzip.BestSpeed))

	// enable CORS for all domain
	r.Use(cors.Default())

	repositoryController := controllers.NewRepositoryController(dal.NewRepositoryStore(store))

	// Ping test
	r.GET("/repositories/", repositoryController.GetRepositories)
	r.GET("/repositories/:owner/:name", repositoryController.GetRepositoryById)

	if gin.Mode() == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return r
}
