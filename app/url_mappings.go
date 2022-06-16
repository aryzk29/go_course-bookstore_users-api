package app

import (
	"github.com/aryzk29/go_course-bookstore_users-api/controllers"
)

func MapUrls() {
	router.GET("/ping", controllers.Ping)
	router.GET("/users/:user_id", controllers.GetUSer)
	router.POST("/users", controllers.CreateUser)
}
