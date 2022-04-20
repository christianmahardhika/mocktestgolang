package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var repo = &TestRepositoryMock{Mock: mock.Mock{}}
var testService = useCase{repo: repo}

func TestService_SaveTodoSuccess(t *testing.T) {
	id1, _ := primitive.ObjectIDFromHex(uuid.New().String())
	id2, _ := primitive.ObjectIDFromHex(uuid.New().String())
	id3, _ := primitive.ObjectIDFromHex(uuid.New().String())
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
		Todo: todo,
		TodoDetail: []TodoDetail{
			{
				ID:     id1,
				TodoID: todo.ID,
				Item:   "item 1",
			},
			{
				ID:     id2,
				TodoID: todo.ID,
				Item:   "item 2",
			},
			{
				ID:     id3,
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
	id, _ := primitive.ObjectIDFromHex(uuid.New().String())
	id1, _ := primitive.ObjectIDFromHex(uuid.New().String())
	id2, _ := primitive.ObjectIDFromHex(uuid.New().String())
	id3, _ := primitive.ObjectIDFromHex(uuid.New().String())
	// Mock Entity Todo
	todo := Todo{
		ID:    id,
		Title: "this is title",
	}

	// Mock Entity TodoDetail Array
	todoDetail := []TodoDetail{
		{
			ID:     id1,
			TodoID: todo.ID,
			Item:   "item 1",
		},
		{
			ID:     id2,
			TodoID: todo.ID,
			Item:   "item 2",
		},
		{
			ID:     id3,
			TodoID: todo.ID,
			Item:   "item 3",
		},
	}

	var todoAll TodoAll
	todoAll.Todo = todo
	todoAll.TodoDetail = todoDetail

	repo.Mock.On("FindTodoDetailById", id1).Return(todoDetail)
	repo.Mock.On("FindTodoById", id).Return(&todo)
	result, err := testService.GetTodoDetail(uuid.New().String())
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, &todoAll, result, "the result should be equal")
}

func TestService_FindTodoFailed(t *testing.T) {
	repo.Mock.On("FindTodoDetailById", uuid.New().String()).Return(nil)
	repo.Mock.On("FindTodoById", uuid.New().String()).Return(nil)
	result, err := testService.GetTodoDetail(uuid.New().String())
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestService_FindTodosSuccess(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex(uuid.New().String())
	id1, _ := primitive.ObjectIDFromHex(uuid.New().String())
	// Mock Entity Todo
	todo := []*Todo{
		{
			ID:    id,
			Title: "this is title",
		},
		{
			ID:    id1,
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

	repo.Mock.On("DeleteTodo", uuid.New().String()).Return(nil)
	result, err := testService.GetTodos()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	// assert.Equal(t, &todoAll, result, "the result should be equal")
}
