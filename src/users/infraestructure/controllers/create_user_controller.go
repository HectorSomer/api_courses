package controllers

import (
	"event_driven/src/users/application"
	"event_driven/src/users/domain/entities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	useCase *application.CreateUserUseCase
}

func NewCreateUserController(useCase *application.CreateUserUseCase) *CreateUserController {
    return &CreateUserController{useCase: useCase}
}

func(cu *CreateUserController) CreateUser(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Creando usuario")
	userCreated, err := cu.useCase.Execute(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	response := gin.H{
	"posts": gin.H{
      		"id": userCreated.ID,
			"username": userCreated.Username,
			"email": userCreated.Email,
			"role": userCreated.Role,
          },
	}
	c.JSON(http.StatusCreated, response)
}