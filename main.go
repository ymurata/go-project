package main

import (
	"go-project/wire"

	echo "github.com/labstack/echo/v4"
)

func router(e *echo.Echo) {
	statusRouter(e)
}

func statusRouter(e *echo.Echo) {
	controller := wire.NewStatusController()
	e.GET("/", controller.Get)
}

func main() {
	e := echo.New()
	router(e)
	e.Logger.Fatal(e.Start(":8000"))
}
