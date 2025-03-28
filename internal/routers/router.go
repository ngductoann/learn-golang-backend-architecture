package routers

import (
	c "learn-golang-backend-architecture/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Groping routes
	v1 := r.Group("/v1")
	{
		// v1.GET("/ping", c.Pong)                     // import basic version
		v1.GET("/ping", c.NewPongController().Pong) // import advanced version

		// v1.GET("/user/1", controller.GetUserByID) // import basic version
		v1.GET("/user/1", c.NewUserController().GetUserByID) // import advanced version

		// v1.PUT("/ping", Pong)
		// v1.PATCH("/ping", Pong)
		// v1.DELETE("/ping", Pong)
		// v1.HEAD("/ping", Pong)
		// v1.OPTIONS("/ping", Pong)
	}

	return r
}
