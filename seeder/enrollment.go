package seeder

import (
	"database/sql"
	"fmt"
	"math/rand"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

func EnrollmentSeeder(db *sql.DB) {
	minCourseEnroll := 3
	studentRepository := repository.NewStudentRepository(db)
	studentIDs := studentRepository.GetStudentIDs()
	courseRepository := repository.NewCourseRepository(db)
	courseIDs := courseRepository.GetCourseIDs()
	// increasing the ratio to approved vs false to 4:1
	approvedOption := []bool{true, true, true, false}
	enrollmentRepository := repository.NewEnrollmentRepository(db)

	var enrollmentModelLinks []model.Enrollment
	for _, studentID := range studentIDs {
		coursesEnroll := rand.Intn(5) + minCourseEnroll
		pickedCourseIDs := pickRandomIDs(courseIDs, coursesEnroll)
		for _, cIDs := range pickedCourseIDs {
			if !enrollmentRepository.IsStudentEnrolledInCourse(studentID, cIDs) {
				modelLink := model.Enrollment{
					StudentID: &studentID,
					CourseID:  &cIDs,
					Approved:  &approvedOption[rand.Intn(len(approvedOption))],
				}
				enrollmentModelLinks = append(enrollmentModelLinks, modelLink)
			}
		}
	}
	enrollmentRepository.InsertMultipleEnrollments(enrollmentModelLinks)
	fmt.Println("Finish seeding Enrollment")
}

func pickRandomIDs(arr []int32, count int) []int32 {
	// Create a copy of the original array to avoid modifying it
	temp := make([]int32, len(arr))
	copy(temp, arr)

	// Shuffle the temporary array
	rand.Shuffle(len(temp), func(i, j int) {
		temp[i], temp[j] = temp[j], temp[i]
	})

	// Return the first 'count' elements
	if count > len(temp) {
		count = len(temp)
	}
	return temp[:count]
}
