package repositories

import (
	"github.com/sadia-54/student-management/models"
	"github.com/sadia-54/student-management/config"

)

type StudentRepository struct {}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

// Create a new student
func (r *StudentRepository) Create(student *models.Student) error {
	result := config.GormDB.Create(student)
	return result.Error
}

// Get all students
func (r *StudentRepository) GetAll() (students []models.Student, err error) {
	result := config.GormDB.Find(&students)
	return students, result.Error
}

// Get a student by ID
func (r *StudentRepository) GetByID(id uint) (student models.Student, err error) {
	result := config.GormDB.First(&student, id)
	return student, result.Error
}

// Update a student
func (r *StudentRepository) Update(student *models.Student) error {
	result := config.GormDB.Save(student)
	return result.Error
}

// Delete a student
func (r *StudentRepository) Delete(id uint) error {
	result := config.GormDB.Delete(&models.Student{}, id)
	return result.Error
}

