package course

import (
	"context"

	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
)

type IUseCase interface {
	CreateCourse(context.Context, *payload.CreateCourseRequest) (*presenter.CourseResponseWrapper, error)
	GetList(context.Context, *payload.ListCourseBySemesterRequest) (*presenter.ListCourseResponseWrapper, error)
	GetByID(context.Context, string) (*presenter.CourseResponseWrapper, error)
	Update(context.Context, *payload.UpdateCourseRequest) (*presenter.CourseResponseWrapper, error)
	Delete(context.Context, string) error
}
