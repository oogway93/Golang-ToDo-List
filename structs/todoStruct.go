package structs

type ToDo struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UpdateToDo struct {
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
}
