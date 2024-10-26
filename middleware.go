// middleware.go
package gust

import (
	"fmt"
	"log"
	"time"
)

type MiddlewareFunc func(next HandlerFunc) HandlerFunc

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
