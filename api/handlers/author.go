package handlers

import (
	"bootcamp/article/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAuthor(c *gin.Context) {
	var entity models.PersonCreateModel
	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(entity)

	err = h.strg.Author().Create(entity)

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, models.SuccessResponse{
		Message: "article has been created",
		Data:    "ok",
	})
}

func (h *Handler) UpdateAuthor(c *gin.Context) {
	var entity models.PersonUpdateModel
	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(entity)

	 err = h.strg.Author().Update(entity)

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, models.SuccessResponse{
		Message: "article has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetAuthorsList(c *gin.Context) {
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

	resp, err := h.strg.Author().GetList(models.Query{Offset: offset, Limit: limit, Search: c.Query("search")})

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}

func (h *Handler) GetAuthorsByIDHandler(c *gin.Context) {

	id := c.Param("id")
	author, err := h.strg.Author().GetByID(id)
	if err!= nil {
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

func (h *Handler) DeleteAuthor(c *gin.Context) {

	id := c.Param("id")
	author, err := h.strg.Author().Delete(id)
	if err!= nil {
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

