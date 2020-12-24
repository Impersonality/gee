package main

import (
	"fmt"
	"gee/gee"
	http "net/http"
)

func main() {
	r := gee.New()

	r.Get("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s", c.Param("name"))
	})
	r.Run(":3773")
	fmt.Println(r)
}
