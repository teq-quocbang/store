package register

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
)

func (u *UseCase) GetListRegisteredHistories(ctx context.Context, req *payload.ListRegisteredHistories) (*presenter.ListRegisterResponseWrapper, error) {
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	req.Format()

	var (
		order = make([]string, 0)
	)

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s %s", req.OrderBy, req.SortBy))
	}

	// get in cache
	registerHistoriesSaveCacheKey := fmt.Sprintf("%d|%s|%d|%d|%s|%s",
		userPrinciple.User.ID,
		req.SemesterID,
		req.Paginator.Page,
		req.Paginator.Limit,
		req.OrderBy,
		req.SortBy)
	registerHistories, err := u.Cache.Register().GetRegisterHistories(ctx, registerHistoriesSaveCacheKey)
	if err != nil && !errors.Is(err, redis.Nil) {
		if err != nil {
			return nil, myerror.ErrFailedToGetCache(err)
		}
	}
	if registerHistories != nil {
		return registerHistories, nil
	}

	registers, total, err := u.Register.GetListRegistered(ctx, userPrinciple.User.ID, req.SemesterID, order, req.Paginator)
	if err != nil {
		return nil, myerror.ErrRegisterGet(err)
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

	// save to cache
	if err := u.Cache.Register().CreateRegisterHistories(ctx, registerHistoriesSaveCacheKey, response); err != nil {
		return nil, myerror.ErrFailedToSaveCache(err)
	}

	return response, err
}
