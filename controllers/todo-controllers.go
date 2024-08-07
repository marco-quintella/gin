package controllers

import (
	"fmt"
	"gin/config"
	"gin/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

var db = config.ConnectDB()

type todoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type todoResponse struct {
	todoRequest
	ID uint `json:"id"`
}

func CreateTodo(context *gin.Context) {
	var data todoRequest

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}
	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Create(todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	if err := db.Find(&todos); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "success",
		"data":    todos,
	})
}

func UpdateTodo(context *gin.Context) {
	var data todoRequest
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}

	todoById := db.Where("id = ?", idTodo).First(&todoResponse{})

	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
	}

	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

func DeleteTodo(context *gin.Context) {
	todo := models.Todo{}

	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	deleteTodo := db.Where("id = ?", idTodo).Unscoped().Delete(&todo)
	fmt.Println(deleteTodo)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idTodo,
	})
}
