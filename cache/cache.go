package cache

import (
	"context"

	"github.com/teq-quocbang/store/presenter"
)

type ICache interface {
	Register() RegisterService
}

type RegisterService interface {
	// registerHistoriesKey is concat by accountID|semesterID|page|limit|orderBy|sortBy
	// example: 1|S0001|1|20|name|name
	CreateRegisterHistories(ctx context.Context, registerHistoriesKey string, listRegisterResp *presenter.ListRegisterResponseWrapper) error
	// registerHistoriesKey is concat by accountID|semesterID|page|limit|orderBy|sortBy
	// example: 1|S0001|1|20|name|name
	GetRegisterHistories(ctx context.Context, registerHistoriesKey string) (*presenter.ListRegisterResponseWrapper, error)
	// registerHistoriesKey is prefix accountID*
	// example: 1*
	ClearRegisterHistories(ctx context.Context, registerHistoriesKey string) error
}
