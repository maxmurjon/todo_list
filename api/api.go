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


	// r.POST("/login",h.Login)

	// ! User endpoints
	r.Use(customCORSMiddleware())
	r.PUT("/updateuser/", h.UpdateUser)
	r.POST("/createuser", h.CreateUser)
	r.GET("/userlist", h.GetUsersList)
	r.GET("/users/:id", h.GetUsersByIDHandler)
	r.DELETE("/deleteuser/:id", h.DeleteUser)


	// ! Task endpoints 
	r.GET("/tasklist", h.GetTaskList)
	r.GET("/task/:id", h.GetTaskByIDHandler)
	r.DELETE("/deletetask/:id", h.DeleteTask)
	r.POST("/createtask", h.CreateTask)
	r.PUT("/updatetask/", h.UpdateTask)

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}


func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
