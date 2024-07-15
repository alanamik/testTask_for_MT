package redis

import (
	"context"
	"errors"
	"mt/internal/config"

	"github.com/redis/go-redis/v9"
)

var (
	ErrInternalServiceRedisDB = errors.New("failed request to RedisDB")
)

type RedisClient struct {
	Client *redis.Client
}

func NewClient(con *config.Config) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     con.Redis.Address,
		Password: "password",
		DB:       con.Redis.DB,
	})

	return &RedisClient{
		Client: client,
	}
}

func (r *RedisClient) AddInCache(ctx context.Context, str string, cypher string) error {
	err := r.Client.Set(ctx, str, cypher, 0).Err()
	return err
}

func (r *RedisClient) GetFromCache(ctx context.Context, str string) (string, error) {
	cypher, err := r.Client.Get(ctx, str).Result()
	return cypher, err
}
