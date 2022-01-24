package repositoryRedis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/golinuxcloudnative/go-url-shortener/domain"
)

type URLRedis struct {
	rdb *redis.Client
}

func NewRepositoryURLRedis(rdb *redis.Client) domain.UrlRepository {
	return &URLRedis{rdb: rdb}
}

func (db *URLRedis) GetURL(key string) (*domain.URLShortener, error) {
	ctx := context.Background()
	result, err := db.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &domain.URLShortener{URL: result, SHORT: key}, nil
}

func (db *URLRedis) CreateURL(url domain.URLShortener) error {
	ctx := context.Background()
	err := db.rdb.Set(ctx, url.SHORT, url.URL, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
func (db *URLRedis) Update(url domain.URLShortener) error {
	return nil
}
func (db *URLRedis) Delete(key domain.URLShortener) error {
	return nil
}
