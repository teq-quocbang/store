package class

import (
	"context"
	"fmt"

	"github.com/teq-quocbang/store/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pgRepository struct {
	getDB func(context.Context) *gorm.DB
}

func NewClassPG(getDB func(context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB: getDB}
}

func (p *pgRepository) Create(ctx context.Context, s *model.Class) error {
	return p.getDB(ctx).Create(&s).Error
}

func (p *pgRepository) GetListBySemester(ctx context.Context, semesterID string) ([]model.Class, error) {
	var class []model.Class
	err := p.getDB(ctx).Where(`semester_id = ?`, semesterID).Find(&class).Error
	return class, err
}

func (p *pgRepository) GetByID(ctx context.Context, classID string) (model.Class, error) {
	var class model.Class
	err := p.getDB(ctx).Where(`id = ?`, classID).Take(&class).Error
	return class, err
}

func (p *pgRepository) Update(ctx context.Context, req *model.Class) error {
	conditions := req.BuildUpdateFields()
	return p.getDB(ctx).Model(&model.Class{}).Where(`id = ?`, req.ID).Updates(conditions).Error
}

func (p *pgRepository) Delete(ctx context.Context, id string) error {
	return p.getDB(ctx).Where(`id = ?`, id).Delete(&model.Class{}).Error
}

func (p *pgRepository) BatchInCreMember(ctx context.Context, id string) error {
	tx := p.getDB(ctx).Begin()
	var class *model.Class
	err := tx.Model(&class).Where(`id = ?`, id).Clauses(clause.Returning{
		Columns: []clause.Column{{Name: "current_slot"}},
	}).Update("current_slot", gorm.Expr("current_slot + 1")).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if class.CurrentSlot >= class.MaxSlot {
		tx.Rollback()
		return fmt.Errorf("class was max slot")
	}
	tx.Commit()
	return err
}

func (p *pgRepository) BatchDeCreMember(ctx context.Context, id string) error {
	return p.getDB(ctx).Model(&model.Class{}).Where(`id = ?`, id).
		Update("current_slot", gorm.Expr("current_slot - 1")).Error
}
