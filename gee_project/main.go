package main

import (
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()

	engine.GET("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	})

	engine.Run(":8081")

}
