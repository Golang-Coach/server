package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
	_ "github.com/Golang-Coach/server/docs"
)

var DB = make(map[string]string)

// @Summary ping
// @Description ping
// @ID get-ping
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /ping [get]
func GetPing(c *gin.Context) {
	c.String(200, "pong")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", GetPing)

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

// @host localhost:8080
func main() {
	r := setupRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Listen and Server in 0.0.0.0:8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	r.Run(":" + port)
}
