package product

import (
	"context"

	"github.com/google/uuid"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) Create(ctx context.Context, req *payload.CreateProductRequest) (*presenter.ProductResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrProductInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	product := model.Product{
		Name:        req.Name,
		ProductType: req.ProductType,
		ProducerID:  uuid.MustParse(req.ProducerID),
		CreatedBy:   userPrinciple.User.ID,
	}
	if err := u.Product.Create(ctx, product); err != nil {
		return nil, myerror.ErrProductCreate(err)
	}

	return &presenter.ProductResponseWrapper{
		Product: &product,
	}, nil
}
