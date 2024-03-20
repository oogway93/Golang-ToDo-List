package service

import "todo_list/repository"

type TodoList interface {
}

type Service struct {
	TodoList
}

func NewService(repository *repository.Repository) *Service {
	return new(Service)
}
