package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//TODO: handle auth error here
		isAuthSuccess := true

		//----------------------------
		if !isAuthSuccess {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credential")
		}

		customContext := &CustomContext{c, 0}
		return next(customContext)
	}
}
