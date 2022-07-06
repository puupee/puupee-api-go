package cli

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mr-doggy/doggy-sdk-go/doggy"
)

type AppOp struct {
	api *doggy.APIClient
}

func NewAppOp(api *doggy.APIClient) *AppOp {
	return &AppOp{
		api: api,
	}
}

func (op *AppOp) Create(dto doggy.CreateOrUpdateAppDto) error {
	resp, _, err := op.api.AppApi.ApiAppAppPost(context.Background()).CreateOrUpdateAppDto(dto).Execute()
	if err != nil {
		return err
	}
	PrettyPrint(resp)
	return nil
}

func (op *AppOp) List() error {
	dto, _, err := op.api.AppApi.ApiAppAppGet(context.Background()).MaxResultCount(100).Execute()
	if err != nil {
		return err
	}
	bts, _ := json.MarshalIndent(dto, "", "  ")
	fmt.Println(string(bts))
	return nil
}
