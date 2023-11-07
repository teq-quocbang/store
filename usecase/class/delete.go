package class

import (
	"context"
	"time"

	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return myerror.ErrClassInvalidParam("id")
	}

	class, err := u.Class.GetByID(ctx, id)
	if err != nil {
		return myerror.ErrClassGet(err)
	}

	if class.StartTime.Before(time.Now()) && class.EndTime.After(time.Now()) {
		return myerror.ErrClassInvalidParam("class is going on")
	}

	if class.EndTime.Before(time.Now()) {
		return myerror.ErrClassInvalidParam("class ended")
	}

	err = u.Class.Delete(ctx, id)
	if err != nil {
		return myerror.ErrSemesterDelete(err)
	}

	return nil
}
