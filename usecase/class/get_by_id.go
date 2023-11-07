package class

import (
	"context"

	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) GetByID(ctx context.Context, id string) (*presenter.ClassResponseWrapper, error) {
	if id == "" {
		return nil, myerror.ErrClassInvalidParam("id")
	}

	class, err := u.Class.GetByID(ctx, id)
	if err != nil {
		return nil, myerror.ErrClassGet(err)
	}

	return &presenter.ClassResponseWrapper{
		Class: class,
	}, nil
}
