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

//2.Получение всех айтемов
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

//3.Получений айтемов по id
func (h *Handler) getItemByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return //--переходим в todo и ставим тег "binding:"required" для поля Title
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}

	var input todo.TodoItem
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	item, err := h.services.TodoItem.GetItemById(userId, itemId)

	c.JSON(http.StatusOK, item)

}

//4.Удаление айтемов
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return //--переходим в todo и ставим тег "binding:"required" для поля Title
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}

	var input todo.TodoItem
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.TodoItem.DeleteItem(userId, itemId)

	c.JSON(http.StatusOK, statusResponse{"ok"})

}

//5.Редактирование
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c) //Функция определена в middleware.go
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.TodoItem.UpdateItem(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, statusResponse{
		"Ok",
	})

}
