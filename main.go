package main

import (
	"os"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-project/context"
	"go-project/extension"
	"go-project/infrastructure/database"
	gm "go-project/middleware"
	"go-project/wire"
)

type handlerFunc func(c *context.Context) error

func cast(next handlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c.(*context.Context))
	}
}

func router(e *echo.Echo, db *database.DB) {
	statusRouter(e)
	userRouter(e, db)
}

func statusRouter(e *echo.Echo) {
	controller := wire.NewStatusController()
	e.GET("/", cast(controller.Get))
}

func userRouter(e *echo.Echo, db *database.DB) {
	controller := wire.NewUserController(db)
	e.GET("/users", cast(controller.List))
	e.POST("/users", cast(controller.Create))
	e.GET("/users/:id", cast(controller.Get))
	e.PUT("/users/:id", cast(controller.Update))
	e.DELETE("/users/:id", cast(controller.Delete))
}

func main() {
	e := echo.New()

	e.HTTPErrorHandler = extension.HTTPErrorHandler
	e.Validator = extension.NewValidator()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(gm.Context())

	db, err := database.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	if _, err := extension.NewHash("salt", 20); err != nil {
		panic(err)
	}

	router(e, db)

	e.Logger.Fatal(e.Start(":8000"))
}
