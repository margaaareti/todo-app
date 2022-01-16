package todo

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
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListenItem struct {
	Id     int
	UserId int
	ListId int
}
