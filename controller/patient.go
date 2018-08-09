package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

type PatientController struct {
	logger *logging.Logger
}

func NewPatientController(logger *logging.Logger) *PatientController {
	return &PatientController{logger}
}

func (c *PatientController) GetList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "patients",
		})
	}
}

func (c *PatientController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{
			"message": fmt.Sprintf("patient id: %s", id),
		})
	}
}
