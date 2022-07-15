package app

import "github.com/herrera-ignacio/go-microservices/users_api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
