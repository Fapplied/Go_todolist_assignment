package controllers

import (
	"net/http"

	"github.com/fapplied/Go_todolist_assignment/initializers"
	"github.com/fapplied/Go_todolist_assignment/models"
	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {

	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		return
	}


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
	if err := initializers.DB.First(&todo, id).First(&todo).Error; err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

	c.IndentedJSON(http.StatusOK, todo)

}

func UpdateTodoDone(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	if err := initializers.DB.First(&todo, id).First(&todo).Error; err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
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




