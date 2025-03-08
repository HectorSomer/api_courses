package main

import (
    "event_driven/src/config"
    "event_driven/src/courses/infraestructure/dependencies"
    "event_driven/src/users/infraestructure"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    config.InitRabbitMQConnection()
    r := gin.Default()
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // Permitir todos los orígenes
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        AllowCredentials: false, // Necesario para usar "*"
    }))
    
    
    infraestructure.InitUsers(r)
    dependencies.InitCourses(r)
    if err := r.Run(); err != nil {
        panic(err)
    }
}
