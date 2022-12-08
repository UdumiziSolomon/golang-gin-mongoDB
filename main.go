package main

import (
	"net/http"
	"os"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/UdumiziSolomon/Gopher-Backend/configs"
	"github.com/UdumiziSolomon/Gopher-Backend/routes"
	"github.com/gin-contrib/cors"

)

// RETRIEVE LOGS FOR REQUESTS
func retrieveRequestLogs() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func main() {
	router := gin.New();
	router.Use(gin.Recovery(), gin.Logger())

	// CORS IMPLEMENTATION
	CorsConfig := cors.DefaultConfig()
	CorsConfig.AllowOrigins = []string{configs.LoadENV("CLIENTURL")}
	// bind cors to router
	router.Use(cors.New(CorsConfig))

	retrieveRequestLogs()  // retrieve log requests

	//default route
	router.GET("/app", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Default route!!",
		})
	})

	// DB Connection
	configs.ConnectDB()

	// Routes (middleware)
	routes.UserRoutes(router)   

	router.Run(configs.LoadENV("LOCALPORT"))
}
