package main

import (
	"final-project/config"
	"final-project/controller"
	"final-project/repository"
	"final-project/route"
	"final-project/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	// Inisialisasi Repository
	userRepo := repository.NewUserRepository(config.DB)
	itemRepo := repository.NewItemRepository(config.DB)
	activityRepo := repository.NewActivityRepository(config.DB)

	// Inisialisasi Service
	userService := service.NewUserService(userRepo)
	itemService := service.NewItemService(itemRepo)
	activityService := service.NewActivityService(activityRepo)

	// Inisialisasi Controller
	userController := controller.NewUserController(userService)
	itemController := controller.NewItemController(itemService)
	activityController := controller.NewActivityController(activityService)

	// Registrasi route
	route.SetupRoutes(r, userController, itemController, activityController)

	r.Run(":8080") // Jalankan di localhost:8080
}
