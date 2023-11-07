package course

import (
	"context"

	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return myerror.ErrCourseInvalidParam("id")
	}

	_, err := u.Course.GetByID(ctx, id)
	if err != nil {
		return myerror.ErrCourseGet(err)
	}

	err = u.Course.Delete(ctx, id)
	if err != nil {
		return myerror.ErrSemesterDelete(err)
	}

	return nil
}
