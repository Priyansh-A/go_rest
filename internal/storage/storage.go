package storage

import "github.com/Priyansh_A/go-rest-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	UpdateStudentById(id int64, name, email string, age int) error
	GetStudents() ([]types.Student, error)
	DeleteStudentsById(int64) error
}
