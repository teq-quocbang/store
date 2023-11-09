package product

import (
	"context"
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func replace[T1 any, T2 any](t1 T1, t2 T2) (T2, error) {
	if reflect.TypeOf(*new(T1)).Kind() == reflect.Pointer ||
		reflect.TypeOf(*new(T2)).Kind() == reflect.Pointer {
		return t2, fmt.Errorf("can not send pointer type")
	}
	valueT1 := reflect.ValueOf(t1)
	typeT1 := reflect.TypeOf(t1)
	valueT2 := reflect.ValueOf(&t2)
	typeT2 := reflect.TypeOf(t2)

	for i := 0; i < valueT1.NumField(); i++ {
		fieldT1 := valueT1.Field(i)
		tag1 := typeT1.Field(i).Tag.Get("json")

		// if is different
		if !fieldT1.IsZero() {
			// check all field of model if found json tag is same let replace at t2(model) value
			for i := 0; i < valueT2.Elem().NumField(); i++ {
				tag2 := typeT2.Field(i).Tag.Get("json")
				if tag1 == tag2 {
					if valueT2.Elem().Field(i).CanSet() {
						if valueT2.Elem().Field(i).Type() != fieldT1.Type() {
							break // TODO: need to same type
						}
						valueT2.Elem().Field(i).Set(fieldT1)
					}
					break
				}
			}
		}
	}

	return t2, nil // return model with replaced value
}

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
	productModel, err := replace[payload.UpdateProductRequest, model.Product](*req, product)
	if err != nil {
		return nil, myerror.ErrProductInvalidParam(err.Error())
	}

	// set dependency
	productModel.ProducerID = producer
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
