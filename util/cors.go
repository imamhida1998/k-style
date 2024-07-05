package util

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
