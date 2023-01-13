package controllers

import (
	"net/http"

	"github.com/fapplied/Go_todolist_assignment/initializers"
	"github.com/fapplied/Go_todolist_assignment/models"
	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {

	var body struct {
		Body    string
		Title   string
		DueDate string
	}

	c.Bind(&body)

	todo := models.Todo{Title: body.Title, Body: body.Body, DueDate: body.DueDate}

	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad"})
		return
	}

	c.IndentedJSON(http.StatusCreated, todo)

}

func GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	initializers.DB.Find(&todos)

	c.IndentedJSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	initializers.DB.First(&todo, id)
	if todo.ID == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo with that id not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)

}

func UpdateTodoDone(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	if err := initializers.DB.Where("id = ?", id).First(&todo).Error; err != nil {
        c.IndentedJSON(404, gin.H{"error": "Todo not found"})
        return
    }


	initializers.DB.Model(&todo).Select("Done").Updates(models.Todo{Done:  !todo.Done})

	c.IndentedJSON(http.StatusOK, todo)

}

func DeleteTodoById(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Todo{}, id)

	c.JSON(http.StatusNoContent, gin.H{"message" : "item deleted"})
}




