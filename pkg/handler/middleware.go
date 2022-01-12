//Middlewaere парсит токены из запроса и предоставляет доступ
//к группе эйндпоинтов API
package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

//Создаем метод для обработчика
func (h *Handler) userIdentity(c *gin.Context) {
	//Получаем значение из хэдера авторизации
	//и валидируем его, что он не пустой
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusAlreadyReported, "empty auth header")
		return
	}

	//Укажем разделить строку по пробелам
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	} //Далее переходим в service.go и добавим метод для парса токена

	// Вызываем функцию ParseToken
	//Если операция успешна- запишем значение id в контекст
	//Для того, чтобы иметь доступ к id пользователя в последующих обработчиках
	//Которые вызываются после данной прослойки
	userID, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userID) //-> Переходим в handler.go и задаем метод userIdentity в качестве обработчика
}
