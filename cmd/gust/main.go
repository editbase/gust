// /cmd/gust/main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/editbase/gust/util/engine"
	"github.com/editbase/gust/util/handler"
)

func main() {
	port := flag.String("port", "3000", "port to serve on")
	flag.Parse()

	eng := engine.NewStandardEngine()
	h := handler.New(eng)

	fmt.Printf("Server starting on http://localhost:%s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, h))
}
