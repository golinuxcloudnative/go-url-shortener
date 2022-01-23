package repositoryRedis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/golinuxcloudnative/go-url-shortener/domain"
)

type HealthzRedis struct {
	rdb *redis.Client
}

func NewRepositoryHealthzRedis(rdb *redis.Client) domain.HealthzRepository {
	return &HealthzRedis{rdb: rdb}
}

func (db *HealthzRedis) Ping() (string, error) {
	ctx := context.Background()
	result, err := db.rdb.Ping(ctx).Result()
	if err != nil {
		return "", err
	}
	return result, err
}
