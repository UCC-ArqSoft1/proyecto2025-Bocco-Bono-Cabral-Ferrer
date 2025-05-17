package main

import (
	"gym-api/backend/db"
)

func main() {
	db.InitDatabase()
	// router := gin.Default
	// router.Use(controllers.AllowCORS)
	// router.POST("/login", userControllers.Login)
}
