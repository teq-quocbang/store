package account

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/payload"
	"github.com/teq-quocbang/store/presenter"
	"github.com/teq-quocbang/store/util/hashing"
	"github.com/teq-quocbang/store/util/myerror"
	"github.com/teq-quocbang/store/util/token"
	"gorm.io/gorm"
)

func (u *UseCase) SignUp(ctx context.Context, req *payload.SignUpRequest) (*presenter.AccountResponseWrapper, error) {
	// validate check
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrAccountInvalidParam(err.Error())
	}

	// check unique constraint
	account, err := u.Account.GetAccountByConstraint(ctx, &model.Account{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, myerror.ErrAccountGet(err)
	}

	// check whether constraint is existed
	if account != nil {
		return nil, myerror.ErrAccountConflictUniqueConstraint("Username or Email was registered")
	}

	// create account
	hashPassword, err := hashing.ToHashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	createAccountRequest := &model.Account{
		ID:           uuid.New(),
		Username:     req.Username,
		HashPassword: hashPassword,
		Email:        req.Email,
	}

	ID, err := u.Account.CreateAccount(ctx, createAccountRequest)
	if err != nil {
		return nil, myerror.ErrAccountCreate(err)
	}

	return &presenter.AccountResponseWrapper{Account: &model.Account{
		ID: ID,
	}}, nil
}

func (p *UseCase) Login(ctx context.Context, req *payload.LoginRequest) (*presenter.AccountLoginResponseWrapper, error) {
	// validate check
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrAccountInvalidParam(err.Error())
	}

	// get account
	account, err := p.Account.GetAccountByUsername(ctx, req.Username)
	if err != nil {
		return nil, myerror.ErrAccountGet(err)
	}

	// compare password
	if err := hashing.CompareHashPassword(req.Password, account.HashPassword); err != nil {
		return nil, myerror.ErrAccountComparePassword(err)
	}

	jwt := token.JWT{
		SecretKey: p.Config.TokenSecretKey,
		User: token.UserInfo{
			ID:       account.ID,
			Username: account.Username,
			Email:    account.Email,
		},
		TokenLifeTime: time.Duration(p.Config.AccessTokenDuration * int64(time.Second)),
	}
	// generate access token
	accessToken, _, err := jwt.GenerateToken()
	if err != nil {
		return nil, myerror.ErrAccountGenerateToken(err)
	}

	// generate refresh token
	jwt.TokenLifeTime = time.Duration(p.Config.RefreshTokenDuration * int64(time.Second))
	refreshToken, _, err := jwt.GenerateToken()
	if err != nil {
		return nil, myerror.ErrAccountGenerateToken(err)
	}

	return &presenter.AccountLoginResponseWrapper{
		Data: presenter.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
