package structs

type ToDo struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}
type ToDoAdd struct {
	Title       string `json:"title" db:"title" binding:"required" extensions:"x-order=0"`
	Description string `json:"description" db:"description" extensions:"x-order=1"`
}

type UpdateToDo struct {
	Title       *string `json:"title" db:"title" extensions:"x-order=0"`
	Description *string `json:"description" db:"description" extensions:"x-order=1"`
}

type ToDoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `default:"false" json:"done" db:"done"`
}

type ToDoItemAdd struct {
	Title       string `json:"title" db:"title" binding:"required" extensions:"x-order=0"`
	Description string `json:"description" db:"description" extensions:"x-order=1"`
	Done        bool   `default:"false" json:"done" db:"done" extensions:"x-order=2"`
}

type ListToDoItems struct {
	Id         int
	ToDoId     int
	ToDoItemId int
}
