package controller

import (
	"net/http"

	"go-project/context"
	"go-project/service"
)

// StatusController ...
type StatusController struct {
	service service.StatusService
}

// NewStatusController ...
func NewStatusController(service service.StatusService) *StatusController {
	return &StatusController{
		service: service,
	}
}

// Get ...
func (s *StatusController) Get(c context.Context) error {
	status := s.service.Get()
	return c.JSON(http.StatusOK, status)
}
