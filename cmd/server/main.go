package main

import (
	"learn-golang-backend-architecture/internal/routers"
)

func main() {
	// fmt.Println("Startin")
	// r := gin.Default()

	// basic route
	// r.GET("/ping", Pong)
	// r.PUT("/ping", Pong)
	// r.PATCH("/ping", Pong)
	// r.DELETE("/ping", Pong)
	// r.HEAD("/ping", Pong)
	// r.OPTIONS("/ping", Pong)

	// Groping routes
	// v1 := r.Group("/v1")
	// {
	// 	r.GET("/ping", Pong)
	// 	r.PUT("/ping", Pong)
	// 	r.PATCH("/ping", Pong)
	// 	r.DELETE("/ping", Pong)
	// 	r.HEAD("/ping", Pong)
	// 	r.OPTIONS("/ping", Pong)
	// }

	// r.Run() // listen server on 0.0.0.0:8080

	r := routers.NewRouter()

	r.Run(":8002")
}
