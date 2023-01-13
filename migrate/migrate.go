package main

import (
	"github.com/fapplied/Go_todolist_assignment/initializers"
	"github.com/fapplied/Go_todolist_assignment/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectTodb()
	
}

func main() {
	initializers.DB.AutoMigrate(&models.Todo{})
}