package controllers

import (
	"go-gin-playground/controllers/user"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	// init router
	router.GET("/hello", user.GetUser)
}
