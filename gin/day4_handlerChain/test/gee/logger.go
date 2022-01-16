package gee

import (
	"fmt"
	"log"
	"time"
)

func Logger() Handler {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func Logger1() Handler {
	return func(c *Context) {
		// Process request
		c.Next()
		// Calculate resolution time
		fmt.Println("middleware 2 has next()")
	}
}

func Logger2() Handler {
	return func(c *Context) {
		fmt.Println("middleware 3 no next()")
	}
}
