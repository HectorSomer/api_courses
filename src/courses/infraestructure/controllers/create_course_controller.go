package controllers

import (
	"event_driven/src/courses/application"
	"event_driven/src/courses/domain/entities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCourseController struct {
	useCase *application.CreateCourseUseCase
}

func NewCreateCourseController(useCase *application.CreateCourseUseCase) *CreateCourseController {
	return &CreateCourseController{useCase: useCase}
}
func (cc *CreateCourseController) CreateCourse(c *gin.Context) {
	var course entities.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Creando curso")
	courseCreated, err := cc.useCase.Execute(&course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
	}
	response := gin.H{
		"posts": gin.H{
			"id": courseCreated.ID,
			"name": courseCreated.Name,
			"description": courseCreated.Description,
            "teacher": courseCreated.Teacher,
			"idUser": courseCreated.IDUser,
        },
	}
	c.JSON(http.StatusCreated, response)
}