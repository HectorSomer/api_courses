package dependencies

import (
    "event_driven/src/courses/application"
    "event_driven/src/courses/infraestructure/controllers"
    "fmt"
    "event_driven/src/courses/infraestructure"
    "github.com/gin-gonic/gin"
)

func InitCourses(r *gin.Engine) {
    fmt.Println("Iniciando cursos")
    ps := infraestructure.NewMySql()
    createCourseUseCase := application.NewCreateCourseUseCase(ps)
    createCourseController := controllers.NewCreateCourseController(createCourseUseCase)
    publisher := infraestructure.NewRabbitMQPublisher()
    registrateToCourseUseCase := application.NewRegistrateToCourseUseCase(ps, publisher)
    registrateToCourseController := controllers.NewRegistrateToCourseController(registrateToCourseUseCase)
    getCoursesController := application.NewGetCoursesUseCase(ps)
    get_courses_controller := controllers.NewGetCoursesController(getCoursesController)
    infraestructure.RegisterCourseRouter(r, createCourseController, registrateToCourseController, get_courses_controller)
}
