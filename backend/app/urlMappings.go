package app

import (
	userController "gym-api/backend/controllers/user"
)

func mapUrls() {

	// User Mapping
	router.POST("/users/login", userController.Login)
	//router.POST("/users/register", userController.Register)

	// activity Mapping

	// enrollment Mapping

}
