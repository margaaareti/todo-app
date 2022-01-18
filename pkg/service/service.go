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
	//Принимаем логин и пароль, а возвращаем- токен и ошибку
	GenerateToken(username, password string) (string, error)
	//Метод длЯ парсинга токена:принимает токен, возвращает id пользователя при успешном парсинге
	ParseToken(token string) (int, error) //-> Переходим в auth.go и имплементируем эту логику
}

type TodoList interface {
	//для создания списков принимаем id пользователя и структуру списка
	//возвращаем id и ошибку
	Create(userId int, list todo.TodoList) (int, error) //->Создаем файл todo.go, где реализуем все методы todo
	//Для получения всех списков принимаем id пользователя и возвращаем слайс со списками
	GetAll(userId int) ([]todo.TodoList, error)
	//Для получения списка по id нам нужно id самого списка и id пользователя
	GetById(userId, listId int) (todo.TodoList, error)
	//Для удаления нужно получать id списка и юзера
	Delete(userId, listId int) error
	//id юзера, списка, и форму
	Update(userId int, id int, input todo.UpdateListInput) error
}

type TodoItem interface {
	CreateItem(userId, listId int, input todo.TodoItem) (int, error)
	GetAllItems(userId, listId int) ([]todo.TodoItem, error)
	GetItemById(userId, itemId int) (todo.TodoItem, error)
	DeleteItem(userId, itemId int) error
	UpdateItem(userId, itemId int, input todo.UpdateItemInput) error
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
		TodoList:      NewTodoListService(repos.TodoList), //->Переходим в обработчик list.go и вызовем метод создаю списков
		TodoItem:      NewToDoItemService(repos.TodoItem, repos.TodoList),
	}
}
