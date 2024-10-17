package seeder

import (
	"fmt"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

type Department struct {
	departmentRepo *repository.DepartmentRepository
}

func NewDepartment(departmentRepo *repository.DepartmentRepository) *Department {
	return &Department{
		departmentRepo: departmentRepo,
	}
}

func (s *Department) Seed() {
	departmentNames := []string{
		"Computer Science",
		"Biology",
		"Chemistry",
		"Physics",
		"Mathematics",
		"Economics",
		"English Literature",
		"History",
		"Psychology",
		"Political Science",
	}

	var departmentModelLinks []model.Department
	departmentIds := s.departmentRepo.GetDepartmentIDs()

	if len(departmentIds) == 0 {
		for _, name := range departmentNames {
			modelLink := model.Department{
				Name: name,
			}
			departmentModelLinks = append(departmentModelLinks, modelLink)
		}
		s.departmentRepo.InsertMultipleDepartments(departmentModelLinks)
		fmt.Println("Finish seeding Department")
	} else {
		fmt.Println("Already created Departments.  Skipping....")
	}
}
