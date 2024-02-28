package api

import (
	"bootcamp/article/api/docs"
	_ "bootcamp/article/api/docs"
	"bootcamp/article/api/handlers"
	"bootcamp/article/config"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @description This is a sample article demo.
// @termsOfService https://udevs.io
func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.Version
	docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.HTTPPort
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.POST("/createarticle", h.CreateArticle)
	r.POST("/createauthor", h.CreateAuthor)
	r.GET("/articlelist", h.GetArticleList)
	r.GET("/authorlist", h.GetAuthorsList)
	r.GET("/authors/:id", h.GetAuthorsByIDHandler)
	r.GET("/articles/:id", h.GetArticleByIDHandler)
	r.DELETE("/deleteauthor/:id", h.DeleteAuthor)
	r.DELETE("/deletearticle/:id", h.DeleteArticle)
	r.PUT("/updateauthor/", h.UpdateAuthor)
	r.PUT("/updatearticle/", h.UpdateArticle)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
