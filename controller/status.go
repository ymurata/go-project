package controller

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

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
func (s *StatusController) Get(c echo.Context) error {
	status := s.service.Get()
	return c.JSON(http.StatusOK, status)
}
