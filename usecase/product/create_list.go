package product

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) CreateList(ctx context.Context, req *payload.CreateListProductRequest) (*presenter.ListProductResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrProductInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	// get list and filter id already existed
	productName := map[string]struct{}{}
	if products, err := u.Product.GetList(ctx); err != nil {
		return nil, myerror.ErrProducerGet(err)
	} else {
		for _, p := range products {
			productName[p.Name] = struct{}{}
		}
	}

	productModel := []model.Product{}
	for i, p := range req.Products {
		if _, ok := productName[p.Name]; !ok {
			producerID, err := uuid.Parse(p.ProducerID)
			if err != nil {
				return nil, myerror.ErrProductInvalidParam(fmt.Sprintf("error at index [%d], error: %v", i, err))
			}
			price, err := decimal.NewFromString(p.Price)
			if err != nil {
				return nil, myerror.ErrProductInvalidParam(err.Error())
			}
			productModel = append(productModel, model.Product{
				Name:        p.Name,
				ProductType: p.ProductType,
				ProducerID:  producerID,
				Price:       price,
				CreatedBy:   userPrinciple.User.ID,
				UpdatedBy:   userPrinciple.User.ID,
			})
		}
	}

	if err := u.Product.CreateList(ctx, productModel); err != nil {
		return nil, myerror.ErrProductCreate(err)
	}

	return &presenter.ListProductResponseWrapper{
		Product: productModel,
	}, nil
}
