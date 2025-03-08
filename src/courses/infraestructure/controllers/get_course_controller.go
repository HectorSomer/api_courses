package controllers

import (
	"event_driven/src/courses/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetCoursesController struct {
	useCase *application.GetCoursesUseCase
}

func NewGetCoursesController(useCase *application.GetCoursesUseCase) *GetCoursesController {
    return &GetCoursesController{useCase: useCase}
}

func (gcc *GetCoursesController) GetCourses(c *gin.Context) {
    courses, err := gcc.useCase.Execute()

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    if len(*courses) == 0 {
        c.JSON(http.StatusOK, gin.H{
            "data_size": 0,
            "message": "No se encontraron cursos.",
        })
        return
    }
    var response []gin.H
    for _, course := range *courses {
        courseResponse := gin.H{
            "idCourse":    course.ID,
            "name":        course.Name,
            "description": course.Description,
            "teacher":     course.Teacher,
            "idUser":      course.IDUser,
        }
        response = append(response, courseResponse)
    }
    c.JSON(http.StatusOK, gin.H{
        "data": response,
    })
}
