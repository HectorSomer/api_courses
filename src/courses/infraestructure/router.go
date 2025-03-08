package infraestructure

import (
	"event_driven/src/courses/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterCourseRouter(r *gin.Engine, createCourseController *controllers.CreateCourseController, registrateToCourseController *controllers.RegistrateToCourseController, getCoursesController *controllers.GetCoursesController) {
	routes := r.Group("/v1/courses")
	{
		routes.POST("/", createCourseController.CreateCourse)
		routes.POST("/registration/", registrateToCourseController.RegistrateToCourse)
		routes.GET("/", getCoursesController.GetCourses)
	}

}