package example

import (
	"context"
	"strings"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) validateCreate(req *payload.CreateExampleRequest) error {
	if req.Name == nil {
		return myerror.ErrExampleInvalidParam("name")
	}

	*req.Name = strings.TrimSpace(*req.Name)
	if len(*req.Name) == 0 {
		return myerror.ErrExampleInvalidParam("name")
	}

	return nil
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateExampleRequest,
) (*presenter.ExampleResponseWrapper, error) {
	if err := u.validateCreate(req); err != nil {
		return nil, err
	}

	myExample := &model.Example{
		Name:      *req.Name,
		CreatedBy: 1, // must be validate logged user.
	}

	err := u.ExampleRepo.Create(ctx, myExample)
	if err != nil {
		return nil, myerror.ErrExampleCreate(err)
	}

	return &presenter.ExampleResponseWrapper{Example: myExample}, nil
}
