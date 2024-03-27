package service

import (
	"todo_list/repository"
	"todo_list/structs"
)

type TodoListService struct {
	repo repository.ToDoList
}

func NewTodoListService(repo repository.ToDoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) GetAll(userId int) ([]structs.ToDo, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, itemId int) (structs.ToDo, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoListService) Create(userId int, itemList structs.ToDo) error {
	return s.repo.Create(userId, itemList)
}

func (s *TodoListService) Delete(userId, itemId int) error { return s.repo.Delete(userId, itemId) }
func (s *TodoListService) Update(userId, itemId int, itemList structs.UpdateToDo) error {
	return s.repo.Update(userId, itemId, itemList)
}
