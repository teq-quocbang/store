package product

import (
	"context"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return myerror.ErrProductInvalidParam("missing id")
	}

	productID, err := uuid.Parse(id)
	if err != nil {
		return myerror.ErrProductInvalidParam(err.Error())
	}
	if err := u.Product.Delete(ctx, productID); err != nil {
		return myerror.ErrProductDelete(err)
	}

	return nil
}
