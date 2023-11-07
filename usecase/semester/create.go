package semester

import (
	"context"
	"fmt"
	"time"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/times"
)

// read wiki for details https://en.wikipedia.org/wiki/Currying
func currying[t1 any, R1 time.Time, R2 error](f func(t1) (R1, R2)) func(t1) (R1, R2) {
	return func(t t1) (R1, R2) {
		return f(t)
	}
}

// read wiki for details https://en.wikipedia.org/wiki/Currying
// currying with one param and one return
func curryingOne[t1, R any](f func(t1) R) func(t1) R {
	return func(t1 t1) R {
		return f(t1)
	}
}

func isLessThanCurrentTime(t time.Time) bool {
	return t.Before(time.Now())
}

func validateCreate(start time.Time, end time.Time, registerStartAt time.Time, registerExpiresAt time.Time) error {
	// start currying
	lessThanCurrentTime := curryingOne(isLessThanCurrentTime)

	// validate whether is less than current time
	if lessThan := lessThanCurrentTime(start); lessThan {
		return myerror.ErrSemesterInvalidParam("start time less than current time")
	}
	// validate whether is less than current time
	if lessThan := lessThanCurrentTime(end); lessThan {
		return myerror.ErrSemesterInvalidParam("end time less than current time")
	}
	// validate whether is less than current time
	if lessThan := lessThanCurrentTime(registerStartAt); lessThan {
		return myerror.ErrSemesterInvalidParam("register start at less than current time")
	}
	// validate whether is less than current time
	if lessThan := lessThanCurrentTime(registerExpiresAt); lessThan {
		return myerror.ErrSemesterInvalidParam("register expires at less than current time")
	}

	// at least 3 month
	if ok := times.IsLessThan(start, end, times.ThreeMonth); ok {
		return myerror.ErrSemesterInvalidParam("a semester at least 3 month")
	}

	// maximum 6 month
	if ok := times.IsMoreThan(start, end, times.SixMonth); ok {
		return myerror.ErrSemesterInvalidParam("a semester maximum is six month")
	}
	return nil
}

func parseStringToTime(
	startTime string,
	endTime string,
	registerStartAt string,
	registerExpiresAt string) (*time.Time, *time.Time, *time.Time, *time.Time, error) {
	// start currying
	parseTime := currying(times.StringToTime)

	// curried start time
	start, err := parseTime(startTime)
	if err != nil {
		return nil, nil, nil, nil, myerror.ErrSemesterInvalidParam(fmt.Sprintf("start time, error: %v", err))
	}

	// curried end time
	end, err := parseTime(endTime)
	if err != nil {
		return nil, nil, nil, nil, myerror.ErrSemesterInvalidParam(fmt.Sprintf("end time, error: %v", err))
	}

	// curried register start at
	rsa, err := parseTime(registerStartAt)
	if err != nil {
		return nil, nil, nil, nil, myerror.ErrSemesterInvalidParam(fmt.Sprintf("register start at, error: %v", err))
	}

	// curried register expiries at
	rexpa, err := parseTime(registerExpiresAt)
	if err != nil {
		return nil, nil, nil, nil, myerror.ErrSemesterInvalidParam(fmt.Sprintf("register expires at, error: %v", err))
	}

	return &start, &end, &rsa, &rexpa, nil
}

func (u *UseCase) CreateSemester(ctx context.Context, req *payload.CreateSemesterRequest) (*presenter.SemesterResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrSemesterInvalidParam(err.Error())
	}

	// get user principle from context
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	startTime, endTime, registerStartAt, registerExpiresAt, err := parseStringToTime(req.StartTime, req.EndTime, req.RegisterStartAt, req.RegisterExpiresAt)
	if err != nil {
		return nil, err
	}

	// validate time
	if err := validateCreate(*startTime, *endTime, *registerStartAt, *registerExpiresAt); err != nil {
		return nil, err
	}

	semester := &model.Semester{
		ID:                req.ID,
		MinCredits:        req.MinCredits,
		StartTime:         *startTime,
		EndTime:           *endTime,
		RegisterStartAt:   *registerStartAt,
		RegisterExpiresAt: *registerExpiresAt,
		CreatedBy:         &userPrinciple.User.ID,
	}
	if err := u.Semester.Create(ctx, semester); err != nil {
		return nil, myerror.ErrSemesterCreate(err)
	}

	return &presenter.SemesterResponseWrapper{
		Semester: *semester,
	}, nil
}
