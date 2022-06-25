package user

import (
	"fmt"
	"go-gin-playground/core/user"

	"github.com/gin-gonic/gin"
)

type userResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser(ctx *gin.Context) {

	u, err := user.GetUser()
	if err != nil {
		ctx.JSON(400, map[string]interface{}{
			"error": "xxxxxxxx",
		})
	}

	fmt.Println(" controller U :", u)

	ctx.JSON(200, gin.H{
		"data": userResponse{
			Id:    u.Id,
			Name:  u.Name,
			Email: u.Email,
		},
	})
}
