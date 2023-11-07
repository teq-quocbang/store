package semester

import (
	"context"

	"github.com/teq-quocbang/store/model"
)

type Repository interface {
	Create(context.Context, *model.Semester) error
	GetListByYear(ctx context.Context, year string) ([]model.Semester, error)
	GetByID(ctx context.Context, semesterID string) (model.Semester, error)
	Update(context.Context, *model.Semester) error
	Delete(context.Context, string) error
}
