package repository

type TodoList interface {
}

type Repository struct {
	TodoList
}

func NewRepository() *Repository {
	return new(Repository)
}
