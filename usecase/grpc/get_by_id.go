package grpc

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/util/myerror"
)

type GetByIDRequest struct {
	ID int64 `json:"-"`
}

func (u *UseCase) GetByID(ctx context.Context, req *GetByIDRequest) (*model.Example, error) {
	myExample, err := u.ExampleRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrExampleNotFound()
		}

		return nil, myerror.ErrExampleGet(err)
	}

	return myExample, nil
}
