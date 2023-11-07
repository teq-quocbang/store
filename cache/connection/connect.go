package connection

import (
	"context"
	"fmt"

	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	"git.teqnological.asia/teq-go/teq-pkg/teqsentry"
	"github.com/redis/go-redis/v9"
	"github.com/teq-quocbang/store/cache"
)

type redisDB struct {
	redis *redis.Client
}

type RedisConfig struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

func NewRedisCache(cfg RedisConfig) cache.ICache {
	client, err := NewRedis(cfg)
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}
	return redisDB{
		redis: client,
	}
}

func NewRedis(cfg RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Address, cfg.Port),
		Password: cfg.Password,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
