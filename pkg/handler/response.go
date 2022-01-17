//определяем функцию для стандартной обработки ошибок
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Структура для ответа на успешное удаление листа
type statusResponse struct {
	Status string `json:"status" `
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, message)
}
