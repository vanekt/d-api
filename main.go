package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/op/go-logging"
	"os"
	"vanekt/dental-api/controller"
	"vanekt/dental-api/model"
)

func main() {
	port := os.Getenv("PORT")
	logger := NewLogger("dental-api", logging.DEBUG)
	db := sqlx.MustConnect("mysql", os.Getenv("SQL_DB_DSN"))

	// init models
	userModel := model.NewUserModel(db, logger)

	// init controllers
	authController := controller.NewAuthController(logger, userModel)
	patientController := controller.NewPatientController(logger)

	// init gin
	r := gin.Default()
	configureRouter(r, authController, patientController)

	r.Run(":" + port)
}
