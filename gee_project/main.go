package main

import (
	"gee"
)

func main() {
	engine := gee.New()

	engine.GET("/", func(ctx *gee.Context) {
		ctx.HTML(200, "<h1>hello</h1>")
	})

	engine.GET("/hello", func(ctx *gee.Context) {
		name := ctx.Query("name")
		age := ctx.Query("age")
		ctx.String(200, "name=%v, age=%v\n", name, age)
	})
	engine.GET("/json", func(ctx *gee.Context) {
		ctx.JSON(200, gee.H{
			"name": "zs",
			"age":  "100",
		})
	})
	engine.Run(":8081")

}
