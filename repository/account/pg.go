package account

import (
	"context"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/model"
)

type pgRepository struct {
	getDB func(context.Context) *gorm.DB
}

func NewAccountPG(getDB func(context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB: getDB}
}

func (p *pgRepository) CreateAccount(ctx context.Context, req *model.Account) (uuid.UUID, error) {
	if err := p.getDB(ctx).Create(req).Error; err != nil {
		return uuid.UUID{}, err
	}
	return req.ID, nil
}

func (p *pgRepository) GetAccountByUsername(ctx context.Context, username string) (*model.Account, error) {
	var account *model.Account
	if err := p.getDB(ctx).Where(`username = ?`, username).Take(&account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (p *pgRepository) GetAccountByConstraint(ctx context.Context, req *model.Account) (*model.Account, error) {
	var account *model.Account
	if err := p.getDB(ctx).Where(`username = ? or email = ?`, req.Username, req.Email).
		Take(&account).Error; err != nil {
		return nil, err
	}

	return account, nil
}

func (p *pgRepository) GetList(ctx context.Context) ([]model.Account, error) {
	var accounts []model.Account
	err := p.getDB(ctx).Find(&accounts).Error
	return accounts, err
}
