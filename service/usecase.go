package service

import(
  "strconv"
)

type UseCase interface {
  SaveTodo(numberOfItems int) (string, error)
}

func NewUseCase(repo Repository) UseCase {
  return &useCase{repo: repo}
}

type useCase struct {
  repo Repository
}

// Menyimpan Todo List berdasarkan jumlah detail item
func (uc *useCase) SaveTodo(numberOfItems int) (string, error){
  todo := Todo{
    Title: "this is title",
  }
  err := uc.repo.CreateTodo(&todo)
  if err != nil {
    return "", err
  }

  for i := 0; i < numberOfItems; i++ {
    todoDetail := TodoDetail{
      ID: todo.ID,
      Item: "item "+ strconv.Itoa(i),
    }
    err := uc.repo.CreateTodoDetail(&todoDetail)
    if err != nil {
    return "", err
  }
  }

  return "success", nil
  
}

// Menamplikan Todo List sesuai ID
func (uc *useCase) GetTodoById(id int) (*TodoAll, error){
  res, err := uc.repo.FindTodoDetailById(id)
  if err != nil {
    return nil, err
  }
  resTodo, err := uc.repo.FindTodoById(res[0].ID)
  if err != nil {
    return nil, err
  }
  var todoResult TodoAll
  todoResult.todoDetail = res
  todoResult.todo = *resTodo
  


  return &todoResult, nil
  
}