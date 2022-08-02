package main

import (
	"gee"
)

func main() {
	engine := gee.New()
	rg := engine.Group("/v1")
	rg.Use(gee.Logger()) // 添加全局logger中间件
	{
		rg.GET("/hello", func(c *gee.Context) {
			c.String(200, "v1 hello world")
		})
		rg.GET("/hello/:name", func(c *gee.Context) {
			name := c.Param("name")
			c.String(200, "hello name:%v\n", name)
		})
	}

	rg2 := engine.Group("/v2")
	{
		rg2.GET("/hello", func(c *gee.Context) {
			c.String(200, "v2 hello world")
		})
	}

	engine.Run(":8081")
}
