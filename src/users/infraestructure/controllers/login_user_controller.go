package controllers

import (
	"event_driven/src/users/application"
	"event_driven/src/users/application/entities"

	"github.com/gin-gonic/gin"
)

type LoginUserController struct {
	useCase *application.LoginUseCase
}

func NewLoginUserController(useCase *application.LoginUseCase) *LoginUserController {
	return &LoginUserController{useCase: useCase}
}
func (luc *LoginUserController) LoginUser(c *gin.Context){
	var login entities.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
        return
	}
    user, err := luc.useCase.Execute(login.Email, login.Password)
	if err != nil {
        c.JSON(401, gin.H{"error": "Invalid email or password"})
        return
    }
	response := gin.H{
		"post": gin.H{
             "id": user.ID,
             "username": user.Username,
             "role": user.Role,
             "email": user.Email,
         },
    }
    c.JSON(200, response)
}