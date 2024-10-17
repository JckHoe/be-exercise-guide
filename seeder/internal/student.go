package seeder

import (
	"fmt"
	"time"

	"be-exerise-go-mod/repository"

	"be-exerise-go-mod/.gen/be-exercise/public/model"

	_ "github.com/lib/pq"
)

type Student struct {
	studentRepo    *repository.StudentRepository
	departmentRepo *repository.DepartmentRepository
}

func NewStudent(
	studentRepo *repository.StudentRepository,
	departmentRepo *repository.DepartmentRepository,
) *Student {
	return &Student{
		studentRepo:    studentRepo,
		departmentRepo: departmentRepo,
	}
}

func (s *Student) Seed(num int32) {
	var departmentIds = s.departmentRepo.GetDepartmentIDs()

	var studentModelLinks []model.Student
	for range num {
		now := time.Now().UTC()
		modelLink := model.Student{
			FirstName:    "s.faker.FirstName()",
			LastName:     "s.faker.LastName()",
			Dob:          now,
			Email:        "s.faker.Email()",
			DepartmentID: &departmentIds[len(departmentIds)],
		}
		studentModelLinks = append(studentModelLinks, modelLink)
	}

	s.studentRepo.InsertMultipleStudents(studentModelLinks)
	fmt.Println("Finish seeding Students")
}
