package url

import "github.com/golinuxcloudnative/go-url-shortener/domain"

//URL repository nterface
type Repository interface {
	GetURL(key string) (*domain.URLShortener, error)
	CreateURL(url domain.URLShortener) error
	Update(url domain.URLShortener) error
	Delete(key domain.URLShortener) error
}

//URL usecase interface
type UseCase interface {
	GetURL(key string) (*domain.URLShortener, error)
	CreateURL(url string) (*domain.URLShortener, error)
	Update(url domain.URLShortener) error
	Delete(key domain.URLShortener) error
}
