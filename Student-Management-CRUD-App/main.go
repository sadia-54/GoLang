package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sadia-54/student-management/config"
	"github.com/sadia-54/student-management/models"
	"github.com/sadia-54/student-management/services"
	"github.com/sadia-54/student-management/api"
	"github.com/sadia-54/student-management/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/sadia-54/student-management/logger"
)


func main () {

	// init logger
	logger.InitLogger()
	logger.Logger.Info().Msg("Starting student management server")

	e := echo.New() 

	// register validator
	e.Validator = &config.CustomValidator{
		Validator: validator.New(),
	}

	// connect to database
	config.ConnectDB()

	// auto migrate the student model
	err := config.GormDB.AutoMigrate(&models.Student{})
	if err != nil {
		e.Logger.Fatal("Failed to auto-migrate student model:", err)
	}
	log.Println("Student table auto-migrated successfully!")

	// initialize repository, service, and handler
	studentRepo := repositories.NewStudentRepository()
	studentService := services.NewStudentService(studentRepo)
	studentHandler := api.NewStudentHandler(studentService)

	e.Use(middleware.RequestLogger())

	// use zerolog for request logging
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			logger.Logger.Info().
				Str("method", c.Request().Method).
				Str("path", c.Request().URL.Path).
				Int("status", c.Response().Status).
				Dur("latency", time.Since(start)).
				Msg("Request completed")
			return err
		}
	})


	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is a simple student management CRUD application!")
	})

	// student routes
	e.POST("/students", studentHandler.CreateStudent)
	e.GET("/students", studentHandler.GetAllStudents)
	e.GET("/students/:id", studentHandler.GetStudentByID)
	e.PUT("/students/:id", studentHandler.UpdateStudent)
	e.DELETE("/students/:id", studentHandler.DeleteStudent)

	// start the server
	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal("Failed to start server:", err)
	}

}
