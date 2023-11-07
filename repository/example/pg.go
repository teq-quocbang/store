package example

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/teq-quocbang/store/codetype"
	"github.com/teq-quocbang/store/model"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (p *pgRepository) Create(ctx context.Context, data *model.Example) error {
	return p.getDB(ctx).Create(data).Error
}

func (p *pgRepository) CreateList(ctx context.Context, data []model.Example) error {
	return p.getDB(ctx).Create(data).Error
}

func (p *pgRepository) CreateOrUpdate(ctx context.Context, data *model.Example) error {
	return p.getDB(ctx).
		Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns([]string{"id"}),
		}).
		Create(data).
		Error
}

func (p *pgRepository) CreateOrUpdateList(ctx context.Context, data []model.Example) error {
	return p.getDB(ctx).
		Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns([]string{"id"}),
		}).
		Create(data).
		Error
}

func (p *pgRepository) Update(ctx context.Context, data *model.Example) error {
	return p.getDB(ctx).Save(data).Error
}

func (p *pgRepository) UpdateList(ctx context.Context, data []model.Example) error {
	return p.getDB(ctx).Save(data).Error
}

func (p *pgRepository) GetByID(ctx context.Context, id int64) (*model.Example, error) {
	var example model.Example

	err := p.getDB(ctx).
		Where("id = ?", id).
		First(&example).
		Error
	if err != nil {
		return nil, err
	}

	return &example, nil
}

func (p *pgRepository) GetByInterface(ctx context.Context, itf interface{}) (*model.Example, error) {
	if itf == nil {
		return nil, errors.New("Nil interface")
	}

	var data model.Example

	err := p.getDB(ctx).Where(itf).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (p *pgRepository) GetListByInterface(ctx context.Context, itf interface{}) ([]model.Example, error) {
	if itf == nil {
		return nil, errors.New("Nil interface")
	}

	var data []model.Example

	err := p.getDB(ctx).Where(itf).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *pgRepository) Delete(ctx context.Context, data *model.Example, unscoped bool) error {
	db := p.getDB(ctx)

	if unscoped {
		db = db.Unscoped()
	}

	return db.Delete(data).Error
}

func (p *pgRepository) DeleteList(ctx context.Context, data []model.Example, unscoped bool) error {
	db := p.getDB(ctx)

	if unscoped {
		db = db.Unscoped()
	}

	return db.Delete(data).Error
}

func (p *pgRepository) GetAll(ctx context.Context, unscoped bool) ([]model.Example, error) {
	var (
		data []model.Example
		db   = p.getDB(ctx)
	)

	if unscoped {
		db = db.Unscoped()
	}

	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (p *pgRepository) GetList(
	ctx context.Context,
	search string,
	paginator codetype.Paginator,
	conditions interface{},
	order []string,
) ([]model.Example, int64, error) {
	var (
		db     = p.getDB(ctx).Model(&model.Example{})
		data   = make([]model.Example, 0)
		total  int64
		offset int
	)

	if conditions != nil {
		db = db.Where(conditions)
	}

	if search != "" {
		db.Where("name LIKE ?", "%"+search+"%")
	}

	for i := range order {
		db = db.Order(order[i])
	}

	if paginator.Page != 1 {
		offset = paginator.Limit * (paginator.Page - 1)
	}

	if paginator.Limit != -1 {
		err := db.Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	}

	err := db.Limit(paginator.Limit).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	if paginator.Limit == -1 {
		total = int64(len(data))
	}

	return data, total, nil
}
