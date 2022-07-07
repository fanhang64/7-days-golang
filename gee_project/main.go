package main

import (
	"fmt"
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

	engine.GET("/hello/:name", func(ctx *gee.Context) {
		name := ctx.Param("name")
		fmt.Println(name, "--------")
		ctx.String(200, "ok")
	})

	engine.GET("/hello/:name/doc", func(ctx *gee.Context) {
		name := ctx.Param("name")
		fmt.Println(name, "--------")
		ctx.String(200, "ok")
	})
	engine.GET("/world/*path", func(ctx *gee.Context) {
		path := ctx.Param("path")
		fmt.Println(path, "----===")
		ctx.String(200, "path: %v\n", path)
	})

	engine.GET("/json", func(ctx *gee.Context) {
		ctx.JSON(200, gee.H{
			"name": "zs",
			"age":  "100",
		})
	})
	engine.Run(":8081")

}
