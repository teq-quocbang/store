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
	return r.getDB(ctx).Save(&p).Error
}

func (r *pgRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.getDB(ctx).Where("id = ?", id).Delete(&model.Product{}).Error
}

func (r *pgRepository) GetByID(ctx context.Context, id uuid.UUID) (model.Product, error) {
	product := model.Product{}
	err := r.getDB(ctx).Where("id = ?", id).Take(&product).Error
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (r *pgRepository) CreateList(ctx context.Context, ps []model.Product) error {
	return r.getDB(ctx).Create(&ps).Error
}

func (r *pgRepository) GetList(ctx context.Context) ([]model.Product, error) {
	products := []model.Product{}
	err := r.getDB(ctx).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *pgRepository) GetListByProductIDs(ctx context.Context, productIDs []uuid.UUID) ([]model.Product, error) {
	products := []model.Product{}
	err := r.getDB(ctx).Where("id in (?)", productIDs).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
