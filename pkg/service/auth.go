//реализовываем интерфейс
package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

const (
	salt = "sdfsdf324gdsfg3"
	//ключи подписи для генерации токена, использующиеся для его расшифровки
	signingKey = "gdfh34hjhhk34ghh3345%$^DF23"
	tokenTTL   = 12 * time.Hour
)

//Структура, имеющая стандартные поля claims + id пользователя
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	//получаем пользователя из базы данных
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	//Если такой юзер существует-- сгенерируем токен

	token := jwt.NewWithClaims(jwt.SigningMethodES256, &tokenClaims{
		jwt.StandardClaims{
			//Укажем значение на 12 часов больше текущего времени,т.е. токен перестанет быть валидным через 12 часов
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			//Время, когда токен был сгенерирован
			IssuedAt: time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

//Имплементируем логику создания пользователя

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}
