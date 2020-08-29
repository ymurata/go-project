package repository

import "go-project/domain/model"

type (
	// StatusRepository ...
	StatusRepository interface {
		Get() *model.Status
	}
	// StatusRepositoryImpl ...
	StatusRepositoryImpl struct{}
)

// NewStatusRepositoryImpl ...
func NewStatusRepositoryImpl() *StatusRepositoryImpl {
	return &StatusRepositoryImpl{}
}

// Get ...
func (s *StatusRepositoryImpl) Get() *model.Status {
	return &model.Status{
		Status: "ok",
	}
}

var _ StatusRepository = (*StatusRepositoryImpl)(nil)
