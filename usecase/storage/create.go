package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) UpsertStorage(ctx context.Context, req *payload.UpsertStorageRequest) (*presenter.StorageResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrStorageInvalidParam(err.Error())
	}

	productID, err := uuid.Parse(req.ProductID)
	if err != nil {
		return nil, myerror.ErrStorageInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	storage := &model.Storage{
		Locat:        req.Locat,
		ProductID:    productID,
		InventoryQty: req.Qty,
		CreatedBy:    userPrinciple.User.ID,
		UpdatedBy:    userPrinciple.User.ID,
	}
	if err := u.Storage.UpsertStorage(ctx, storage); err != nil {
		return nil, myerror.ErrStorageCreate(err)
	}

	return &presenter.StorageResponseWrapper{
		Storage: storage,
	}, nil
}
