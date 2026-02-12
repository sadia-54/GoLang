package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sadia-54/student-management/models"
	"github.com/sadia-54/student-management/services"
)

type StudentHandler struct {
	service *services.StudentService
}

//constructor for student handler
func NewStudentHandler(service *services.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

// create student 
func (h *StudentHandler) CreateStudent(c echo.Context) error {
	var student models.Student

	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	if err := h.service.CreateStudent(&student); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create student"})
	}

	return c.JSON(http.StatusCreated, student)
}

// get all students
func (h *StudentHandler) GetAllStudents(c echo.Context) error {
	students, err := h.service.GetAllStudents()	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve students"})
	}

	return c.JSON(http.StatusOK, students)
}

// get student by id
func (h *StudentHandler) GetStudentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid student ID"})
	}

	student, err := h.service.GetStudentByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Student not found"})
	}

	return c.JSON(http.StatusOK, student)
}

// update student
func (h *StudentHandler) UpdateStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid student ID"})
	}

	var student models.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	student.ID = uint(id) 

	if err := h.service.UpdateStudent(&student); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update student"})
	}

	return c.JSON(http.StatusOK, student)
}

// delete student
func (h *StudentHandler) DeleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid student ID"})
	}

	if err := h.service.DeleteStudent(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete student"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Student deleted successfully"})
}