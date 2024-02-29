package api

import (
	"todo/api/docs"
	_ "todo/api/docs"
	"todo/api/handlers"
	"todo/config"

	"github.com/gin-gonic/gin"
	// swagger embed files
	// gin-swagger middleware
)

// @description This is a sample article demo.
// @termsOfService https://udevs.io
func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.Version
	docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.HTTPPort
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.PUT("/updateuser/", h.UpdateUser)
	r.POST("/createuser", h.CreateUser)
	r.GET("/userlist", h.GetUsersList)
	r.GET("/users/:id", h.GetUsersByIDHandler)
	r.DELETE("/deleteuser/:id", h.DeleteUser)
	
	r.GET("/todolist", h.GetTodoList)
	r.GET("/todos/:id", h.GetTodoByIDHandler)
	r.DELETE("/deletetodo/:id", h.DeleteTodo)
	r.POST("/createtodo", h.CreateTodo)
	r.PUT("/updatetodo/", h.UpdateTodo)

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
