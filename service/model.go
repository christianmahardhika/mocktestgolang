package service

type Todo struct {
  ID int
  Title string
}

type TodoDetail struct {
  ID int
  TodoID int
  Item string
}

type TodoAll struct {
  todo Todo
  todoDetail []*TodoDetail
}