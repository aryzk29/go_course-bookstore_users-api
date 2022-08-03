package app

import (
	"github.com/aryzk29/go_course-bookstore_users-api/controllers"
)

func MapUrls() {
	router.GET("/ping", controllers.Ping)
	router.GET("/users/:user_id", controllers.Get)
	router.POST("/users", controllers.Create)
	router.PUT("/users/:user_id", controllers.Update)
	router.PATCH("/users/:user_id", controllers.Update)
	router.DELETE("/users/:user_id", controllers.Delete)
	router.GET("/internal/users/search", controllers.Search)
	router.POST("/users/login", controllers.Login)
}
