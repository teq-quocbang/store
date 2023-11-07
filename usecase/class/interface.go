package class

import (
	"context"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

type IUseCase interface {
	CreateClass(context.Context, *payload.CreateClassRequest) (*presenter.ClassResponseWrapper, error)
	GetList(context.Context, *payload.ListClassBySemesterRequest) (*presenter.ListClassResponseWrapper, error)
	GetByID(context.Context, string) (*presenter.ClassResponseWrapper, error)
	Update(context.Context, *payload.UpdateClassRequest) (*presenter.ClassResponseWrapper, error)
	Delete(context.Context, string) error
	InCreMember(context.Context, string) error
	DeCreMember(context.Context, string) error
}
