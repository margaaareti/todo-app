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
	Create(userdId int, list todo.TodoList) (int, error) //->Создаем файл работы со списками todo_list_postgress.go
	GetAll(userdId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId int, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	CreateItem(listId int, input todo.TodoItem) (int, error)
	GetAllItems(userId, listId int) ([]todo.TodoItem, error)
	GetItemById(userId, itemId int) (todo.TodoItem, error)
	DeleteItem(userId, itemId int) error
	UpdateItem(userId, itemId int, input todo.UpdateItemInput) error
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
		TodoList:      NewTodoListPostgres(db), //-> также в сервисе
		TodoItem:      NewTodoItemPostgres(db),
	}
}
