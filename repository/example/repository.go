package example

import (
	"context"

	"github.com/teq-quocbang/store/codetype"
	"github.com/teq-quocbang/store/model"
)

type Repository interface {
	Create(ctx context.Context, data *model.Example) error
	CreateList(ctx context.Context, data []model.Example) error
	CreateOrUpdate(ctx context.Context, data *model.Example) error
	CreateOrUpdateList(ctx context.Context, data []model.Example) error
	Update(ctx context.Context, data *model.Example) error
	UpdateList(ctx context.Context, data []model.Example) error
	GetByID(ctx context.Context, id int64) (*model.Example, error)
	GetByInterface(ctx context.Context, itf interface{}) (*model.Example, error)
	GetListByInterface(ctx context.Context, itf interface{}) ([]model.Example, error)
	Delete(ctx context.Context, data *model.Example, unscoped bool) error
	DeleteList(ctx context.Context, data []model.Example, unscoped bool) error
	GetAll(ctx context.Context, unscoped bool) ([]model.Example, error)
	GetList(
		ctx context.Context,
		search string,
		paginator codetype.Paginator,
		conditions interface{},
		order []string,
	) ([]model.Example, int64, error)
}
