package util

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HandleCors(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{
			"X-Requested-With",
			"Content-Type",
			"Authorization"},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE"},
		AllowOrigins: []string{"*"},
	}))
}
