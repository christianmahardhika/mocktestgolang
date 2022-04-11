package service

type Repository interface {
  CreateTodo(todo *Todo) error 
  CreateTodoDetail(todoDetail *TodoDetail) error
  FindTodoById(id int) (*Todo, error)
  FindTodoDetailById(id int) ([]*TodoDetail, error)
}