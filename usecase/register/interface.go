package register

import (
	"context"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

type IUseCase interface {
	Create(context.Context, *payload.CreateRegisterRequest) (*presenter.RegisterResponseWrapper, error)
	GetListBySemester(context.Context, *payload.ListRegisterInformationRequest) (*presenter.ListRegisterResponseWrapper, error)
	UnRegister(context.Context, *payload.UnRegisterRequest) (*presenter.RegisterResponseWrapper, error)
	GetListRegisteredHistories(context.Context, *payload.ListRegisteredHistories) (*presenter.ListRegisterResponseWrapper, error)
	TracingInsufficientCreditsStatistics(context.Context) error
}
