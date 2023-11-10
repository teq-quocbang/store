package product

import (
	"context"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

type IUseCase interface {
	Create(context.Context, *payload.CreateProductRequest) (*presenter.ProductResponseWrapper, error)
	CreateList(context.Context, *payload.CreateListProductRequest) (*presenter.ListProductResponseWrapper, error)
	Update(context.Context, *payload.UpdateProductRequest) (*presenter.ProductResponseWrapper, error)
	Delete(context.Context, string) error
	Export(context.Context, *payload.ExportProductRequest) (*presenter.ListProductResponseWrapper, error)
}
