# Golang Learning & Homework Repository

This repository contains a collection of Go projects developed as part of a learning journey or homework assignments.

## Project Structure

The repository is organized into several modules, each representing a stage of development or a specific homework assignment:

- **[hw1](file:///home/nrmn/Desktop/golang/hw1)**: A basic HTTP server implementation using the standard library, demonstrating JSON response handling.
- **[hw2](file:///home/nrmn/Desktop/golang/hw2)**: Introduction to the [Echo](https://echo.labstack.com/) web framework, featuring basic routing and query parameter handling. Includes database preparation scripts.
- **[hw3](file:///home/nrmn/Desktop/golang/hw3)**: A more advanced backend for student management and class schedules. Integrates with PostgreSQL using `pgx/v5`.
- **[hw4](file:///home/nrmn/Desktop/golang/hw4)**: Extended functionality building upon `hw3`, adding features for recording and retrieving student attendance.
## Technology Stack

- **Backend**:
  - [Go](https://go.dev/) (v1.25.5+)
  - [Echo](https://echo.labstack.com/) Web Framework
  - [PostgreSQL](https://www.postgresql.org/) with [pgx/v5](https://github.com/jackc/pgx)
  - [godotenv](https://github.com/joho/godotenv) for environment variable management
## Setup & Installation

### Prerequisites

- Go installed on your system.
- PostgreSQL server running.

### Database Setup

Various SQL scripts are provided for setting up the database schema and seeding data:

1. **Homework 2 Setup**:
   - Navigate to `hw2/db prep/`
   - Run `create_db.sql`, `create_tables.sql`, and `insert_data.sql`.
2. **Homework 4 Setup**:
   - Navigate to `hw4/db_prep/`
   - Run `add_table.sql` and `insert_data.sql` to extend the schema for attendance tracking.

### Running the Projects

#### Go Backends (`hw3` or `hw4`)

1. Navigate to the desired project directory (e.g., `cd hw4`).
2. Copy `env.example` to `.env` and configure your database credentials.
3. Run the application:
   ```bash
   go run cmd/main.go
   ```
   The server will start on `http://localhost:8080`.


## API Endpoints (hw4)

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| GET | `/student/:id` | Get student details by ID |
| GET | `/all_class_schedule` | Get all group schedules |
| GET | `/schedule/group/:id` | Get schedule for a specific group |
| POST | `/attendance/schedule` | Record student attendance |
| GET | `/attendance/schedule/:id` | Get attendance for a specific schedule |
| GET | `/attendance/student/:id` | Get attendance history for a student |
