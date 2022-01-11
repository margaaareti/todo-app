//опишем структуру, который имплементирует интерфейс репозитория
//и работает с базой NewPostgresDB
package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/todo-app"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostqres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

//Пишем запрос
func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name,username,password_hash) values($1,$2,$3) RETURNING id", userTable)
	//row хранит инф.о возвращаемой строке из базы
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	//Scan записывает значение из БД в переменную передавая ее по ссылке
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

//Метод для получения пользователя по его логину и паролю
func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash=$2", userTable)
	//Get записывает в структуру user результат выборки
	err := r.db.Get(&user, query, username, password)

	//Возвращаем пользователя
	return user, err
}
