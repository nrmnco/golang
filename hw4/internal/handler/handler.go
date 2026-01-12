package handler

import (
	"hw4/internal/service"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
}

type AttendanceRequest struct {
	ScheduleID string `json:"schedule_id"`
	DateTime   string `json:"date_time"`
	StudentID  string `json:"student_id"`
	Attendance bool   `json:"attendance"`
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

func (h *Handler) RecordAttendance(c echo.Context) error {
	var req AttendanceRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, "Invalid request payload")
	}

	err := h.service.RecordAttendance(req.ScheduleID, req.DateTime, req.StudentID, req.Attendance)
	if err != nil {
		return c.JSON(500, "Failed to record attendance")
	}

	return c.JSON(200, "Attendance recorded successfully")

}

func (h *Handler) GetSubjectAttendance(c echo.Context) error {
	subjectId := c.Param("id")
	attendanceRecords, err := h.service.GetSubjectAttendance(subjectId)
	if err != nil {
		return c.JSON(404, "Attendance records not found for the subject")
	}

	return c.JSON(200, attendanceRecords)
}

func (h *Handler) GetStudentAttendance(c echo.Context) error {
	studentId := c.Param("id")
	attendanceRecords, err := h.service.GetStudentAttendance(studentId)
	if err != nil {
		return c.JSON(404, "Attendance records not found for the student")
	}

	return c.JSON(200, attendanceRecords)
}
