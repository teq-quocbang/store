package class

import (
	"context"

	"github.com/teq-quocbang/store/model"
)

type Repository interface {
	Create(context.Context, *model.Class) error
	GetListBySemester(ctx context.Context, semesterID string) ([]model.Class, error)
	GetByID(ctx context.Context, classID string) (model.Class, error)
	Update(context.Context, *model.Class) error
	Delete(context.Context, string) error
	BatchInCreMember(context.Context, string) error
	BatchDeCreMember(context.Context, string) error
}
