package semester

import (
	"context"

	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return myerror.ErrSemesterInvalidParam("id")
	}

	// check semester
	_, err := u.Semester.GetByID(ctx, id)
	if err != nil {
		return myerror.ErrSemesterGet(err)
	}

	err = u.Semester.Delete(ctx, id)
	if err != nil {
		return myerror.ErrSemesterDelete(err)
	}

	return nil
}
