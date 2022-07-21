package app

import (
	"github.com/aryzk29/go_course-bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	MapUrls()

	logger.Info("starting the application...")
	router.Run(":8080")
}
