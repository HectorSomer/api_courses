package infraestructure

import (
	"event_driven/src/users/application"
	"event_driven/src/users/infraestructure/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitUsers(r *gin.Engine){
	fmt.Println("Iniciando usuarios")
	ps := NewMySql()

	createUserUseCase := application.NewCreateUserUseCase(ps)
	create_user_controller := controllers.NewCreateUserController(createUserUseCase)
	getUserUseCase := application.NewGetUserUseCase(ps)
	get_user_controller := controllers.NewGetUserController(getUserUseCase)
	loginUserUseCase := application.NewLoginUseCase(ps)
	login_user_controller := controllers.NewLoginUserController(loginUserUseCase)
    RegisterUserRoutes(r, create_user_controller, get_user_controller, login_user_controller)
}