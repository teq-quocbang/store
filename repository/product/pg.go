package product

import (
	"context"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/model"
	"gorm.io/gorm"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (r *pgRepository) Create(ctx context.Context, p *model.Product) error {
	return r.getDB(ctx).Create(&p).Error
}

func (r *pgRepository) Update(ctx context.Context, p *model.Product) error {
	conditions := p.BuildUpdateFields()
	return r.getDB(ctx).Model(&model.Product{}).Where("id = ?", p.ID).Updates(conditions).Error
}

func (r *pgRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.getDB(ctx).Where("id = ?", id).Delete(&model.Product{}).Error
}
