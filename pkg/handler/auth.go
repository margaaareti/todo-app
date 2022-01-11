//Создаем обработчики sighUp и sighIn и присваиваем их маршрутам
//Объявляем эндпоинты для регистрации и авторизации
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app"
)

//Регистрация
func (h *Handler) sighUp(c *gin.Context) {
	//Структура, куда будут записываться данные из json от пользователей
	var input todo.User

	//Принимает ссылку на объект, в который мы хотим распарсить тело json
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//После того, как мы распарсили тело запроса и валидировали его
	//передаем эти данные в servise, где реализована бизнес-логика регистрации

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		//500- код ошибки о том, что пользователь ввел некорректные данные
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

//Аутентификация

//Структура для получения логина и пароля в запросе:
type signInInput struct {
	Username string `json:"username" `
	Password string `json:"password" `
}

func (h *Handler) sighIn(c *gin.Context) {

	//Структура, куда будут записываться данные из json от пользователей
	var input signInInput

	//Принимает ссылку на объект, в который мы хотим распарсить тело json
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Генерируем токен, передавая логин и пароль
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		//500- код ошибки о том, что пользователь ввел некорректные данные
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
