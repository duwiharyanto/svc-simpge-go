package router

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func SetResponseTimeout(ctx context.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Request().Header.Set("X-Member", "ahmad")
			dur, err := time.ParseDuration(os.Getenv("RESPONSE_TIMEOUT_MS") + "ms")
			if err != nil {
				dur = time.Second * 40 // default
			}
			ctx, cancel := context.WithTimeout(ctx, dur)
			defer cancel()
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

func HandleError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if he, ok := err.(*echo.HTTPError); ok {
			fmt.Printf("he: %+v\n", he)
		}
		return next(c)
		// return echo.NewHTTPError(404, "not found")
	}
}
