package register

import (
	"context"

	"github.com/teq-quocbang/store/codetype"
	"github.com/teq-quocbang/store/model"
)

type Repository interface {
	Create(context.Context, *model.Register) error
	GetListBySemesterID(ctx context.Context, accountID uint, semesterID string, order []string, paginator codetype.Paginator) ([]model.Register, int64, error)
	// GetListByFirstCourseChar is get list all the course that student registered
	// use the first character of course_id
	// ex:
	//		student registered S0001, T0001, M0001
	// with S in param so the result is:
	// 		[S0001]
	GetListByFirstCourseChar(ctx context.Context, firstChar string, accountID uint, semesterID string) ([]model.Register, error)
	Get(context.Context, *model.Register) (*model.Register, error)
	// swap the state of the is_canceled field
	// false -> true and true -> false
	BatchUpdateSwapIsCanCeledStatus(context.Context, *model.Register) error
	GetListRegistered(ctx context.Context, accountID uint, semesterID string, order []string, paginator codetype.Paginator) ([]model.Register, int64, error)
}
