package handler

import (
	"github.com/labstack/echo/v4"
	"hw3/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(srvc *service.Service) *Handler {
	return &Handler{service: srvc}
}

func (h *Handler) GetStudent(c echo.Context) error {
	id := c.Param("id")
	student, err := h.service.GetStudent(id)
	if err != nil {
		return c.JSON(404, "Student not found")
	}
	return c.JSON(200, student)
}

func (h *Handler) GetAllGroupSchedule(c echo.Context) error {
	AllGroupsSchedule, err := h.service.GetAllGroupSchedule()
	if err != nil {
		return c.JSON(404, "AllGroupsSchedule not found")
	}
	return c.JSON(200, AllGroupsSchedule)
}

func (h *Handler) GetOneGroupSchedule(c echo.Context) error {
	groupId := c.Param("id")
	groupSchedules, err := h.service.GetOneGroupSchedule(groupId)
	if err != nil {
		return c.JSON(404, "GroupSchedules not found")
	}

	return c.JSON(200, groupSchedules)
}
