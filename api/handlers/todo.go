package handlers

import (
	"fmt"
	"net/http"
	"todo/models"

	"github.com/gin-gonic/gin"
)

// CreateTodo godoc
// @tags Todo
// @ID create-article-handler
// @Summary Create Article
// @Description Create Article By Given Info and Author ID
// @Param data body models.ArticleCreateModel true "Article Body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SuccessResponse{data=string}
// @Failure default {object} models.DefaultError
// @Router /articles [POST]
func (h *Handler) CreateTodo(c *gin.Context) {
	var entity models.TodoCreateModel

	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(entity)

	uuid, err := h.strg.Todo().Create(entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	fmt.Println(err)

	resp, err := h.strg.Todo().GetByID(uuid)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, models.SuccessResponse{
		Message: "Todo has been created",
		Data:    resp,
	})
}

// GetTodoList godoc
// @tags Todos
// @ID get-all-handler
// @Summary List Todos
// @Description jimgina yaxshilab ishlat, yaxshi bola bol
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param search query string false "search string"
// @Param something query string false "something"
// @Accept  json
// @Produce  json
// @Success 200 {array} models.TodoListItem
// @Failure default {object} models.DefaultError
// @Router /Todos [get]
func (h *Handler) GetTodoList(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	resp, err := h.strg.Todo().GetList(models.Query{Offset: offset, Limit: limit, Search: c.Query("search")})

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}

func (h *Handler) GetTodoByIDHandler(c *gin.Context) {

	id := c.Param("id")
	author, err := h.strg.Todo().GetByID(id)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    author,
	})
}

func (h *Handler) DeleteTodo(c *gin.Context) {

	id := c.Param("id")
	author, err := h.strg.Todo().Delete(id)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    author,
	})
}

func (h *Handler) UpdateTodo(c *gin.Context) {
	var entity models.TodoUpdateModel
	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(entity)

	err = h.strg.Todo().Update(entity)

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, models.SuccessResponse{
		Message: "Todo has been updated",
		Data:    "ok",
	})
}
