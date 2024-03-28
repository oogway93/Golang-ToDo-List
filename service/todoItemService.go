package service

import (
	"todo_list/repository"
	"todo_list/structs"
)

type ToDoItemService struct {
	repo     repository.ToDoItem
	listRepo repository.ToDoList
}

func NewToDoItemService(repo repository.ToDoItem, listRepo repository.ToDoList) *ToDoItemService {
	return &ToDoItemService{repo: repo, listRepo: listRepo}
}

func (s *ToDoItemService) GetAll(userId, listId int) ([]structs.ToDoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *ToDoItemService) Create(listId int, item structs.ToDoItem) error {
	return s.repo.Create(listId, item)
}
