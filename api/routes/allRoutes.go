package routes

import (
	_ "github.com/BetterGR/api-gateway/docs" // This is required for swagger docs
	"github.com/BetterGR/api-gateway/pkg/controllers"
	"github.com/gin-gonic/gin"
)

// @title BetterGR API Gateway
// @version 1.0
// @description Learning Management System API Gateway
// @host localhost:1234
// @BasePath /
// @schemes http
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @Summary User login
// @Description Authenticate user and receive JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Login credentials"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/login [post]
func login(c *gin.Context) {
	controllers.HandleCallback(c)
}

// The following endpoints are not implemented yet:

// @Summary Get student courses
// @Description Get all courses for a student
// @Tags Students
// @Accept json
// @Produce json
// @Param studentId path string true "Student ID"
// @Success 200 {object} models.getStudentCoursesResponse
// @Router /api/students/{studentId}/courses [get]
func getStudentCourses(c *gin.Context) {}

// @Summary Get student's course grades
// @Description Get grades for a student's course
// @Tags Grades
// @Accept json
// @Produce json
// @Param studentId path string true "Student ID"
// @Param courseId path string true "Course ID"
// @Success 200 {object} models.getStudentCourseGradesResponse
// @Router /api/grades/{student_id}/{courseId} [get]
func getStudentCourseGrades(c *gin.Context) {}

// @Summary Get all student grades
// @Description Get all grades for a student
// @Tags Grades
// @Accept json
// @Produce json
// @Param studentId path string true "Student ID"
// @Success 200 {object} models.NotImplementedYet
// @Router /api/grades/{student_id} [get]
func getAllStudentGrades(c *gin.Context) {}

// @Summary Create new student
// @Description Create a new student
// @Tags Students
// @Accept json
// @Produce json
// @Param credentials body models.createStudentRequest true "Student data"
// @Success 200 {object} models.createStudentResponse
// @Router /api/students [post]
func createStudent(c *gin.Context) {}

// @Summary Get student details
// @Description Get details for a student
// @Tags Students
// @Accept json
// @Produce json
// @Param studentId path string true "Student ID"
// @Success 200 {object} models.getStudentResponse
// @Router /api/students/{studentId} [get]
func getStudent(c *gin.Context) {}

// @Summary Update student
// @Description Update a student's information
// @Tags Students
// @Accept json
// @Produce json
// @Param studentId path string true "Student ID"
// @Param credentials body models.updateStudentRequest true "Student ID"
// @Success 200 {object} models.updateStudentResponse
// @Router /api/students/{studentId} [put]
func updateStudent(c *gin.Context) {}

// @Summary Delete student
// @Description Delete a student
// @Tags Students
// @Accept json
// @Produce json
// @Param studentId path string true "Student ID"
// @Success 200 {object} models.deleteStudentResponse
// @Router /api/students/{studentId} [delete]
func deleteStudent(c *gin.Context) {}

// @Summary Get course details
// @Description Get details for a course
// @Tags Courses
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Success 200 {object} models.getCourseResponse
// @Router /api/courses/{courseId} [get]
func getCourse(c *gin.Context) {}

// @Summary Create course
// @Description Create a new course
// @Tags Courses
// @Accept json
// @Produce json
// @Param credentials body models.createCourseRequest true "Student ID"
// @Success 200 {object} models.createCourseResponse
// @Router /api/courses [post]
func createCourse(c *gin.Context) {}

// @Summary Update course
// @Description Update a course's information
// @Tags Courses
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Param credentials body models.updateCourseRequest true "Student ID"
// @Success 200 {object} models.updateCourseResponse
// @Router /api/courses/{courseId} [put]
func updateCourse(c *gin.Context) {}

// @Summary Add student to course
// @Description Add a student to a course
// @Tags Courses
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Param credentials body models.addStudentToCourseRequest true "Student ID"
// @Success 200 {object} models.addStudentToCourseResponse
// @Router /api/courses/{courseId}/students [post]
func addStudentToCourse(c *gin.Context) {}

// @Summary Remove student from course
// @Description Remove a student from a course
// @Tags Courses
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Param studentId path string true "Student ID"
// @Success 200 {object} models.removeStudentFromCourseResponse
// @Router /api/courses/{courseId}/students/{studentId} [delete]
func removeStudentFromCourse(c *gin.Context) {}

// @Summary Add staff to course
// @Description Add a staff member to a course
// @Tags Courses
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Param credentials body models.addStaffToCourseRequest true "Student ID"
// @Success 200 {object} models.addStaffToCourseResponse
// @Router /api/courses/{courseId}/staff [post]
func addStaffToCourse(c *gin.Context) {}

// @Summary Remove staff from course
// @Description Remove a staff member from a course
// @Tags Courses
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Param staffId path string true "Staff ID"
// @Success 200 {object} models.removeStaffFromCourseResponse
// @Router /api/courses/{courseId}/staff/{staffId} [delete]
func removeStaffFromCourse(c *gin.Context) {}

// @Summary Delete course
// @Description Delete a course
// @Tags Courses
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Success 200 {object} models.deleteCourseResponse
// @Router /api/courses/{courseId} [delete]
func deleteCourse(c *gin.Context) {}

// @Summary Get course students
// @Description Get all students for a course
// @Tags Courses
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Success 200 {object} models.getCourseStudentsResponse
// @Router /api/courses/{courseId}/students [get]
func getCourseStudents(c *gin.Context) {}

