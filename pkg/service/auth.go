//реализовываем интерфейс
package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

//Структура, имеющая стандартные поля claims + id пользователя
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	//получаем пользователя из базы данных
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	//Если такой юзер существует-- сгенерируем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			//Укажем значение на 12 часов больше текущего времени,т.е. токен перестанет быть валидным через 12 часов
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			//Время, когда токен был сгенерирован
			IssuedAt: time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

//Парсим токен
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		//Проверка метода подписи токена: если это не HMAC- то ошибка, если да- то возвращаем ключ-подпись
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	//Ф-ия ParseWithClaims вовзвращает объект токена, в котором есть поле claims типа интерфейс
	//приведем его к нашей структуре и проверим, все ли хорошо
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims arent of type *tokenClaims")
	}
	//Если успешно распарсили токен, вернем значение id пользователя
	return claims.UserId, nil //-> Возвращаемся в middleware.go и вызываем метод
}

//Хэшируем пароль
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}
