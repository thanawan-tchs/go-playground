package user

import (
	"github.com/gin-gonic/gin"
)

func PostUser(ctx *gin.Context) {

	ctx.JSON(200, map[string]interface{}{
		"data": "post user",
	})
}
