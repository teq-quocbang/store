package register

import (
	"context"
	"fmt"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) GetListBySemester(ctx context.Context, req *payload.ListRegisterInformationRequest) (*presenter.ListRegisterResponseWrapper, error) {
	req.Format()
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrRegisterInvalidParam(err.Error())
	}

	var (
		order = make([]string, 0)
	)
	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s %s", req.OrderBy, req.SortBy))
	}

	// get user principle from context
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	registers, total, err := u.Register.GetListBySemesterID(ctx, userPrinciple.User.ID, req.SemesterID, order, req.Paginator)
	if err != nil {
		return nil, myerror.ErrSemesterGet(err)
	}

	response := &presenter.ListRegisterResponseWrapper{
		Register: make([]presenter.RegisterResponseCustom, len(registers)),
		Meta: map[string]interface{}{
			"page":  req.Paginator.Page,
			"limit": req.Paginator.Limit,
			"total": total,
		},
	}
	for i, r := range registers {
		semester, err := u.Semester.GetByID(ctx, r.SemesterID)
		if err != nil {
			return nil, myerror.ErrSemesterGet(err)
		}

		class, err := u.Class.GetByID(ctx, r.ClassID)
		if err != nil {
			return nil, myerror.ErrClassGet(err)
		}

		course, err := u.Course.GetByID(ctx, r.CourseID)
		if err != nil {
			return nil, myerror.ErrCourseGet(err)
		}

		response.Register[i] = presenter.RegisterResponseCustom{
			AccountID: r.AccountID,
			Semester:  semester,
			Class:     class,
			Course:    course,
		}
	}

	return response, nil
}
