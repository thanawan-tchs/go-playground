package main

import (
	"go-gin-playground/controllers"
	"go-gin-playground/gateway"

	"github.com/gin-gonic/gin"
)

func main() {
	// init gateway
	gateway.InitGateway()

	// init router
	router := gin.Default()

	controllers.Setup(router)

	router.Run("localhost:8080")
}
