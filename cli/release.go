package cli

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mr-doggy/doggy-sdk-go/doggy"
)

type ReleaseOp struct {
	api *doggy.APIClient
}

func NewReleaseOp(api *doggy.APIClient) *ReleaseOp {
	return &ReleaseOp{
		api: api,
	}
}

func (op *ReleaseOp) Create(dto doggy.CreateOrUpdateAppReleaseDto) error {
	resp, _, err := op.api.AppReleaseApi.ApiAppAppReleasePost(context.Background()).CreateOrUpdateAppReleaseDto(dto).Execute()
	if err != nil {
		return err
	}
	PrettyPrint(resp)
	return nil
}

func (op *ReleaseOp) List(appName string) error {
	appDto, _, err := op.api.AppApi.ApiAppAppByNameGet(context.Background()).Name(appName).Execute()
	if err != nil {
		return err
	}
	dto, _, err := op.api.AppReleaseApi.ApiAppAppReleaseGet(context.Background()).AppId(*appDto.Id).MaxResultCount(100).Execute()
	if err != nil {
		return err
	}
	bts, _ := json.MarshalIndent(dto, "", "  ")
	fmt.Println(string(bts))
	return nil
}
