package service

type UseCase interface {
	SaveTodo(todoAll *TodoAll) (*TodoAll, error)
	GetTodoDetail(id string) (*TodoAll, error)
	GetTodos() ([]*Todo, error)
	DeleteTodo(id string) error
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{repo: repo}
}

type useCase struct {
	repo Repository
}

// DeleteTodo implements UseCase
func (uc *useCase) DeleteTodo(id string) error {
	err := uc.repo.DeleteTodoDetail(id)
	if err != nil {
		return err
	}
	return uc.repo.DeleteTodo(id)
}

// GetTodos implements UseCase
func (uc *useCase) GetTodos() ([]*Todo, error) {
	res, err := uc.repo.FindTodos()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Menyimpan Todo List berdasarkan jumlah detail item
func (uc *useCase) SaveTodo(todoAll *TodoAll) (*TodoAll, error) {
	res, err := uc.repo.CreateTodo(&todoAll.Todo)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(todoAll.TodoDetail); i++ {
		todoDetail := todoAll.TodoDetail[i]
		todoDetail.TodoID = *res
		_, err := uc.repo.CreateTodoDetail(&todoDetail)
		if err != nil {
			return nil, err
		}
	}

	return todoAll, nil

}

// Menamplikan Todo List sesuai ID
func (uc *useCase) GetTodoDetail(id string) (*TodoAll, error) {
	res, err := uc.repo.FindTodoDetailById(id)
	if err != nil {
		return nil, err
	}
	resTodo, err := uc.repo.FindTodoById(id)
	if err != nil {
		return nil, err
	}
	var todoResult TodoAll
	todoResult.TodoDetail = res
	todoResult.Todo = *resTodo

	return &todoResult, nil

}
