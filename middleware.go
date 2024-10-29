// /middleware.go
// MiddlewareFunc defines the middleware function signature

package stardust

import (
	"fmt"
	"log"
	"time"
)

// Takes a handler function and returns a wrapped handler function
type MiddlewareFunc func(next HandlerFunc) HandlerFunc

// Logger returns a middleware that logs request details and timing
func Logger() MiddlewareFunc {
	return func(next HandlerFunc) HandlerFunc {
		return func(c *Context) error {
			start := time.Now()
			err := next(c)
			log.Printf("[%s] %s %s %v", c.Request.Method, c.Request.URL.Path, time.Since(start), err)
			return err
		}
	}
}

// Recover returns a middleware that recovers from panics in handlers
func Recover() MiddlewareFunc {
	return func(next HandlerFunc) HandlerFunc {
		return func(c *Context) (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("panic: %v", r)
				}
			}()
			return next(c)
		}
	}
}
