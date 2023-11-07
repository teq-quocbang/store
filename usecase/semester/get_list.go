package semester

import (
	"context"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) GetList(ctx context.Context, req *payload.GetListSemesterRequest) (*presenter.ListSemesterResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrSemesterInvalidParam(err.Error())
	}

	semesters, err := u.Semester.GetListByYear(ctx, req.Year)
	if err != nil {
		return nil, myerror.ErrSemesterGet(err)
	}

	return &presenter.ListSemesterResponseWrapper{
		Semester: semesters,
	}, nil
}
