package cli

import (
	"github.com/mr-doggy/doggy-sdk-go/doggy"
)

type TodoOp struct {
	api *doggy.APIClient
}

func NewTodoOp(api *doggy.APIClient) *TodoOp {
	return &TodoOp{
		api: api,
	}
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
