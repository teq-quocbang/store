package course

import (
	"context"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) GetList(ctx context.Context, req *payload.ListCourseBySemesterRequest) (*presenter.ListCourseResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrCourseInvalidParam(err.Error())
	}

	course, err := u.Course.GetListBySemester(ctx, req.SemesterID)
	if err != nil {
		return nil, myerror.ErrClassGet(err)
	}

	return &presenter.ListCourseResponseWrapper{
		Course: course,
		Meta:   nil,
	}, nil
}
