package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"vanekt/dental-api/model"
)

type AuthController struct {
	logger    *logging.Logger
	userModel *model.UserModel
}

func NewAuthController(logger *logging.Logger, userModel *model.UserModel) *AuthController {
	return &AuthController{logger, userModel}
}

func (c *AuthController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := c.userModel.GetUserByLogin()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "login",
		})
		return
	}
}

func (c *AuthController) Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "logout",
		})
		return
	}
}
