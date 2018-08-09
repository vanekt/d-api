package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"net/http"
)

type PatientController struct {
	logger *logging.Logger
}

func NewPatientController(logger *logging.Logger) *PatientController {
	return &PatientController{logger}
}

func (c *PatientController) GetList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "patients",
		})
	}
}

func (c *PatientController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("patient id: %s", id),
		})
	}
}
