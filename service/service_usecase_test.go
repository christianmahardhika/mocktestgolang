package service

import(
  "testing"
  "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repository = &TestRepositoryMock{Mock: mock.Mock{}}
var testService = useCase{repo: repository}

func TestService_SaveTodoSuccess(t *testing.T) {
  // Mock Entity Todo
  todo := Todo{
    Title: "this is title",
  }

  // Mock Entity TodoDetail
  todoDetail1 := TodoDetail{
      ID: todo.ID,
      Item: "item 0",
    }
  todoDetail2 := TodoDetail{
      ID: todo.ID,
      Item: "item 1",
    }
  todoDetail3 := TodoDetail{
      ID: todo.ID,
      Item: "item 2",
    }
  repository.Mock.On("CreateTodo", &todo).Return(todo)
  repository.Mock.On("CreateTodoDetail", &todoDetail1).Return(todoDetail1)
  repository.Mock.On("CreateTodoDetail", &todoDetail2).Return(todoDetail2)
  repository.Mock.On("CreateTodoDetail", &todoDetail3).Return(todoDetail3)
  result, err := testService.SaveTodo(3)
  assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestService_FindTodoSuccess(t *testing.T) {
  // Mock Entity Todo
  todo := Todo{
    ID: 1,
    Title: "this is title",
  }

  // Mock Entity TodoDetail Array
  todoDetail := []*TodoDetail{
    {
      ID: 1,
    TodoID: todo.ID,
      Item: "item 1",
    },
  {
      ID: 2,
    TodoID: todo.ID,
      Item: "item 2",
    },
  {
      ID: 3,
    TodoID: todo.ID,
      Item: "item 3",
    },
  }

  var todoAll TodoAll
  todoAll.todo = todo
  todoAll.todoDetail = todoDetail
  
  repository.Mock.On("FindTodoDetailById", 1).Return(todoDetail)
  repository.Mock.On("FindTodoById", 1).Return(&todo)
  result, err := testService.GetTodoById(1)
  assert.Nil(t, err)
	assert.NotNil(t, result)
  assert.Equal(t, &todoAll, result, "the result should be equal")
}

func TestService_FindTodoFailed(t *testing.T) {
  repository.Mock.On("FindTodoDetailById", 5).Return(nil)
  repository.Mock.On("FindTodoById", 5).Return(nil)
  result, err := testService.GetTodoById(5)
  assert.Nil(t, result)
	assert.NotNil(t, err)
}