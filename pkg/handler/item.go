package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app"
)

//1.Создаем item
func (h *Handler) createItem(c *gin.Context) {
	//Получаем id пользователя и списка, в котором создается элемент
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return //--переходим в todo и ставим тег "binding:"required" для поля Title
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}

	var input todo.TodoItem
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.TodoItem.CreateItem(userId, listId, input)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

//Получение всех айтемов
func (h *Handler) getAllItems(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return //--переходим в todo и ставим тег "binding:"required" для поля Title
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAllItems(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)

}

func (h *Handler) getItemByID(c *gin.Context) {

}
func (h *Handler) updateItem(c *gin.Context) {

}
func (h *Handler) deleteItem(c *gin.Context) {

}
