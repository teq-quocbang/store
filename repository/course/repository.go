package course

import (
	"context"

	"github.com/teq-quocbang/store/model"
)

type Repository interface {
	Create(context.Context, *model.Course) error
	GetListBySemester(ctx context.Context, semesterID string) ([]model.Course, error)
	GetByID(ctx context.Context, courseID string) (model.Course, error)
	Update(context.Context, *model.Course) error
	Delete(context.Context, string) error
}
