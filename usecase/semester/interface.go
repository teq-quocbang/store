package semester

import (
	"context"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

type IUseCase interface {
	CreateSemester(context.Context, *payload.CreateSemesterRequest) (*presenter.SemesterResponseWrapper, error)
	GetList(context.Context, *payload.GetListSemesterRequest) (*presenter.ListSemesterResponseWrapper, error)
	GetByID(context.Context, string) (*presenter.SemesterResponseWrapper, error)
	Update(context.Context, *payload.UpdateSemesterRequest) (*presenter.SemesterResponseWrapper, error)
	Delete(context.Context, string) error
}
