package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "ngductoan")

	uid := ctx.Query("uid")
	ctx.JSON(http.StatusOK, gin.H{ // map String
		"message": "pong" + name,
		"uid":     uid,
		"users":   []string{"user1", "user2", "ngductoan"},
	})
}
