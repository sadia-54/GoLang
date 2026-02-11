package services

import (
	"time"
	"github.com/sadia-54/student-management/models"
	"github.com/sadia-54/student-management/repositories"
)

type StudentService struct {
	repo *repositories.StudentRepository
}

func NewStudentService(repo *repositories.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

// Create a new student
func (s *StudentService) CreateStudent(student *models.Student) error {
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()
	return s.repo.Create(student)
}

// Get all students
func (s *StudentService) GetAllStudents() ([]models.Student, error) {
	return s.repo.GetAll()
}	

// Get a student by ID
func (s *StudentService) GetStudentByID(id uint) (models.Student, error) {
	return s.repo.GetByID(id)
}	

// Update a student
func (s *StudentService) UpdateStudent(student *models.Student) error	 {
	student.UpdatedAt = time.Now()
	return s.repo.Update(student)
}

// Delete a student		
func (s *StudentService) DeleteStudent(id uint) error {
	return s.repo.Delete(id)
}