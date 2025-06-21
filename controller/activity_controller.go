package controller

import (
	"net/http"
	"strconv"

	"final-project/entity"
	"final-project/service"

	"github.com/gin-gonic/gin"
)

type ActivityController struct {
	ActivityService service.ActivityService
}

func NewActivityController(s service.ActivityService) *ActivityController {
	return &ActivityController{ActivityService: s}
}

func (c *ActivityController) GetActivities(ctx *gin.Context) {
	activities, err := c.ActivityService.GetAllActivities()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get activities"})
		return
	}
	ctx.JSON(http.StatusOK, activities)
}

func (c *ActivityController) GetActivityDetail(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	activity, err := c.ActivityService.GetActivityByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}
	ctx.JSON(http.StatusOK, activity)
}

func (c *ActivityController) CreateActivity(ctx *gin.Context) {
	var activity entity.Activity
	if err := ctx.ShouldBindJSON(&activity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	println("DEBUG: ItemID =", activity.ItemID)
	println("DEBUG: UserID =", activity.UserID)
	println("DEBUG: Action =", activity.Action)
	println("DEBUG: Date =", activity.Date.String())

	// Validasi wajib
	if activity.ItemID == 0 || activity.UserID == 0 || activity.Action == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ItemID, UserID, dan Action wajib diisi"})
		return
	}

	if activity.Date.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Tanggal aktivitas tidak valid (format: YYYY-MM-DD)"})
		return
	}
	err := c.ActivityService.CreateActivity(&activity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create activity"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Activity created"})
}
