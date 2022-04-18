package service

import "database/sql"

type Repository interface {
	CreateTodo(todo *Todo) error
	CreateTodoDetail(todoDetail *TodoDetail) error
	FindTodoById(id int) (*Todo, error)
	FindTodoDetailById(id int) ([]*TodoDetail, error)
	FindTodos() ([]*Todo, error)
	DeleteTodo(id int) error
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

// CreateTodo implements Repository
func (*repository) CreateTodo(todo *Todo) error {
	panic("unimplemented")
}

// CreateTodoDetail implements Repository
func (*repository) CreateTodoDetail(todoDetail *TodoDetail) error {
	panic("unimplemented")
}

// DeleteTodo implements Repository
func (*repository) DeleteTodo(id int) error {
	panic("unimplemented")
}

// FindTodoById implements Repository
func (*repository) FindTodoById(id int) (*Todo, error) {
	panic("unimplemented")
}

// FindTodoDetailById implements Repository
func (*repository) FindTodoDetailById(id int) ([]*TodoDetail, error) {
	panic("unimplemented")
}

// FindTodos implements Repository
func (*repository) FindTodos() ([]*Todo, error) {
	panic("unimplemented")
}
