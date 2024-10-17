package seeder

import (
	"fmt"
	"math/rand"

	"be-exerise-go-mod/.gen/be-exercise/public/model"
	"be-exerise-go-mod/repository"
)

type Score struct {
	scoreRepo      *repository.ScoreRepository
	teacherRepo    *repository.TeacherRepository
	studentRepo    *repository.StudentRepository
	submissionRepo *repository.SubmissionRepository
}

func NewScore(teacherRepo *repository.TeacherRepository, studentRepo *repository.StudentRepository, submissionRepo *repository.SubmissionRepository, scoreRepo *repository.ScoreRepository) *Score {
	return &Score{
		teacherRepo:    teacherRepo,
		studentRepo:    studentRepo,
		submissionRepo: submissionRepo,
		scoreRepo:      scoreRepo,
	}
}

func (s *Score) Seed() {
	teachers := s.teacherRepo.GetAllTeachers()
	submissions := s.submissionRepo.GetSubmissionIDsAndDepartmentIDs()

	// group teacher by department
	teachersByDepartment := make(map[int32][]int32)
	for _, teacher := range teachers {
		deptID := *teacher.DepartmentID
		teachersByDepartment[deptID] = append(teachersByDepartment[deptID], int32(teacher.ID))
	}

	var scoreModelLinks []model.Score
	for _, submission := range submissions {
		// skipping assignment with submission time over due date
		// currently using UTC time as a cutoff, can review if this is a correct appraoch or not
		if !submission.IsAssignment || (submission.IsAssignment && submission.AssignmentDueDate.AddDate(0, 0, 1).Before(submission.SubmittedAt)) {
			modelLink := model.Score{
				SubmissionID: &submission.ID,
				TeacherID:    &teachersByDepartment[submission.DepartmentID][rand.Intn(len(teachersByDepartment[submission.DepartmentID]))],
				Value:        int32(rand.Intn(101)),
			}
			scoreModelLinks = append(scoreModelLinks, modelLink)
		}
	}
	// Define the batch size
	batchSize := 5000

	// Process submissions in batches
	for i := 0; i < len(scoreModelLinks); i += batchSize {
		end := i + batchSize
		if end > len(scoreModelLinks) {
			end = len(scoreModelLinks)
		}
		batch := scoreModelLinks[i:end]
		s.scoreRepo.InsertMultipleScores(batch)
	}

	fmt.Println("Finish seeding Score")
}

func (s *Score) ClearAll() {
	s.scoreRepo.ClearAllScores()
}
