// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package wire

import (
	"github.com/google/wire"

	"go-project/controller"
	"go-project/domain/repository"
	"go-project/infrastructure/database"
	"go-project/service"
)

// NewStatusController ...
func NewStatusController() *controller.StatusController {
	wire.Build(
		repository.NewStatusRepositoryImpl,
		service.NewStatusServiceImpl,
		controller.NewStatusController,
		wire.Bind(new(service.StatusService), new(*service.StatusServiceImpl)),
		wire.Bind(new(repository.StatusRepository), new(*repository.StatusRepositoryImpl)),
	)
	return &controller.StatusController{}
}

// NewUserController ...
func NewUserController(db *database.DB) *controller.UserController {
	wire.Build(
		repository.NewUserRepositoryImpl,
		service.NewUserServiceImpl,
		controller.NewUserController,
		wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
		wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	)
	return &controller.UserController{}
}
