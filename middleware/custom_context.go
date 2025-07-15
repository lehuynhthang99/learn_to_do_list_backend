package middleware

import (
	"github.com/labstack/echo/v4"
)

// CustomContext extends echo.Context to add your own fields
type CustomContext struct {
	echo.Context
	UserID uint // Optional: attach user ID from token later
}

// Helper to cast echo.Context to your CustomContext
func GetCustomContext(c echo.Context) (value *CustomContext, isSuccess bool) {
	value, isSuccess = c.(*CustomContext)
	return
}
