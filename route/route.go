package route

import (
	"final-project/controller"
	"final-project/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine,
	userController *controller.UserController,
	itemController *controller.ItemController,
	activityController *controller.ActivityController) {

	// AUTH
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	// Middleware JWT dibutuhkan untuk semua route setelah login
	auth := r.Group("/")
	auth.Use(middleware.JWTAuth())

	// ITEM CRUD
	auth.GET("/items", itemController.GetItems)
	auth.POST("/items", middleware.AdminOnly(), itemController.CreateItem)
	auth.PUT("/items/:id", middleware.AdminOnly(), itemController.UpdateItem)
	auth.DELETE("/items/:id", middleware.AdminOnly(), itemController.DeleteItem)

	// ACTIVITY
	auth.GET("/activities", activityController.GetActivities)
	auth.GET("/activities/:id", activityController.GetActivityDetail)
	auth.POST("/activities", activityController.CreateActivity)

}
