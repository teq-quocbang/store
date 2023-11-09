package product

import (
	"context"
	"reflect"

	"github.com/google/uuid"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/replace"
)

func (u *UseCase) validateUpdate(ctx context.Context, req *payload.UpdateProductRequest, productID uuid.UUID) (*model.Product, error) {
	producer := uuid.UUID{}
	var err error
	if req.ProducerID != "" {
		producer, err = uuid.Parse(req.ProducerID)
		if err != nil {
			return nil, myerror.ErrProductInvalidParam(err.Error())
		}
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	product, err := u.Product.GetByID(ctx, productID)
	if err != nil {
		return nil, myerror.ErrProductGet(err)
	}
	productModel, err := replace.Replace[payload.UpdateProductRequest, model.Product](*req, product)
	if err != nil {
		return nil, myerror.ErrProductInvalidParam(err.Error())
	}

	// set dependency
	if !reflect.DeepEqual(producer, uuid.UUID{}) {
		productModel.ProducerID = producer
	}
	productModel.ID = productID
	productModel.UpdatedBy = userPrinciple.User.ID

	return &productModel, nil
}

func (u *UseCase) Update(ctx context.Context, req *payload.UpdateProductRequest) (*presenter.ProductResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrProductInvalidParam(err.Error())
	}

	productID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, myerror.ErrProductInvalidParam(err.Error())
	}

	product, err := u.validateUpdate(ctx, req, productID)
	if err != nil {
		return nil, err
	}
	if err := u.Product.Update(ctx, product); err != nil {
		return nil, myerror.ErrProductUpdate(err)
	}

	return &presenter.ProductResponseWrapper{
		Product: product,
	}, nil
}
