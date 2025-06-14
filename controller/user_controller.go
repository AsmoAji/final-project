package controller

import (
	"final-project/entity"
	"final-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (c *UserController) Register(ctx *gin.Context) {
	var input entity.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	token, err := c.UserService.Register(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registrasi berhasil", "token": token})
}

// Struct khusus untuk login input
type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *UserController) Login(ctx *gin.Context) {
	var input LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	token, err := c.UserService.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login berhasil", "token": token})
}
