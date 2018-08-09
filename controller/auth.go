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
		user, err := c.userModel.GetUserByLogin("admin")
		if err != nil {
			c.logger.Error("AuthController Login: ", err.Error())
			ctx.JSON(500, gin.H{
				"error": "Internal Server Error",
			})
			return
		}

		c.logger.Debug("User: ", user)

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
