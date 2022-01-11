//Объявляем заготовки интерфейсов для наших сущностей
package service

import (
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

//Названия интерфейсов даются исходя из доменной зоны- участков бизнес логики приложения, за которые они отвечают

type Authorization interface {
	//Принимает структуру User и возвращает id созданного в базе пользователя
	CreateUser(user todo.User) (int, error)
	//Принимаем лгин и пароль, а возвращаем- токен и ошибку
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

//Создаем структуру-сервис, собирающую все наши сервисы в одном месте
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// ... и сразу же объявим конструктор
//Сервисы будут обращаться к БД, потому объявляем указатель на структуру repository в качестве аргумента конструктора
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
