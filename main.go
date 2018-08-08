package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	r := gin.Default()
	configureRouter(r)

	r.Run(":" + port)
}
