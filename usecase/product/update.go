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

func (u *UseCase) Update(ctx context.Context, req *payload.UpdateProductRequest) (*presenter.ProductResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrProductInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	producer := uuid.UUID{}
	var err error
	if req.ProducerID != "" {
		producer, err = uuid.Parse(req.ProducerID)
		if err != nil {
			return nil, myerror.ErrProductInvalidParam(err.Error())
		}
	}
	product := &model.Product{
		ID:          uuid.MustParse(req.ID),
		Name:        req.Name,
		ProductType: req.ProductType,
		ProducerID:  producer,
		UpdatedBy:   userPrinciple.User.ID,
	}
	if err := u.Product.Update(ctx, product); err != nil {
		return nil, myerror.ErrProductUpdate(err)
	}

	return &presenter.ProductResponseWrapper{
		Product: product,
	}, nil
}
