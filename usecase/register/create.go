package register

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/contexts"
	"github.com/teq-quocbang/store/util/myerror"
	"gorm.io/gorm"
)

func (u *UseCase) Create(ctx context.Context, req *payload.CreateRegisterRequest) (*presenter.RegisterResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrRegisterInvalidParam(err.Error())
	}

	// get user principle from context
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	// already existed
	register, err := u.Register.Get(ctx, &model.Register{
		AccountID:  userPrinciple.User.ID,
		SemesterID: req.SemesterID,
		ClassID:    req.ClassID,
		CourseID:   req.CourseID,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, myerror.ErrRegisterGet(err)
	}

	// if existed
	if !reflect.DeepEqual(register, &model.Register{}) {
		// do swap isCanceled if register already exited and isCanceled == true
		if register.IsCanceled {
			err := u.Register.BatchUpdateSwapIsCanCeledStatus(ctx, register)
			if err != nil {
				return nil, myerror.ErrRegisterUpdate(err)
			}
			register.IsCanceled = !register.IsCanceled // swap status

			semester, err := u.Semester.GetByID(ctx, req.SemesterID)
			if err != nil {
				return nil, myerror.ErrSemesterGet(err)
			}

			class, err := u.Class.GetByID(ctx, req.ClassID)
			if err != nil {
				return nil, myerror.ErrClassGet(err)
			}

			course, err := u.Course.GetByID(ctx, req.CourseID)
			if err != nil {
				return nil, myerror.ErrCourseGet(err)
			}

			return &presenter.RegisterResponseWrapper{
				Register: presenter.RegisterResponseCustom{
					AccountID: register.AccountID,
					Semester:  semester,
					Class:     class,
					Course:    course,
				},
			}, nil
		} else {
			semester, err := u.Semester.GetByID(ctx, register.SemesterID)
			if err != nil {
				return nil, myerror.ErrSemesterGet(err)
			}

			class, err := u.Class.GetByID(ctx, register.ClassID)
			if err != nil {
				return nil, myerror.ErrClassGet(err)
			}

			course, err := u.Course.GetByID(ctx, register.CourseID)
			if err != nil {
				return nil, myerror.ErrCourseGet(err)
			}

			return &presenter.RegisterResponseWrapper{
				Register: presenter.RegisterResponseCustom{
					AccountID: register.AccountID,
					Semester:  semester,
					Class:     class,
					Course:    course,
				},
			}, nil
		}
	}

	// can not register with same course code
	firstCourseChar := string(req.CourseID[0])
	registers, err := u.Register.GetListByFirstCourseChar(ctx, firstCourseChar, userPrinciple.User.ID, req.SemesterID)
	if err != nil {
		return nil, myerror.ErrRegisterGet(err)
	}
	if len(registers) != 0 {
		return nil, myerror.ErrCanNotRegisterSameCourse("Can not register with same course in one semester")
	}

	register = &model.Register{
		AccountID:  userPrinciple.User.ID,
		SemesterID: req.SemesterID,
		ClassID:    req.ClassID,
		CourseID:   req.CourseID,
		CreatedBy:  userPrinciple.User.ID,
	}
	if err := u.Register.Create(ctx, register); err != nil {
		return nil, myerror.ErrRegisterCreate(err)
	}

	semester, err := u.Semester.GetByID(ctx, req.SemesterID)
	if err != nil {
		return nil, myerror.ErrSemesterGet(err)
	}

	class, err := u.Class.GetByID(ctx, req.ClassID)
	if err != nil {
		return nil, myerror.ErrClassGet(err)
	}

	course, err := u.Course.GetByID(ctx, req.CourseID)
	if err != nil {
		return nil, myerror.ErrCourseGet(err)
	}

	// clear cache with prefix accountID*
	if err := u.Cache.Register().ClearRegisterHistories(ctx, fmt.Sprintf("%d", userPrinciple.User.ID)); err != nil {
		return nil, myerror.ErrFailedToRemoveCache(err)
	}

	return &presenter.RegisterResponseWrapper{
		Register: presenter.RegisterResponseCustom{
			Semester: semester,
			Class:    class,
			Course:   course,
		},
	}, nil
}
