// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"go-project/controller"
	"go-project/domain/repository"
	"go-project/infrastructure/database"
	"go-project/service"
)

// Injectors from wire.go:

func NewStatusController() *controller.StatusController {
	statusRepositoryImpl := repository.NewStatusRepositoryImpl()
	statusServiceImpl := service.NewStatusServiceImpl(statusRepositoryImpl)
	statusController := controller.NewStatusController(statusServiceImpl)
	return statusController
}

func NewUserController(db *database.DB) *controller.UserController {
	userRepositoryImpl := repository.NewUserRepositoryImpl(db)
	userServiceImpl := service.NewUserServiceImpl(userRepositoryImpl)
	userController := controller.NewUserController(userServiceImpl)
	return userController
}