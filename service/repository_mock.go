package service

import(
  "errors"
  "github.com/stretchr/testify/mock"
)
type TestRepositoryMock struct {
	Mock mock.Mock
}

func (repository *TestRepositoryMock) CreateTodo(todo *Todo) error {
	arguments := repository.Mock.Called(todo)
	if arguments.Get(0) == nil {
		return errors.New("Error CreateTodo")
	} else {
		return nil
	}
}

func (repository *TestRepositoryMock) CreateTodoDetail(todo *TodoDetail) error {
	arguments := repository.Mock.Called(todo)
	if arguments.Get(0) == nil {
		return errors.New("Error CreateTodoDetail")
	} else {
		return nil
	}
}

func (repository *TestRepositoryMock) FindTodoById(id int) (*Todo, error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("Error")
	} else {
    todo := arguments.Get(0).(*Todo)
		return todo, nil
	}
}

func (repository *TestRepositoryMock) FindTodoDetailById(id int) ([]*TodoDetail, error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("Error")
	} else {
		todo := arguments.Get(0).([]*TodoDetail)
		return todo, nil
	}
}