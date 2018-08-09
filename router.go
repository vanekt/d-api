package main

import (
	"github.com/gin-gonic/gin"
	"vanekt/dental-api/controller"
)

func configureRouter(
	r *gin.Engine,
	authController *controller.AuthController,
	patientController *controller.PatientController,
) {
	v1 := r.Group("/v1")

	// -------------------------------
	auth := v1.Group("/auth").Use()

	loginHandler := authController.Login()
	auth.GET("/login", loginHandler)

	logoutHandler := authController.Logout()
	auth.GET("/logout", logoutHandler)

	// -------------------------------
	patients := v1.Group("/patients")

	patientsHandler := patientController.GetList()
	patients.GET("/", patientsHandler)

	patientHandler := patientController.GetById()
	patients.GET("/:id", patientHandler)
}
