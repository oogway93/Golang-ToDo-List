package service

import (
	"todo_list/repository"
	"todo_list/structs"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) GetAll() ([]structs.ToDo, error) {
	return s.repo.GetAll()
}

func (s *TodoListService) GetById(itemId int) (structs.ToDo, error) {
	return s.repo.GetById(itemId)
}

func (s *TodoListService) Create(itemList structs.ToDo) error {
	return s.repo.Create(itemList)
}

func (s *TodoListService) Delete(itemId int) error { return s.repo.Delete(itemId) }
func (s *TodoListService) Update(itemId int, itemList structs.UpdateToDo) error {
	return s.repo.Update(itemId, itemList)
}
