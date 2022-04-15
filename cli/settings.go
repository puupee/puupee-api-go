package cli

import "github.com/mr-doggy/doggy-sdk-go/doggy"

type SettingsOp struct {
}

func NewSettingsOp() *SettingsOp {
	return &SettingsOp{}
}

func (op *SettingsOp) Set(value doggy.SettingsDto) error {
	return nil
}

func (op *SettingsOp) Get() (doggy.SettingsDto, error) {
	return doggy.SettingsDto{}, nil
}
