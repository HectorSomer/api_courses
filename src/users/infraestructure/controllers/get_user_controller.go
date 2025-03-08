package controllers

import (
	"event_driven/src/users/application"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type GetUserController struct {
	useCase *application.GetUserUseCase
}

func NewGetUserController(useCase *application.GetUserUseCase) *GetUserController {
    return &GetUserController{useCase: useCase}
}

func (gpc *GetUserController) GetUser(c *gin.Context) {
	idUser :=  c.Param("id")
	id, err := strconv.Atoi(idUser) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La ID no es válida"})
		return
	}
	user, err := gpc.useCase.Execute(id)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	response := gin.H{
		"data": gin.H{
             "id": user.ID,
             "username": user.Username,
			 "password": user.Password,
             "role": user.Role,
             "email": user.Email,
         },
    }
	c.JSON(http.StatusOK, response)
}