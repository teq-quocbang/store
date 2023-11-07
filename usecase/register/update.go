package register

import (
	"context"
	"fmt"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) UnRegister(ctx context.Context, req *payload.UnRegisterRequest) (*presenter.RegisterResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrRegisterInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	register, err := u.Register.Get(ctx, &model.Register{
		AccountID:  userPrinciple.User.ID,
		SemesterID: req.SemesterID,
		ClassID:    req.ClassID,
		CourseID:   req.CourseID,
	})
	if err != nil {
		return nil, myerror.ErrRegisterGet(err)
	}
	if register.IsCanceled {
		return &presenter.RegisterResponseWrapper{
			Register: presenter.RegisterResponseCustom{},
		}, nil
	}
	if register.CreatedBy != userPrinciple.User.ID {
		return nil, myerror.ErrRegisterInvalidParam("not permission")
	}

	register = &model.Register{
		AccountID:  userPrinciple.User.ID,
		SemesterID: req.SemesterID,
		ClassID:    req.ClassID,
		CourseID:   req.CourseID,
		IsCanceled: register.IsCanceled,
	}

	if err := u.Register.BatchUpdateSwapIsCanCeledStatus(ctx, register); err != nil {
		return nil, myerror.ErrRegisterUpdate(err)
	}
	if err != nil {
		return nil, myerror.ErrRegisterGet(err)
	}

	// clear cache with prefix accountID*
	if err := u.Cache.Register().ClearRegisterHistories(ctx, fmt.Sprintf("%d", userPrinciple.User.ID)); err != nil {
		return nil, myerror.ErrFailedToRemoveCache(err)
	}

	return &presenter.RegisterResponseWrapper{
		Register: presenter.RegisterResponseCustom{},
	}, nil
}
