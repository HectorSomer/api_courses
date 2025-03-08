package infraestructure

import (
	"event_driven/src/users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, createUserController *controllers.CreateUserController, getUserController *controllers.GetUserController, loginUserController *controllers.LoginUserController) {

    routes := r.Group("/v1/users")
	{
      routes.POST("/", createUserController.CreateUser);
	  routes.GET("/:id", getUserController.GetUser);
	  routes.POST("/login", loginUserController.LoginUser);
	}
}