package main

import (
	"github.com/fapplied/Go_todolist_assignment/controllers"
	"github.com/fapplied/Go_todolist_assignment/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectTodb()
}

func main() {
	r := gin.Default()

	r.POST("/todos", controllers.CreateTodo)
	r.GET("/todos", controllers.GetAllTodos)
	r.GET("/todos/:id", controllers.GetTodoById)
	r.PATCH("/todos/:id/done", controllers.UpdateTodoDone)
	r.DELETE("/todos/:id", controllers.DeleteTodoById)
	r.Run()
}
