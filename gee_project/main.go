package main

import (
	"gee"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(c *gee.Context) {
		c.Json(200, gee.H{
			"name": "zs",
			"age":  12,
		})
	})
	engine.GET("/hello", func(c *gee.Context) {
		age := c.Query("age")
		c.String(200, "hello %s? Query age: %s", "zs", age)
	})
	engine.GET("/hello/:name", func(c *gee.Context) {
		s := c.Param("name")

		c.String(200, "hello wolrd-%v", s)
	})

	engine.GET("/assets/*filepath", func(c *gee.Context) {
		s := c.Param("filepath")
		c.String(200, "assets/----%v\n", s)
	})

	rg := engine.Group("/v1")
	{
		rg.GET("/hello", func(c *gee.Context) {
			c.String(200, "new version hello world")
		})
		rg.GET("/hello/:name", func(c *gee.Context) {
			name := c.Param("name")
			c.String(200, "hello name:%v\n", name)
		})
	}

	engine.Run(":8081")
}
