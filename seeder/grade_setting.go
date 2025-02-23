package seeder

import (
	"database/sql"
	"fmt"
	"math/rand"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

func GradeSettingSeeder(db *sql.DB) {
	courseRepository := repository.NewCourseRepository(db)
	gradeSettingRepository := repository.NewGradeSettingRepository(db)
	courseIDs := courseRepository.GetCourseIDs()
	var gradeSettingModelLinks []model.GradeSetting

	assignmentPercentRandomChoice := []int32{20, 25, 30, 35, 40, 45}
	passingGradeRandomChoice := []int32{60, 65, 70, 75, 80}

	for _, courseID := range courseIDs {
		assignmentPercent := assignmentPercentRandomChoice[rand.Intn(len(assignmentPercentRandomChoice))]
		modelLink := model.GradeSetting{
			AssignmentPercent: assignmentPercent,
			ExamPercent:       100 - assignmentPercent,
			PassingGrade:      passingGradeRandomChoice[rand.Intn(len(passingGradeRandomChoice))],
			CourseID:          &courseID,
		}
		gradeSettingModelLinks = append(gradeSettingModelLinks, modelLink)
	}
	gradeSettingRepository.InsertMultipleGradeSettings(gradeSettingModelLinks)
	fmt.Println("Finish seeding GradeSetting")
}
