package semester

import (
	"context"
	"time"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) Update(ctx context.Context, req *payload.UpdateSemesterRequest) (*presenter.SemesterResponseWrapper, error) {
	// validate check
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrSemesterInvalidParam(err.Error())
	}

	// get semester
	semester, err := u.Semester.GetByID(ctx, req.ID)
	if err != nil {
		return nil, myerror.ErrSemesterGet(err)
	}

	// check whether the semester is ended
	if semester.EndTime.Before(time.Now()) {
		return nil, myerror.ErrSemesterInvalidParam("the semester was ended")
	}

	start, end, registerStartAt, registerExpiresAt, err := parseStringToTime(req.StartTime, req.EndTime, req.RegisterStartAt, req.RegisterExpiresAt)
	if err != nil {
		return nil, err
	}

	// update
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	semesterModel := &model.Semester{
		ID:                req.ID,
		MinCredits:        req.MinCredits,
		StartTime:         *start,
		EndTime:           *end,
		RegisterStartAt:   *registerStartAt,
		RegisterExpiresAt: *registerExpiresAt,
		UpdatedBy:         &userPrinciple.User.ID,
	}
	err = u.Semester.Update(ctx, semesterModel)
	if err != nil {
		return nil, myerror.ErrSemesterUpdate(err)
	}

	return &presenter.SemesterResponseWrapper{
		Semester: *semesterModel,
	}, nil
}
