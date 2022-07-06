package cli

import "github.com/mr-doggy/doggy-sdk-go/doggy"

type ProfileOp struct {
}

func NewProfileOp() *ProfileOp {
	return &ProfileOp{}
}

func (op *ProfileOp) Update(value doggy.ProfileDto) error {

	return nil
}
