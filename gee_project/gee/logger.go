package gee

import (
	"log"
	"time"
)

func Logger() HandleFunc {
	return func(c *Context) {
		start := time.Now()

		c.Next() // 处理请求

		log.Printf("[%d] %s in %v", c.statuscode, c.Req.RequestURI, time.Since(start))
	}
}
