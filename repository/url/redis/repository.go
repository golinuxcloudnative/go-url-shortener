package repositoryRedis

import (
	"github.com/go-redis/redis/v8"
	"github.com/golinuxcloudnative/go-url-shortener/domain"
	"github.com/golinuxcloudnative/go-url-shortener/usecases/url"
)

type URLRedis struct {
	rdb *redis.Client
}

func NewRepositoryRedis(rdb *redis.Client) url.Repository {
	return &URLRedis{rdb: rdb}
}

func (db *URLRedis) GetURL(key string) (*domain.URLShortener, error) {
	return nil, nil
}

func (db *URLRedis) CreateURL(url domain.URLShortener) error {
	return nil
}
func (db *URLRedis) Update(url domain.URLShortener) error {
	return nil
}
func (db *URLRedis) Delete(key domain.URLShortener) error {
	return nil
}
