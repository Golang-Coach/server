package main

import (
	"fmt"
	"github.com/Golang-Coach/server/controllers"
	"github.com/Golang-Coach/server/dal"
	"github.com/Golang-Coach/server/db"
	_ "github.com/Golang-Coach/server/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
	"runtime"
)

var DB = make(map[string]string)

func setupRouter(store *db.DataStore) *gin.Engine {
	r := gin.Default()

	//// use compression
	r.Use(gzip.Gzip(gzip.BestSpeed))

	// enable CORS for all domain
	r.Use(cors.Default())

	repositoryController := controllers.NewRepositoryController(dal.NewRepositoryStore(store))

	// Ping test
	r.GET("/repositories/", repositoryController.GetRepositories)
	r.GET("/repositories/:id", repositoryController.GetRepositoryById)

	return r
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3005
func main() {
	fmt.Println("server has been started !!!!!")

	ConfigRuntime()

	dataStore := db.Connect()
	defer dataStore.Session.Close()

	r := setupRouter(dataStore)

	if gin.Mode() == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Listen and Server in 0.0.0.0:8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "3005"
	}
	r.Run(":" + port)
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}
