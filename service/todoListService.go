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

func (s *TodoListService) GetById(userId, listId int) (structs.ToDo, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Create(userId int, itemList structs.ToDo) error {
	return s.repo.Create(userId, itemList)
}

func (s *TodoListService) Delete(userId, listId int) error { return s.repo.Delete(userId, listId) }
func (s *TodoListService) Update(userId, listId int, itemList structs.UpdateToDo) error {
	return s.repo.Update(userId, listId, itemList)
}
