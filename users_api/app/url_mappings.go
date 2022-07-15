package app

import (
	"github.com/herrera-ignacio/go-microservices/users_api/controllers/ping"
	"github.com/herrera-ignacio/go-microservices/users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/find", users.FindUser)
	router.POST("/users", users.CreateUser)
}
