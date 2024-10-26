// /examples/blog/main.go

package main

import (
	"log"

	"github.com/editbase/gust"
)

func main() {
	app := gust.New().
		WithPort("3000").
		WithTemplateDir("./templates").
		WithStaticDir("./static")

	app.GET("/", func(c *gust.Context) error {
		return c.Render("index.html", nil)
	})

	log.Fatal(app.Run())
}