// @Summary Get course staff
// @Description Get all staff members for a course
// @Tags Courses
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Success 200 {object} models.getCourseStaffResponse
// @Router /api/courses/{courseId}/staff [get]
func getCourseStaff(c *gin.Context) {}

// @Summary Add homework grade
// @Description Add a homework grade
// @Tags Grades
// @Accept json
// @Produce json
// @Param credentials body models.addHomeworkGradeRequest true "Student ID"
// @Success 200 {object} models.addHomeworkGradeResponse
// @Router /api/grades/homework [post]
func addHomeworkGrade(c *gin.Context) {}

// @Summary Add exam grade
// @Description Add an exam grade
// @Tags Grades
// @Accept json
// @Produce json
// @Param credentials body models.addExamGradeRequest true "Student ID"
// @Success 200 {object} models.addExamGradeResponse
// @Router /api/grades/exam [post]
func addExamGrade(c *gin.Context) {}

// @Summary Update homework grade
// @Description Update a homework grade
// @Tags Grades
// @Accept json
// @Produce json
// @Param credentials body models.updateHomeworkGradeRequest true "Student ID"
// @Success 200 {object} models.updateHomeworkGradeResponse
// @Router /api/grades/homework [put]
func updateHomeworkGrade(c *gin.Context) {}

// @Summary Update exam grade
// @Description Update an exam grade
// @Tags Grades
// @Accept json
// @Produce json
// @Param credentials body models.updateExamGradeRequest true "Student ID"
// @Success 200 {object} models.updateExamGradeResponse
// @Router /api/grades/exam [put]
func updateExamGrade(c *gin.Context) {}

// @Summary Delete homework grade
// @Description Delete a homework grade
// @Tags Grades
// @Accept json
// @Produce json
// @Param credentials body models.deleteHomeworkGradeRequest true "Student ID"
// @Success 200 {object} models.deleteHomeworkGradeResponse
// @Router /api/grades/homework [delete]
func deleteHomeworkGrade(c *gin.Context) {}

// @Summary Delete exam grade
// @Description Delete an exam grade
// @Tags Grades
// @Accept json
// @Produce json
// @Param credentials body models.deleteExamGradeRequest true "Student ID"
// @Success 200 {object} models.deleteExamGradeResponse
// @Router /api/grades/exam [delete]
func deleteExamGrade(c *gin.Context) {}

// @Summary Get homework grade
// @Description Get a specific homework grade for a student
// @Tags Grades
// @Accept json
// @Produce json
// @Param homeworkId path string true "Homework ID"
// @Success 200 {object} models.getHomeworkGradeResponse
// @Router /api/grades/homework/{homeworkId} [get]
func getHomeworkGrade(c *gin.Context) {}

// @Summary Get exam grade
// @Description Get a specific exam grade for a student
// @Tags Grades
// @Accept json
// @Produce json
// @Param examId path string true "Exam ID"
// @Success 200 {object} models.getExamGradeResponse
// @Router /api/grades/exam/{examId} [get]
func getExamGrade(c *gin.Context) {}

// @Summary Get course grade for a student
// @Description Get a specific course grade for a student
// @Tags Grades
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Success 200 {object} models.getStudentCourseGradeResponse
// @Router /api/grades/course/{courseId} [get]
func getStudentCourseGrade(c *gin.Context) {}

// @Summary Get homework
// @Description Get homework for a course
// @Tags Homework
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Success 200 {object} models.getHomeworkResponse
// @Router /api/homework/{courseId} [get]
func getHomework(c *gin.Context) {}

// @Summary Create homework
// @Description Create a new homework
// @Tags Homework
// @Accept json
// @Produce json
// @Param credentials body models.createHomeworkRequest true "Student ID"
// @Success 200 {object} models.createHomeworkResponse
// @Router /api/homework [post]
func createHomework(c *gin.Context) {}

// @Summary Get staff member
// @Description Get details for a staff member
// @Tags Staff
// @Accept json
// @Produce json
// @Param staffId path string true "Staff ID"
// @Success 200 {object} models.getStaffMemberResponse
// @Router /api/staff/{staffId} [get]
func getStaffMember(c *gin.Context) {}

// @Summary Get staff courses
// @Description Get all courses for a staff member
// @Tags Staff
// @Accept json
// @Produce json
// @Param staffId path string true "Staff ID"
// @Success 200 {object} models.getStaffCoursesResponse
// @Router /api/staff/{staffId}/courses [get]
func getStaffCourses(c *gin.Context) {}

// @Summary Create staff member
// @Description Create a new staff member
// @Tags Staff
// @Accept json
// @Produce json
// @Param credentials body models.createStaffMemberRequest true "Student ID"
// @Success 200 {object} models.createStaffMemberResponse
// @Router /api/staff [post]
func createStaffMember(c *gin.Context) {}

// @Summary Update staff member
// @Description Update a staff member's information
// @Tags Staff
// @Accept json
// @Produce json
// @Param staffId path string true "Staff ID"
// @Param credentials body models.updateStaffMemberRequest true "Student ID"
// @Success 200 {object} models.updateStaffMemberResponse
// @Router /api/staff/{staffId} [put]
func updateStaffMember(c *gin.Context) {}

// @Summary Delete staff member
// @Description Delete a staff member
// @Tags Staff
// @Accept json
// @Produce json
// @Param staffId path string true "Staff ID"
// @Success 200 {object} models.deleteStaffMemberResponse
// @Router /api/staff/{staffId} [delete]
func deleteStaffMember(c *gin.Context) {}
