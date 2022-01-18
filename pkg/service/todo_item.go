package service

import (
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

type ToDoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewToDoItemService(repo repository.TodoItem, listRepo repository.TodoList) *ToDoItemService {
	return &ToDoItemService{repo: repo, listRepo: listRepo}
}

//1.Создаем item
func (s *ToDoItemService) CreateItem(userId, listId int, item todo.TodoItem) (int, error) {
	//Создаем item по id пользователя и списка
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		//лист не существует или не принадлежит пользователю
		return 0, err
	}
	return s.repo.CreateItem(listId, item)
}

//2.Получаем item
func (s *ToDoItemService) GetAllItems(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAllItems(userId, listId)
}
