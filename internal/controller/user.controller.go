package controller

import (
	"learn-golang-backend-architecture/internal/service"
	"learn-golang-backend-architecture/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// basic version
func GetUserByID(ctx *gin.Context) {
	// do something
	name := ctx.DefaultQuery("name", "ngductoan")
	uid := ctx.Query("uid")
	ctx.JSON(http.StatusOK, gin.H{ // map String
		"message": "pong" + name,
		"uid":     uid,
		"users":   []string{"user1", "user2", "ngductoan"},
	})
}

// advanced version use struct
type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> service -> repo -> models -> dbs
func (uc *UserController) GetUserByID(ctx *gin.Context) {
	// do something

	// ctx.JSON(http.StatusOK, gin.H{ // map String
	// 	"message": uc.userService.GetInfoUser(),
	// 	"users":   []string{"user1", "user2", "ngductoan"},
	// })

	// 	Code:    20001,
	// ctx.JSON(http.StatusOK, response.ResponseData{
	// 	Message: "Success",
	// 	Data:    []string{"user1", "user2", "ngductoan"},
	// })

	response.SuccessResponse(ctx, 20001, []string{"user1", "user2", "ngductoan"})
}
