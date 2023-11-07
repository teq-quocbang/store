package semester

import (
	"context"

	"github.com/teq-quocbang/store/model"
	"gorm.io/gorm"
)

type pgRepository struct {
	getDB func(context.Context) *gorm.DB
}

func NewSemesterPG(getDB func(context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB: getDB}
}

func (p *pgRepository) Create(ctx context.Context, s *model.Semester) error {
	return p.getDB(ctx).Create(&s).Error
}

func (p *pgRepository) GetListByYear(ctx context.Context, year string) ([]model.Semester, error) {
	var semesters []model.Semester
	err := p.getDB(ctx).Where(`year(start_time) = ?`, year).Order("start_time DESC").Find(&semesters).Error
	return semesters, err
}

func (p *pgRepository) GetByID(ctx context.Context, semesterID string) (model.Semester, error) {
	var semester model.Semester
	err := p.getDB(ctx).Where(`id = ?`, semesterID).Take(&semester).Error
	return semester, err
}

func (p *pgRepository) Update(ctx context.Context, req *model.Semester) error {
	conditions := req.BuildUpdateFields()
	return p.getDB(ctx).Model(&model.Semester{}).Where(`id = ?`, req.ID).Updates(conditions).Error
}

func (p *pgRepository) Delete(ctx context.Context, id string) error {
	return p.getDB(ctx).Where(`id = ?`, id).Delete(&model.Semester{}).Error
}
