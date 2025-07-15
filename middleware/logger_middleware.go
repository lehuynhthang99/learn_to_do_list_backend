package middleware

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		method := c.Request().Method
		path := c.Request().URL.Path

		err := next(c)

		log.Printf("%s %s | %v", method, path, time.Since(start))
		return err
	}
}
