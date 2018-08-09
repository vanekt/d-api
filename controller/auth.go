package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"net/http"
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
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "Not found",
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
			return
		}

		c.logger.Debug("User: ", user)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "login",
		})
		return
	}
}

func (c *AuthController) Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "logout",
		})
		return
	}
}
