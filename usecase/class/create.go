package class

import (
	"context"
	"reflect"
	"time"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/times"
)

func isLessThanCurrentTime(t time.Time) bool {
	return t.Before(time.Now())
}

func parseTime(classStart string, classEnd string) (*time.Time, *time.Time, error) {
	start, end := time.Time{}, time.Time{}
	var err error
	if classStart != "" {
		start, err = times.StringToTime(classStart)
		if err != nil {
			return nil, nil, myerror.ErrClassInvalidParam(err.Error())
		}
	}

	if classEnd != "" {
		end, err = times.StringToTime(classEnd)
		if err != nil {
			return nil, nil, myerror.ErrClassInvalidParam(err.Error())
		}
	}

	if !reflect.DeepEqual(start, time.Time{}) {
		if isLessThanCurrentTime(start) {
			return nil, nil, myerror.ErrClassInvalidParam("start is less than current time")
		}
	}

	if !reflect.DeepEqual(end, time.Time{}) {
		if isLessThanCurrentTime(end) {
			return nil, nil, myerror.ErrClassInvalidParam("end is less than current time")
		}
	}

	return &start, &end, nil
}

func validateCreate(classStart string, classEnd string, semesterStart time.Time, semesterEnd time.Time) (*time.Time, *time.Time, error) {
	start, end, err := parseTime(classStart, classEnd)
	if err != nil {
		return nil, nil, err
	}
	if ok := isClassTimeOutOfSemester(*start, *end, semesterStart, semesterEnd); ok {
		return nil, nil, myerror.ErrClassInvalidParam("out of semester time")
	}

	return start, end, nil
}

func isClassTimeOutOfSemester(
	classStart time.Time,
	classEnd time.Time,
	semesterStart time.Time,
	semesterEnd time.Time,
) bool {
	if classStart.Before(semesterStart) || classEnd.Before(semesterStart) {
		return true
	}
	if classStart.After(semesterEnd) || classEnd.After(semesterEnd) {
		return true
	}
	return false
}

func (u *UseCase) CreateClass(ctx context.Context, req *payload.CreateClassRequest) (*presenter.ClassResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrClassInvalidParam(err.Error())
	}

	// check whether out of the semester
	semester, err := u.Semester.GetByID(ctx, req.SemesterID)
	if err != nil {
		return nil, myerror.ErrSemesterGet(err)
	}

	// check time
	start, end, err := validateCreate(req.StartTime, req.EndTime, semester.StartTime, semester.EndTime)
	if err != nil {
		return nil, err
	}

	// check course
	_, err = u.Course.GetByID(ctx, req.CourseID)
	if err != nil {
		return nil, myerror.ErrCourseGet(err)
	}

	// get user principle from context
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	class := &model.Class{
		ID:         req.ID,
		CourseID:   req.CourseID,
		SemesterID: req.SemesterID,
		StartTime:  *start,
		EndTime:    *end,
		Credits:    uint(req.Credits),
		MaxSlot:    uint(req.MaxSlot),
		CreatedBy:  &userPrinciple.User.ID,
	}
	if err := u.Class.Create(ctx, class); err != nil {
		return nil, myerror.ErrClassCreate(err)
	}

	return &presenter.ClassResponseWrapper{
		Class: *class,
	}, nil
}
