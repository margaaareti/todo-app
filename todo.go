package todo

import "errors"

//Добавляем теги db для того,чтобы делать выборку из базы
type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"` //переходим в service и реализуем метод
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int `json:"id"`
	UserId int `json:"title"`
	ListId int `json:"description"`
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done"`
}

type ListenItem struct {
	Id     int
	UserId int
	ListId int
}

//Структура для изменения списков
type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"` //->вернемся к list.go
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"` //->вернемся к list.go
	Done        *bool   `json:"done"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
