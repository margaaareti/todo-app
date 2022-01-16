//Эндпоинты для работы со списками
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app"
)

//Создаем списки
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c) //Функция определена в middleware.go
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return //--переходим в todo и ставим тег "binding:"required" для поля Title
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Возвращаем тело ответа при успешном запросе, в кот. возвращаем id созданного списка
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	}) //->Реализуем получение всех списков
}

//Для response getAllLists используем допю структуру
type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

//Получение всех списков
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c) //Функция определена в middleware.go
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return //--переходим в todo и ставим тег "binding:"required" для поля Title
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return //-Переходим в service для указания в интерфейсе данного метода
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	}) //->Реализуем получение списков по id списка
}

//Получение списков по id списка
func (h *Handler) getListByID(c *gin.Context) {
	userId, err := getUserId(c) //Функция определена в middleware.go
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return //--переходим в todo и ставим тег "binding:"required" для поля Title
	}

	//Получение параметра id списка из пути запроса
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return //-Переходим в service для указания в интерфейсе данного метода
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
