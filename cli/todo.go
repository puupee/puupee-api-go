package cli

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
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

func (op *TodoOp) List(pageIndex, pageSize int32) error {
	special, _, err := op.api.ItemApi.ApiAppItemSpecialItemsGet(context.Background()).Execute()
	if err != nil {
		return nil
	}
	var noteParent doggy.ItemDto
	for _, item := range special.Items {
		if item.Type == doggy.NOTE.Ptr() {
			noteParent = item
			break
		}
	}
	result, _, err := op.api.ItemApi.ApiAppItemGet(context.Background()).
		ParentItemId(*noteParent.Id).
		SkipCount((pageIndex - 1) * pageSize).
		MaxResultCount(pageSize).
		Execute()
	if err != nil {
		return err
	}
	prompt := &survey.Select{
		Message: "Choose a color:",
		Options: []string{"red", "blue", "green"},
	}
}
