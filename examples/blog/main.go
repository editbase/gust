// /examples/blog/main.go

package main

import (
	"log"

	"github.com/editbase/stardust"
)

func main() {
	app := stardust.New().
		WithPort("3000").
		WithTemplateDir("./templates").
		WithStaticDir("./static")

	app.GET("/", func(c *stardust.Context) error {
		return c.Render("index.html", nil)
	})

	log.Fatal(app.Run())
}
