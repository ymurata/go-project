package service

import (
	"go-project/domain/model"
	"go-project/domain/repository"
)

type (
	// StatusService ...
	StatusService interface {
		Get() *model.Status
	}
	// StatusServiceImpl ...
	StatusServiceImpl struct {
		repo repository.StatusRepository
	}
)

// NewStatusServiceImpl ...
func NewStatusServiceImpl(repo repository.StatusRepository) *StatusServiceImpl {
	return &StatusServiceImpl{
		repo: repo,
	}
}

// Get ...
func (s *StatusServiceImpl) Get() *model.Status {
	return s.repo.Get()
}

var _ StatusService = (*StatusServiceImpl)(nil)
