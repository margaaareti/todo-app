package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/todo-app"
)

//Создаем структуру репозитория
type TodoListPostgres struct {
	db *sqlx.DB
}

//И ее конструктор
func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

//Имплементируем метод создания списков используя транзакции БД
func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	//Запрос на создание записи в таблицe todo_list возвращая id нового списка
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	//Делаем вставку в таблицу users_list, в которой свяжем id пользователя и id нового списка
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	//Exec- метод для простого чтения, без возврата значений
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		//Rollback- откатывает все изменения в бд до начала транзакций
		tx.Rollback()
		return 0, err
	}

	//tx.Commit()- вызывает изменения в базе данных и заканчивает транзакции
	return id, tx.Commit()
	//=> Переходим в repository и реализуем конструктор NewTodoListPostgress
}

//Имплементируем метод возвращения всех списков
func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id=$1", todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err //-> Возвращаемся к обработчику list.go чтобы добавить тело ответа
}

//Имплементируем метод возвращения списков по id
func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err //-> Возвращаемся к обработчику list.go чтобы добавить тело ответа
}
