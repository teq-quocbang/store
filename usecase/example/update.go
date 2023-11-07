package example

import (
	"context"
	"strings"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) validateUpdate(ctx context.Context, req *payload.UpdateExampleRequest) (*model.Example, error) {
	myExample, err := u.ExampleRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrExampleNotFound()
		}

		return nil, myerror.ErrExampleGet(err)
	}

	if req.Name != nil {
		*req.Name = strings.TrimSpace(*req.Name)
		if len(*req.Name) == 0 {
			return nil, myerror.ErrExampleInvalidParam("name")
		}

		myExample.Name = *req.Name
	}

	myExample.UpdatedBy = teq.Int64(1)

	return myExample, nil
}

func (u *UseCase) Update(
	ctx context.Context,
	req *payload.UpdateExampleRequest,
) (*presenter.ExampleResponseWrapper, error) {
	myExample, err := u.validateUpdate(ctx, req)
	if err != nil {
		return nil, err
	}

	err = u.ExampleRepo.Update(ctx, myExample)
	if err != nil {
		return nil, myerror.ErrExampleUpdate(err)
	}

	return &presenter.ExampleResponseWrapper{Example: myExample}, nil
}
