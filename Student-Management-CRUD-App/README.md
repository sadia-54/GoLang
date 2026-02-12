# Student Management CRUD App

A simple **RESTful API application** for managing students using **Go**, **Echo framework**, **GORM**, **PostgreSQL**, **Validator** and **Zerolog** for structured logging.

---

## Table of Contents

- [Features](#features)  
- [Project Structure](#project-structure)  
- [Tech Stack](#tech-stack)  
- [Getting Started](#getting-started)  
- [Environment Variables](#environment-variables)  
- [Database](#database)  
- [API Endpoints](#api-endpoints)  
- [Logging](#logging)  
- [Validation](#validation)  
- [License](#license)  

---

## Features

- Create, Read, Update, and Delete students  
- Auto migration of database tables using GORM  
- Structured logging using [Zerolog]  
- Request logging middleware for API monitoring  
- Input validation using [Go Playground Validator]  
- Modular project structure for maintainability  

---

## Project Structure
```
Student-Management-CRUD-App/
│
├── api/              # Handlers for student routes
├── config/           # DB connection and validator setup
├── logger/           # Zerolog configuration
├── models/           # GORM models
├── repositories/     # Database access layer
├── services/         # Business logic layer
├── tmp/              # Temporary files (if needed)
├── .env              # Environment variables
├── app.log           # Application log file
├── go.mod
├── go.sum
├── main.go           # Entry point
└── requirements.txt
```

---

## Tech Stack

- **Language:** Go 1.25.5
- **Framework:** Echo v4
- **ORM:** GORM v1.31.1
- **Database:** PostgreSQL
- **Logger:** Zerolog v1.34.0
- **Validation:** Go Playground Validator v10

---

## Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/sadia-54/GoLang.git
cd Student-Management-CRUD-App
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Configure environment variables

Create a `.env` file in the root directory and set your database variables:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=student_db
DB_SSLMODE=disable
DB_TIMEZONE=Asia/Dhaka
```

### 4. Run the application
```bash
go run main.go
```

The server will start on `http://localhost:8080`.

---

## Database

- **Auto migration** is enabled; tables are created automatically using GORM.
- **Model:** `Student`

| Field      | Type     | Description                |
|------------|----------|----------------------------|
| ID         | uint     | Primary key                |
| Name       | string   | Student full name          |
| Age        | int      | Student age                |
| Email      | string   | Unique student email       |
| Department | string   | Student department         |
| Session    | string   | Academic session           |
| CreatedAt  | time     | Record creation timestamp  |
| UpdatedAt  | time     | Last update timestamp      |

---

## API Endpoints

| Method | Endpoint         | Description           | Request Body         |
|--------|------------------|-----------------------|----------------------|
| POST   | `/students`      | Create a new student  | JSON student object  |
| GET    | `/students`      | Get all students      | None                 |
| GET    | `/students/:id`  | Get a student by ID   | None                 |
| PUT    | `/students/:id`  | Update a student      | JSON student object  |
| DELETE | `/students/:id`  | Delete a student      | None                 |

### Example Request Body (POST/PUT)
```json
{
  "name": "Sadia Rahman",
  "age": 22,
  "email": "sadia@example.com",
  "department": "Computer Science",
  "session": "2020-2021"
}
```

---

## Logging

- **Request logs:** Automatically logged via middleware
- **Structured logs:** Using `logger.Logger` in handlers
- **Log output:** Saved in `app.log` and printed to console

### Example Log Entry
```json
{
  "level": "info",
  "method": "POST",
  "path": "/students",
  "status": 201,
  "latency": "2.123ms",
  "time": "2026-02-12T22:53:51+06:00",
  "message": "Student created successfully",
  "email": "sadia@example.com"
}
```

---

## Validation

- Input validation using [Go Playground Validator](https://github.com/go-playground/validator)
- Applied in `CreateStudent` and `UpdateStudent` handlers
- Returns proper error messages for invalid payloads

---

## License

This project is licensed under the MIT License.

---

