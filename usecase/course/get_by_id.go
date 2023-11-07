package course

import (
	"context"

	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) GetByID(ctx context.Context, id string) (*presenter.CourseResponseWrapper, error) {
	if id == "" {
		return nil, myerror.ErrCourseInvalidParam("id")
	}
	course, err := u.Course.GetByID(ctx, id)
	if err != nil {
		return nil, myerror.ErrCourseGet(err)
	}

	return &presenter.CourseResponseWrapper{
		Course: course,
	}, nil
}
