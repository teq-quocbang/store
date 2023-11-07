package semester

import (
	"context"

	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) GetByID(ctx context.Context, id string) (*presenter.SemesterResponseWrapper, error) {
	if id == "" {
		return nil, myerror.ErrSemesterInvalidParam("id")
	}
	semester, err := u.Semester.GetByID(ctx, id)
	if err != nil {
		return nil, myerror.ErrSemesterGet(err)
	}

	return &presenter.SemesterResponseWrapper{
		Semester: semester,
	}, nil
}
