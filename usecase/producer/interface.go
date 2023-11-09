package producer

import (
	"context"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

type IUseCase interface {
	Create(context.Context, *payload.CreateProducerRequest) (*presenter.ProducerResponseWrapper, error)
}
