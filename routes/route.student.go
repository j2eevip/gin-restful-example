package route

import (
	"github.com/gin-gonic/gin"
	createStudent "github.com/j2eevip/gin-restful-example/controllers/student-controllers/create"
	deleteStudent "github.com/j2eevip/gin-restful-example/controllers/student-controllers/delete"
	resultStudent "github.com/j2eevip/gin-restful-example/controllers/student-controllers/result"
	resultsStudent "github.com/j2eevip/gin-restful-example/controllers/student-controllers/results"
	updateStudent "github.com/j2eevip/gin-restful-example/controllers/student-controllers/update"
	handlerCreateStudent "github.com/j2eevip/gin-restful-example/handlers/student-handlers/create"
	handlerDeleteStudent "github.com/j2eevip/gin-restful-example/handlers/student-handlers/delete"
	handlerResultStudent "github.com/j2eevip/gin-restful-example/handlers/student-handlers/result"
	handlerResultsStudent "github.com/j2eevip/gin-restful-example/handlers/student-handlers/results"
	handlerUpdateStudent "github.com/j2eevip/gin-restful-example/handlers/student-handlers/update"
	middleware "github.com/j2eevip/gin-restful-example/middlewares"
	"gorm.io/gorm"
)

func InitStudentRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Student
	*/
	createStudentRepository := createStudent.NewRepositoryCreate(db)
	createStudentService := createStudent.NewServiceCreate(createStudentRepository)
	createStudentHandler := handlerCreateStudent.NewHandlerCreateStudent(createStudentService)

	resultsStudentRepository := resultsStudent.NewRepositoryResults(db)
	resultsStudentService := resultsStudent.NewServiceResults(resultsStudentRepository)
	resultsStudentHandler := handlerResultsStudent.NewHandlerResultsStudent(resultsStudentService)

	resultStudentRepository := resultStudent.NewRepositoryResult(db)
	resultStudentService := resultStudent.NewServiceResult(resultStudentRepository)
	resultStudentHandler := handlerResultStudent.NewHandlerResultStudent(resultStudentService)

	deleteStudentRepository := deleteStudent.NewRepositoryDelete(db)
	deleteStudentService := deleteStudent.NewServiceDelete(deleteStudentRepository)
	deleteStudentHandler := handlerDeleteStudent.NewHandlerDeleteStudent(deleteStudentService)

	updateStudentRepository := updateStudent.NewRepositoryUpdate(db)
	updateStudentService := updateStudent.NewServiceUpdate(updateStudentRepository)
	updateStudentHandler := handlerUpdateStudent.NewHandlerUpdateStudent(updateStudentService)

	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/student", createStudentHandler.CreateStudentHandler)
	groupRoute.GET("/student", resultsStudentHandler.ResultsStudentHandler)
	groupRoute.GET("/student/:id", resultStudentHandler.ResultStudentHandler)
	groupRoute.DELETE("/student/:id", deleteStudentHandler.DeleteStudentHandler)
	groupRoute.PUT("/student/:id", updateStudentHandler.UpdateStudentHandler)
}
