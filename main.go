package main

import (
	"gee/gee"
	http "net/http"
)

func main() {
	r := gee.New()

	r.Get("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.Run(":3773")
}
