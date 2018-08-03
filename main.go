package main

import (
	"github.com/kataras/iris"
	"os"
)

func main() {
	app := iris.Default()
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})
	port := os.Getenv("PORT")
	app.Run(iris.Addr(":" + port))
}
