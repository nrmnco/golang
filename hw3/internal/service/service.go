package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"

	"hw3/internal/model"
)

type Service struct {
	pool *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

func (s *Service) GetStudent(id string) (*model.Students, error) {
	row := s.pool.QueryRow(context.Background(), "SELECT * FROM students WHERE id = $1", id)

	var student model.Students
	if err := row.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Sex, &student.BirthDate, &student.GroupID); err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *Service) GetAllGroupSchedule() (*[]model.Schedule, error) {
	rows, err := s.pool.Query(context.Background(), "SELECT * FROM schedule")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []model.Schedule
	for rows.Next(){
		var schedule model.Schedule
		if err := rows.Scan(&schedule.ID, &schedule.GroupID, &schedule.FacultyID, &schedule.DayOfWeek, &schedule.StartTime, &schedule.EndTime, &schedule.Subject); err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	
	return &schedules, nil
}

func (s *Service) GetOneGroupSchedule(groupId string) (*[]model.Schedule, error) {
	rows, err := s.pool.Query(context.Background(), "SELECT * FROM schedule WHERE group_id = $1", groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groupSchedules []model.Schedule
	for rows.Next() {
		var schedule model.Schedule

		if err := rows.Scan(&schedule.ID, &schedule.GroupID, &schedule.FacultyID, &schedule.DayOfWeek, &schedule.StartTime, &schedule.EndTime, &schedule.Subject); err != nil {
			return nil, err
		}
		
		groupSchedules = append(groupSchedules, schedule)
	}

	if len(groupSchedules) == 0 {
		return nil, errors.New("no schedule found")
	}

	return &groupSchedules, nil
}

