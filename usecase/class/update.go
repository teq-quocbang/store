package class

import (
	"context"
	"time"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func validateUpdate(startTime time.Time, endTime time.Time, class *model.Class) error {
	// class is starting
	if class.StartTime.After(time.Now()) && class.EndTime.Before(time.Now()) {
		return myerror.ErrClassInvalidParam("class starting")
	}

	// class is ended
	if class.EndTime.Before(time.Now()) {
		return myerror.ErrClassInvalidParam("class ended")
	}

	return nil
}

func (u *UseCase) Update(ctx context.Context, req *payload.UpdateClassRequest) (*presenter.ClassResponseWrapper, error) {
	// validate check
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrClassInvalidParam(err.Error())
	}

	// get class
	class, err := u.Class.GetByID(ctx, req.ID)
	if err != nil {
		return nil, myerror.ErrClassGet(err)
	}

	start, end, err := parseTime(req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}

	if err := validateUpdate(*start, *end, &class); err != nil {
		return nil, err
	}

	// update
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	classModel := &model.Class{
		ID:         req.ID,
		CourseID:   req.CourseID,
		SemesterID: req.SemesterID,
		StartTime:  *start,
		EndTime:    *end,
		Credits:    uint(req.Credits),
		MaxSlot:    uint(req.MaxSlot),
		UpdatedBy:  &userPrinciple.User.ID,
	}

	err = u.Class.Update(ctx, classModel)
	if err != nil {
		return nil, myerror.ErrClassUpdate(err)
	}

	return &presenter.ClassResponseWrapper{
		Class: *classModel,
	}, nil
}

func (u *UseCase) InCreMember(ctx context.Context, id string) error {
	class, err := u.Class.GetByID(ctx, id)
	if err != nil {
		return myerror.ErrClassGet(err)
	}

	if class.EndTime.Before(time.Now()) {
		return myerror.ErrClassInvalidParam("class ended")
	}
	if class.CurrentSlot >= class.MaxSlot {
		return myerror.ErrClassInvalidParam("full slot")
	}

	if err := u.Class.BatchInCreMember(ctx, id); err != nil {
		return myerror.ErrClassUpdate(err)
	}

	return nil
}

func (u *UseCase) DeCreMember(ctx context.Context, id string) error {
	// check can cancel?
	_, err := u.Class.GetByID(ctx, id)
	if err != nil {
		return myerror.ErrClassGet(err)
	}

	if err := u.InCreMember(ctx, id); err != nil {
		return myerror.ErrClassUpdate(err)
	}
	return nil
}
