package storage

import (
	"context"

	"github.com/teq-quocbang/store/model"
)

type Repository interface {
	UpsertStorage(context.Context, *model.Storage) error
	GetListStorageByLocat(ctx context.Context, locat string) ([]model.Storage, error)
}
