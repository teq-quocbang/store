package producer

import (
	"context"

	"github.com/teq-quocbang/store/model"
	"gorm.io/gorm"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (r *pgRepository) Create(ctx context.Context, p *model.Producer) error {
	return r.getDB(ctx).Create(&p).Error
}
