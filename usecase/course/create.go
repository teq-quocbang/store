package course

import (
	"context"
	"time"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func validateCreate(semester model.Semester) error {
	if semester.EndTime.Before(time.Now()) {
		return myerror.ErrCourseInvalidParam("semester ended")
	}
	return nil
}

func (u *UseCase) CreateCourse(ctx context.Context, req *payload.CreateCourseRequest) (*presenter.CourseResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrClassInvalidParam(err.Error())
	}

	// check whether out of the semester
	semester, err := u.Semester.GetByID(ctx, req.SemesterID)
	if err != nil {
		return nil, myerror.ErrSemesterGet(err)
	}

	if err := validateCreate(semester); err != nil {
		return nil, err
	}

	// get user principle from context
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	course := &model.Course{
		ID:         req.ID,
		SemesterID: req.SemesterID,
		CreatedBy:  &userPrinciple.User.ID,
	}

	if err := u.Course.Create(ctx, course); err != nil {
		return nil, myerror.ErrClassCreate(err)
	}

	return &presenter.CourseResponseWrapper{
		Course: *course,
	}, nil
}
