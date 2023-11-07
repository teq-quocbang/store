package course

import (
	"context"

	"github.com/teq-quocbang/store/model"
	"gorm.io/gorm"
)

type pgRepository struct {
	getDB func(context.Context) *gorm.DB
}

func NewCoursePG(getDB func(context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB: getDB}
}

func (p *pgRepository) Create(ctx context.Context, s *model.Course) error {
	return p.getDB(ctx).Create(&s).Error
}

func (p *pgRepository) GetListBySemester(ctx context.Context, courseID string) ([]model.Course, error) {
	var course []model.Course
	err := p.getDB(ctx).Where(`semester_id = ?`, courseID).Find(&course).Error
	return course, err
}

func (p *pgRepository) GetByID(ctx context.Context, courseID string) (model.Course, error) {
	var course model.Course
	err := p.getDB(ctx).Where(`id = ?`, courseID).Take(&course).Error
	return course, err
}

func (p *pgRepository) Update(ctx context.Context, req *model.Course) error {
	conditions := req.BuildUpdateFields()
	return p.getDB(ctx).Model(&model.Course{}).Where(`id = ?`, req.ID).Updates(conditions).Error
}

func (p *pgRepository) Delete(ctx context.Context, id string) error {
	return p.getDB(ctx).Where(`id = ?`, id).Delete(&model.Course{}).Error
}
