package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func configureRouter(r *gin.Engine) {
	v1 := r.Group("/v1")

	auth := v1.Group("/auth").Use()
	auth.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login",
		})
	})
	auth.GET("/logout", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "logout",
		})
	})
	auth.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "check",
		})
	})

	patients := v1.Group("/patients")
	patients.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "patients",
		})
	})
	patients.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("patient id: %s", id),
		})
	})
}
