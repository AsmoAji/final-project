package controller

import (
	"net/http"
	"strconv"

	"final-project/entity"
	"final-project/service"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	ItemService service.ItemService
}

func NewItemController(itemService service.ItemService) *ItemController {
	return &ItemController{ItemService: itemService}
}

func (c *ItemController) GetItems(ctx *gin.Context) {
	items, err := c.ItemService.GetAllItems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get items"})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func (c *ItemController) CreateItem(ctx *gin.Context) {
	var item entity.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	// Validasi dasar
	if item.Name == "" || item.SerialNumber == "" || item.Category == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama, serial number, dan kategori wajib diisi"})
		return
	}

	err := c.ItemService.CreateItem(&item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Item created"})
}

func (c *ItemController) UpdateItem(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	var item entity.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	item.ID = uint(id)
	err := c.ItemService.UpdateItem(&item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Item updated"})
}

func (c *ItemController) DeleteItem(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	err := c.ItemService.DeleteItem(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}
