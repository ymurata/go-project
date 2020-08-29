package main

import (
	"go-project/infrastructure/database"
	"go-project/wire"
	"os"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func router(e *echo.Echo, db *database.DB) {
	statusRouter(e)
	userRouter(e, db)
}

func statusRouter(e *echo.Echo) {
	controller := wire.NewStatusController()
	e.GET("/", controller.Get)
}

func userRouter(e *echo.Echo, db *database.DB) {
	controller := wire.NewUserController(db)
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

	db, err := database.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	router(e, db)

	e.Logger.Fatal(e.Start(":8000"))
}
