package routes

import (
	"gin/controllers"
	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	route.POST("/todo", controllers.CreateTodo)
	route.GET("/todo", controllers.GetAllTodos)
	route.PUT("/todo", controllers.UpdateTodo)
	route.DELETE("/todo", controllers.DeleteTodo)

	err := route.Run()
	if err != nil {
		panic("Error starting routes")
		return
	}
}
