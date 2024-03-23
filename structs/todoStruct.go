package structs

type ToDo struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}
type ToDoAdd struct {
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UpdateToDo struct {
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
}

type ToDoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListToDoItems struct {
	Id         int
	ToDoId     int
	ToDoItemId int
}
