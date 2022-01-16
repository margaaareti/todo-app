//Реализовываем методы работы со списками
package service

import (
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

//При создании списка мы будем передавать данные в репозиторий
//поэтому в сервисе(здесь) мы возвращаем лишь аутентичный метод репозитория
func (s *TodoListService) Create(userdId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userdId, list) //-> переходим в репозиторий и создаем метод Create
}

func (s *TodoListService) GetAll(userdId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userdId) //-> переходим в репозиторий и создаем метод GetAll
}

func (s *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId) //-> переходим в репозиторий и создаем метод GetById
}
