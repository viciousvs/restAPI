package service

import (
	restapi "github.com/viciousvs/restAPI"
	"github.com/viciousvs/restAPI/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user restapi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list restapi.TodoList) (int, error)
	GetAll(userId int) ([]restapi.TodoList, error)
	GetById(userId, listId int) (restapi.TodoList, error)
	Delete(userId, listId int) error
	// Update(userId, listId int, input restapi.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item restapi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]restapi.TodoItem, error)
	GetById(userId, itemId int) (restapi.TodoItem, error)
	Delete(userId, itemId int) error
	// Update(userId, itemId int, input restapi.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

// func NewService(repos *repository.Repository) *Service {
// 	return &Service{
// 		Authorization: NewAuthService(repos.Authorization),
// 		TodoList:      NewTodoListService(repos.TodoList),
// 		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
// 	}
// }

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
