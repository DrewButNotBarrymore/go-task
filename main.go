package main

import (
	"github.com/DrewButNotBarrymore/go-task/controllers"
	"github.com/DrewButNotBarrymore/go-task/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize gin
	r := gin.Default()

	// connect db
	models.ConnectDatabase()

	// routes
	r.GET("/tasks", controllers.FindTasks)
	r.GET("/tasks/user/:userid", controllers.FindUserTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.POST("/users", controllers.CreateUser)
	r.PATCH("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

	// run server
	r.Run(":8080")
}
