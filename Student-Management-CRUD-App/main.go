package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sadia-54/student-management/config"
	"github.com/sadia-54/student-management/models"
)


func main () {
	e := echo.New() 

	// connect to database
	config.ConnectDB()

	// auto migrate the student model
	err := config.GormDB.AutoMigrate(&models.Student{})
	if err != nil {
		e.Logger.Fatal("Failed to auto-migrate student model:", err)
	}
	log.Println("Student table auto-migrated successfully!")

	e.Use(middleware.RequestLogger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is a simple student management CRUD application!")
	})

	// start the server
	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal("Failed to start server:", err)
	}

}
