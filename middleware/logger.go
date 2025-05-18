package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		body := c.Body()
		if len(body) > 0 {
			fmt.Printf("Request Body: %s\n", string(body))
		}

		err := c.Next()

		stop := time.Now()
		latency := stop.Sub(start)

		method := c.Method()
		path := c.Path()
		status := c.Response().StatusCode()
		ip := c.IP()
		userAgent := c.Get("User-Agent")

		fmt.Printf("[%s] %s %s %d %s %s %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			method,
			path,
			status,
			ip,
			userAgent,
			latency,
		)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		return err
	}
}
