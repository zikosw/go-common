package echo

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

func ShutdownGracefully(e *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
