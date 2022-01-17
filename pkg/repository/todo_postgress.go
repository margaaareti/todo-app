package repository

import (
	"fmt"
	"strings"

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

//1.Имплементируем метод создания списков используя транзакции БД
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

//2.Имплементируем метод возвращения всех списков
func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id=$1", todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err //-> Возвращаемся к обработчику list.go чтобы добавить тело ответа
}

//3.Имплементируем метод возвращения списков по id
func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err //-> Возвращаемся к обработчику list.go чтобы добавить тело ответа
}

//4.Имплементируем метод удаления списков по id
func (r *TodoListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2", todoListsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err //-> Возвращаемся к обработчику list.go чтобы добавить тело ответа
}

//5.Имплементируем метод изменения списков по id
//Реализация обновления
func (r *TodoListPostgres) Update(userId, listId int, input todo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	//Выполняем проверку полей:если они не nil- добавляем в наши слайсы элементы для
	//формирования запросов в базу с их обновлением
	if input.Title != nil {
		//Присвоение полю Title  и значение для плейсхолдера
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		//Присвоение полю Title  и значение для плейсхолдера
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	//объявим перем.,в кот. соед. элементы слайса строк в одну строку
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)

	//Добавим еще 2 элемента в слайс аргументов: id пользователя и списков
	args = append(args, listId, userId)

	//Выполним запрос, в кот. передадим запрос и список аргументов
	_, err := r.db.Exec(query, args...)
	return err

}
