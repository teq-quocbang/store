package storage

import (
	"context"

	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) GetListByLocat(ctx context.Context, locat string) (*presenter.ListStorageResponseWrapper, error) {
	if locat == "" {
		return nil, myerror.ErrStorageInvalidParam("missing locat")
	}

	storages, err := u.Storage.GetListStorageByLocat(ctx, locat)
	if err != nil {
		return nil, myerror.ErrStorageGet(err)
	}

	return &presenter.ListStorageResponseWrapper{
		Storage: storages,
	}, nil
}
