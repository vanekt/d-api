package main

import (
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"os"
	"vanekt/dental-api/controller"
)

func main() {
	port := os.Getenv("PORT")
	logger := NewLogger("dental-api", logging.DEBUG)

	authController := controller.NewAuthController(logger)
	patientController := controller.NewPatientController(logger)

	r := gin.Default()
	configureRouter(r, authController, patientController)

	r.Run(":" + port)
}
