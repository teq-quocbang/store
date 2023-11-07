package register

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	"github.com/redis/go-redis/v9"
	"github.com/teq-quocbang/store/cache"
	"github.com/teq-quocbang/store/cache/database"
	"github.com/teq-quocbang/store/presenter"
	"go.uber.org/zap"
)

type cacheService struct {
	redis *redis.Client
}

func NewRedisRegisterCache(redis *redis.Client) cache.RegisterService {
	return &cacheService{
		redis: redis,
	}
}

func (c *cacheService) CreateRegisterHistories(ctx context.Context, registerHistoriesKey string, listRegisterResp *presenter.ListRegisterResponseWrapper) error {
	// set redis database
	c.redis.Conn().Select(ctx, int(database.Database_Cache))

	values, err := json.Marshal(listRegisterResp)
	if err != nil {
		return fmt.Errorf("failed to marshal register histories, error: %v", err)
	}

	// save value
	_, err = c.redis.Set(ctx, registerHistoriesKey, values, time.Hour*1000).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *cacheService) GetRegisterHistories(ctx context.Context, registerHistoriesKey string) (*presenter.ListRegisterResponseWrapper, error) {
	// select database
	c.redis.Conn().Select(ctx, int(database.Database_Cache))

	// get database from cache
	value, err := c.redis.Get(ctx, registerHistoriesKey).Result()
	if err != nil {
		return nil, err
	}

	registerHistories := &presenter.ListRegisterResponseWrapper{}
	if err := json.Unmarshal([]byte(value), registerHistories); err != nil {
		return nil, err
	}

	return registerHistories, nil
}

func (c *cacheService) ClearRegisterHistories(ctx context.Context, registerHistoriesKey string) error {
	// select cache database
	c.redis.Conn().Select(ctx, int(database.Database_Cache))

	// remove from cache
	iter := c.redis.Scan(ctx, 0, fmt.Sprintf("%s*", registerHistoriesKey), 0).Iterator()
	for iter.Next(ctx) {
		if err := c.redis.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
		teqlogger.GetLogger().Info("Cache Tracing, status: REMOVE", zap.String("key", iter.Val()))
	}

	return nil
}
