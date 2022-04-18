package service

type UseCase interface {
	SaveTodo(todoAll *TodoAll) (string, error)
	GetTodoDetail(id int) (*TodoAll, error)
	GetTodos() ([]*Todo, error)
	DeleteTodo(id int) error
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{repo: repo}
}

type useCase struct {
	repo Repository
}

// DeleteTodo implements UseCase
func (*useCase) DeleteTodo(id int) error {
	panic("unimplemented")
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
func (uc *useCase) SaveTodo(todoAll *TodoAll) (string, error) {

	err := uc.repo.CreateTodo(&todoAll.todo)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(todoAll.todoDetail); i++ {
		todoDetail := todoAll.todoDetail[i]
		err := uc.repo.CreateTodoDetail(todoDetail)
		if err != nil {
			return "", err
		}
	}

	return "success", nil

}

// Menamplikan Todo List sesuai ID
func (uc *useCase) GetTodoDetail(id int) (*TodoAll, error) {
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
