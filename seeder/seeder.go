package seeder

import (
	"database/sql"

	"be-exerise-go-mod/repository"
	"be-exerise-go-mod/seeder/internal"

	_ "github.com/lib/pq"
)

type Seeder struct {
	// Repository
	assignmentRepo   *repository.AssignmentRepository
	teacherRepo      *repository.TeacherRepository
	studentRepo      *repository.StudentRepository
	departmentRepo   *repository.DepartmentRepository
	courseRepo       *repository.CourseRepository
	enrollmentRepo   *repository.EnrollmentRepository
	gradeSettingRepo *repository.GradeSettingRepository
	examRepo         *repository.ExamRepository
	scoreRepo        *repository.ScoreRepository
	submissionRepo   *repository.SubmissionRepository

	// Seeder functions
	assignment   *seeder.Assignment
	teacher      *seeder.Teacher
	student      *seeder.Student
	department   *seeder.Department
	course       *seeder.Course
	enrollment   *seeder.Enrollment
	gradeSetting *seeder.GradeSetting
	exam         *seeder.Exam
	score        *seeder.Score
	submission   *seeder.Submission
}

func NewSeeder(db *sql.DB) *Seeder {
	// TODO refactor this one DB pass create one repository struct
	assignmentRepo := repository.NewAssignmentRepository(db)
	teacherRepo := repository.NewTeacherRepository(db)
	studentRepo := repository.NewStudentRepository(db)
	departmentRepo := repository.NewDepartmentRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	enrollmentRepo := repository.NewEnrollmentRepository(db)
	gradeSettingRepo := repository.NewGradeSettingRepository(db)
	examRepo := repository.NewExamRepository(db)
	scoreRepo := repository.NewScoreRepository(db)
	submissionRepo := repository.NewSubmissionRepository(db)

	return &Seeder{
		// Repository
		// TODO refactor this one DB pass create one repository struct
		assignmentRepo:   assignmentRepo,
		teacherRepo:      teacherRepo,
		studentRepo:      studentRepo,
		departmentRepo:   departmentRepo,
		courseRepo:       courseRepo,
		enrollmentRepo:   enrollmentRepo,
		gradeSettingRepo: gradeSettingRepo,
		examRepo:         examRepo,
		scoreRepo:        scoreRepo,
		submissionRepo:   submissionRepo,

		// Seeders
		assignment:   seeder.NewAssignment(assignmentRepo, courseRepo),
		teacher:      seeder.NewTeacher(teacherRepo, departmentRepo),
		student:      seeder.NewStudent(studentRepo, departmentRepo),
		department:   seeder.NewDepartment(departmentRepo),
		course:       seeder.NewCourse(courseRepo, departmentRepo, teacherRepo),
		enrollment:   seeder.NewEnrollment(studentRepo, courseRepo, enrollmentRepo),
		gradeSetting: seeder.NewGradeSetting(courseRepo, gradeSettingRepo),
		exam:         seeder.NewExam(courseRepo, examRepo),
		score:        seeder.NewScore(teacherRepo, studentRepo, submissionRepo, scoreRepo),
		submission:   seeder.NewSubmission(assignmentRepo, courseRepo, examRepo, submissionRepo, enrollmentRepo),
	}
}

func (s *Seeder) SeedAll(teacherSize, studentSize int32) {
	s.department.Seed()
	s.teacher.Seed(teacherSize)
	s.course.Seed()
	s.gradeSetting.Seed()
	s.student.Seed(studentSize)
	s.enrollment.Seed()
	s.assignment.Seed()
	s.exam.Seed()
	s.submission.Seed()
	s.score.Seed()
}

func (s *Seeder) DeseedAll() {
	s.scoreRepo.ClearAllScores()
	s.submissionRepo.ClearAllSubmissions()
	s.examRepo.ClearAllExams()
	s.assignmentRepo.ClearAllAssignments()
	s.enrollmentRepo.ClearAllEnrollments()
	s.studentRepo.ClearAllStudents()
	s.gradeSettingRepo.ClearAllGradeSettings()
	s.courseRepo.ClearAllCourses()
	s.teacherRepo.ClearAllTeachers()
	s.departmentRepo.ClearAllDepartments()
}
