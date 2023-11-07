package class

import (
	"context"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) GetList(ctx context.Context, req *payload.ListClassBySemesterRequest) (*presenter.ListClassResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrClassInvalidParam(err.Error())
	}

	classes, err := u.Class.GetListBySemester(ctx, req.SemesterID)
	if err != nil {
		return nil, myerror.ErrClassGet(err)
	}

	return &presenter.ListClassResponseWrapper{
		Class: classes,
		Meta:  nil,
	}, nil
}
