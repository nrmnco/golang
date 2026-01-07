package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"

	"hw3/internal/db"
	"hw3/internal/handler"
	"hw3/internal/service"
)

func main() {
	ctx := context.Background()
	pool, err := db.NewPostgresDB(ctx)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}
	defer pool.Close()

	myService := service.NewService(pool)
	myHandler := handler.NewHandler(myService)

	e := echo.New()

	e.GET("/student/:id", myHandler.GetStudent)
	e.GET("/all_class_schedule", myHandler.GetAllGroupSchedule)
	e.GET("/schedule/group/:id", myHandler.GetOneGroupSchedule)

	e.Logger.Fatal(e.Start(":8080"))
}
