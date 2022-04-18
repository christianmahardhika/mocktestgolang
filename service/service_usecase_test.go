package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo = &TestRepositoryMock{Mock: mock.Mock{}}
var testService = useCase{repo: repo}

func TestService_SaveTodoSuccess(t *testing.T) {
	// Mock Entity Todo
	todo := Todo{
		Title: "this is title",
	}

	// Mock Entity TodoDetail
	todoDetail1 := TodoDetail{
		ID:   todo.ID,
		Item: "item 0",
	}
	todoDetail2 := TodoDetail{
		ID:   todo.ID,
		Item: "item 1",
	}
	todoDetail3 := TodoDetail{
		ID:   todo.ID,
		Item: "item 2",
	}

	todoAll := TodoAll{
		todo: todo,
		todoDetail: []*TodoDetail{
			{
				ID:     1,
				TodoID: todo.ID,
				Item:   "item 1",
			},
			{
				ID:     2,
				TodoID: todo.ID,
				Item:   "item 2",
			},
			{
				ID:     3,
				TodoID: todo.ID,
				Item:   "item 3",
			},
		},
	}

	repo.Mock.On("CreateTodo", &todo).Return(todo)
	repo.Mock.On("CreateTodoDetail", &todoDetail1).Return(todoDetail1)
	repo.Mock.On("CreateTodoDetail", &todoDetail2).Return(todoDetail2)
	repo.Mock.On("CreateTodoDetail", &todoDetail3).Return(todoDetail3)
	result, err := testService.SaveTodo(&todoAll)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestService_FindTodoSuccess(t *testing.T) {
	// Mock Entity Todo
	todo := Todo{
		ID:    1,
		Title: "this is title",
	}

	// Mock Entity TodoDetail Array
	todoDetail := []*TodoDetail{
		{
			ID:     1,
			TodoID: todo.ID,
			Item:   "item 1",
		},
		{
			ID:     2,
			TodoID: todo.ID,
			Item:   "item 2",
		},
		{
			ID:     3,
			TodoID: todo.ID,
			Item:   "item 3",
		},
	}

	var todoAll TodoAll
	todoAll.todo = todo
	todoAll.todoDetail = todoDetail

	repo.Mock.On("FindTodoDetailById", 1).Return(todoDetail)
	repo.Mock.On("FindTodoById", 1).Return(&todo)
	result, err := testService.GetTodoDetail(1)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, &todoAll, result, "the result should be equal")
}

func TestService_FindTodoFailed(t *testing.T) {
	repo.Mock.On("FindTodoDetailById", 5).Return(nil)
	repo.Mock.On("FindTodoById", 5).Return(nil)
	result, err := testService.GetTodoDetail(5)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestService_FindTodosSuccess(t *testing.T) {
	// Mock Entity Todo
	todo := []*Todo{
		{
			ID:    1,
			Title: "this is title",
		},
		{
			ID:    2,
			Title: "this is title2",
		},
	}

	repo.Mock.On("FindTodos").Return(todo)
	result, err := testService.GetTodos()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	// assert.Equal(t, &todoAll, result, "the result should be equal")
}

func TestService_DeleteTodoSuccess(t *testing.T) {

	repo.Mock.On("DeleteTodo", 1).Return(nil)
	result, err := testService.GetTodos()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	// assert.Equal(t, &todoAll, result, "the result should be equal")
}
