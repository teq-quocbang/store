package producer

import (
	"context"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) Create(ctx context.Context, req *payload.CreateProducerRequest) (*presenter.ProducerResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrProducerInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	producer := &model.Producer{
		Name:      req.Name,
		Country:   req.Country,
		CreatedBy: userPrinciple.User.ID,
		UpdatedBy: userPrinciple.User.ID,
	}
	if err := u.Producer.Create(ctx, producer); err != nil {
		return nil, myerror.ErrProducerCreate(err)
	}

	return &presenter.ProducerResponseWrapper{
		Producer: producer,
	}, nil
}
