package account

import (
	"context"

	"gorm.io/gorm"

	"github.com/teq-quocbang/store/model"
)

type pgRepository struct {
	getDB func(context.Context) *gorm.DB
}

func NewAccountPG(getDB func(context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB: getDB}
}

func (p *pgRepository) CreateAccount(ctx context.Context, req *model.Account) (uint, error) {
	if err := p.getDB(ctx).Create(req).Error; err != nil {
		return 0, err
	}
	return req.ID, nil
}

func (p *pgRepository) GetAccountByID(ctx context.Context, studentID uint) (*model.Account, error) {
	var account *model.Account
	if err := p.getDB(ctx).Where(`id = ?`, studentID).Take(&account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (p *pgRepository) CreateVerifyAccount(ctx context.Context, req *model.AccountVerify) error {
	return nil
}

func (p *pgRepository) GetVerifyAccountByID(ctx context.Context, studentID uint) (*model.AccountVerify, error) {
	return nil, nil
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
