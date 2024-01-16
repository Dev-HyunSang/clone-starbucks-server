package db

import (
	"errors"
	"github.com/dev-hyunsang/clone-stackbuck-backend/config"
	"github.com/redis/go-redis/v9"
)

var (
	ErrEnvSetUp        string = "not setting too Env"
	ErrConnectionRedis string = "failed to connect Redis"
)

func ConnectRedis() (*redis.Client, error) {
	dsn := config.GetDotEnv("REDIS_ADR")
	redisPw := config.GetDotEnv("REDIS_PASSWORD")

	if len(dsn) == 0 {
		return nil, errors.New(ErrEnvSetUp)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: redisPw,
	})

	return client, nil
}
