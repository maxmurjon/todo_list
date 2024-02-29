package handlers

import (
	"fmt"
	"net/http"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var entity models.UserCreateModel
	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}
	uuid, err := h.strg.User().Create(entity)

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}
	var data models.User
	data, err = h.strg.User().GetByID(uuid)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, models.SuccessResponse{
		Message: "user has been created",
		Data:    data,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var entity models.UserUpdateModel
	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(entity)

	err = h.strg.User().Update(entity)

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, models.SuccessResponse{
		Message: "user has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetUsersList(c *gin.Context) {
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

	resp, err := h.strg.User().GetList(models.Query{Offset: offset, Limit: limit, Search: c.Query("search")})

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}

func (h *Handler) GetUsersByIDHandler(c *gin.Context) {

	id := c.Param("id")
	User, err := h.strg.User().GetByID(id)
	fmt.Println("EWeeeeeeeeeeeeeeeeeeeeeee ", User)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    User,
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")
	User, err := h.strg.User().Delete(id)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    User,
	})
}
