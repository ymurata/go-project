package main

import (
	"go-project/wire"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func router(e *echo.Echo) {
	statusRouter(e)
	userRouter(e)
}

func statusRouter(e *echo.Echo) {
	controller := wire.NewStatusController()
	e.GET("/", controller.Get)
}

func userRouter(e *echo.Echo) {
	controller := wire.NewUserController()
	e.GET("/users", controller.List)
	e.POST("/users", controller.Create)
	e.GET("/users/:id", controller.Get)
	e.PUT("/users/:id", controller.Update)
	e.DELETE("/users/:id", controller.Delete)
}

func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	router(e)

	e.Logger.Fatal(e.Start(":8000"))
}
