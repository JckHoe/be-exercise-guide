package seeder

import (
	"fmt"
	"time"

	"be-exerise-go-mod/repository"

	"be-exerise-go-mod/.gen/be-exercise/public/model"

	_ "github.com/lib/pq"
)

// TODO this not used yet, clean up and use in future
// type faker interface {
// 	FirstName() string
// 	LastName() string
// 	DateRange(time.Time, time.Time) time.Time
// 	Email() string
// }
//
// type randomizer interface {
// 	Intn(int) int
// }

type Teacher struct {
	teacherRepo    *repository.TeacherRepository
	departmentRepo *repository.DepartmentRepository
}

func NewTeacher(
	teacherRepo *repository.TeacherRepository,
	departmentRepo *repository.DepartmentRepository,
) *Teacher {
	return &Teacher{
		departmentRepo: departmentRepo,
		teacherRepo:    teacherRepo,
	}
}

func (s *Teacher) Seed(num int32) {
	departmentIds := s.departmentRepo.GetDepartmentIDs()

	var teacherModelLinks []model.Teacher
	for range num {
		now := time.Now().UTC()
		modelLink := model.Teacher{
			FirstName:    "s.faker.FirstName()",
			LastName:     "s.faker.LastName()",
			Dob:          now,
			Email:        "s.faker.Email()",
			DepartmentID: &departmentIds[len(departmentIds)],
		}
		teacherModelLinks = append(teacherModelLinks, modelLink)
	}
	s.teacherRepo.InsertMultipleTeachers(teacherModelLinks)
	fmt.Println("Finish seeding Teachers")
}

func (s *Teacher) ClearAll() {
	s.teacherRepo.ClearAllTeachers()
}
