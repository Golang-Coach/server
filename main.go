package main

import (
	"fmt"
	"github.com/Golang-Coach/server/db"
	_ "github.com/Golang-Coach/server/docs"
	"github.com/Golang-Coach/server/route"
	"github.com/gin-gonic/gin"
	"os"
	"runtime"
)

var DB = make(map[string]string)

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

	// set gin in release mode
	gin.SetMode(gin.ReleaseMode)

	ConfigRuntime()

	dataStore := db.Connect()
	defer dataStore.Session.Close()

	r := route.SetupRouter(dataStore)

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
