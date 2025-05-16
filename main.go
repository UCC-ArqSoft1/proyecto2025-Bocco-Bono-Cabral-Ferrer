package main

import (
	"ucc-gorm/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDatabase()
	router := gin.Default
	router.Use(controllers.AllowCORS)
	router.POST("/login", userControllers.Login)
}
