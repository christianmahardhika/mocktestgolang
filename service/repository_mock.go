package service

import (
	"context"
	"errors"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestRepositoryMock struct {
	Mock mock.Mock
}

func (repository *TestRepositoryMock) CreateTodo(ctx context.Context, todo *Todo) (*primitive.ObjectID, error) {
	arguments := repository.Mock.Called(todo)
	if arguments.Get(0) == nil {
		return nil, errors.New("error CreateTodo")
	} else {
		return &todo.ID, nil
	}
}

func (repository *TestRepositoryMock) CreateTodoDetail(ctx context.Context, todo *TodoDetail) (*primitive.ObjectID, error) {
	arguments := repository.Mock.Called(todo)
	if arguments.Get(0) == nil {
		return nil, errors.New("error CreateTodoDetail")
	} else {
		return &todo.ID, nil
	}
}

func (repository *TestRepositoryMock) FindTodoById(id string) (*Todo, error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("error FindTodoById")
	} else {
		todo := arguments.Get(0).(*Todo)
		return todo, nil
	}
}

func (repository *TestRepositoryMock) FindTodoDetailById(id string) ([]TodoDetail, error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("error FindTodoDetailById")
	} else {
		todo := arguments.Get(0).([]TodoDetail)
		return todo, nil
	}
}

func (repository *TestRepositoryMock) FindTodos() ([]*Todo, error) {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("error FindTodos")
	} else {
		todo := arguments.Get(0).([]*Todo)
		return todo, nil
	}
}

func (repository *TestRepositoryMock) DeleteTodo(id string) error {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return errors.New("error DeleteTodo")
	} else {
		return nil
	}
}

func (repository *TestRepositoryMock) DeleteTodoDetail(id string) error {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return errors.New("error DeleteTodo")
	} else {
		return nil
	}
}
