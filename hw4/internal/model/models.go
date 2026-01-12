package model

import "time"

type Faculties struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Groups struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	FacultyID int    `db:"faculty_id"`
}

type Students struct {
	ID        int       `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Sex       string    `db:"sex"`
	BirthDate time.Time `db:"birth_date"`
	GroupID   int       `db:"group_id"`
}

type Schedule struct {
	ID        int    `db:"id"`
	Subject   string `db:"subject"`
	DayOfWeek string `db:"day_of_week"`
	StartTime string `db:"start_time"`
	EndTime   string `db:"end_time"`
	GroupID   int    `db:"group_id"`
	FacultyID int    `db:"faculty_id"`
}

type Attendance struct {
	ID         int       `db:"id"`
	ScheduleID int       `db:"schedule_id"`
	StudentID  int       `db:"student_id"`
	DateTime   time.Time `db:"date_time"`
	Attendance bool      `db:"attendance"`
}
