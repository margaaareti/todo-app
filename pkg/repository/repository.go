//Объявляем заготовки интерфейсов для наших сущностей
package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/todo-app"
)

//Названия интерфейсов даются исходя из доменной зоны- участков бизнес логики приложения, за которые они отвечают

type Authorization interface {
	CreateUser(todo.User) (int, error)
	//получение пользователя из базы.Если пользователь есть- генерируем токен, в кот. записываем id пользователя
	//Если такого пользователя нет- ошибку
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

//Создаем структуру-сервис, собирающую все наши сервисы в одном месте
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

//Т.к. репозитории должны работать с БД, передадим объект sqlx.DB в качестве аргумента в конструктор
// ... и сразу же объявим конструктор
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostqres(db),
	}
}