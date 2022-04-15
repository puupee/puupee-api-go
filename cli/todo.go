package cli

type TodoOp struct {
}

func NewTodoOp() *TodoOp {
	return &TodoOp{}
}

func (op *TodoOp) Create() error {
	return nil
}

func (op *TodoOp) Update() error {
	return nil
}

func (op *TodoOp) Delete() error {
	return nil
}

func (op *TodoOp) List() error {
	return nil
}
