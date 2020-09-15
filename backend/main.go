package main

import (
	"os"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-project/extension"
	"go-project/extension/types"
	"go-project/infrastructure/database"
	gm "go-project/interface/middleware"
	"go-project/wire"
)

func router(r *echo.Group, db database.Handler) {
	statusRouter(r)
	userRouter(r, db)
}

func statusRouter(r *echo.Group) {
	controller := wire.NewStatusController()
	r.GET("", gm.Cast(controller.Get))
}

func userRouter(r *echo.Group, db database.Handler) {
	controller := wire.NewUserController(db)
	r.GET("/users", gm.Cast(controller.List))
	r.POST("/users", gm.Cast(controller.Create))
	r.GET("/users/:id", gm.Cast(controller.Get))
	r.PUT("/users/:id", gm.Cast(controller.Update))
	r.DELETE("/users/:id", gm.Cast(controller.Delete))
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

	if _, err := types.NewHash(os.Getenv("HASH_SALT"), 20); err != nil {
		panic(err)
	}

	g := e.Group("/api")
	router(g, db)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
