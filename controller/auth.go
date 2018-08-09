package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

type AuthController struct {
	logger *logging.Logger
}

func NewAuthController(logger *logging.Logger) *AuthController {
	return &AuthController{logger}
}

func (c *AuthController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "login",
		})
	}
}

func (c *AuthController) Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "logout",
		})
	}
}
