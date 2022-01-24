package usecaseUrl

import (
	"github.com/golinuxcloudnative/go-url-shortener/domain"
	"github.com/lithammer/shortuuid/v3"
)

//Service url usecase
type Service struct {
	repo domain.UrlRepository
}

//NewService create a url service
func NewService(r domain.UrlRepository) *Service {
	return &Service{repo: r}
}

//Get an url
func (s *Service) GetURL(key string) (*domain.URLShortener, error) {
	short, err := s.repo.GetURL(key)
	if err != nil {
		return nil, err
	}
	if short == nil {
		return nil, nil
	}
	return short, err
}

func (s *Service) CreateURL(url string) (*domain.URLShortener, error) {
	short, err := s.GetURL(url)
	if err != nil {
		return nil, err
	}

	if short == nil {
		id := shortuuid.New()
		short := &domain.URLShortener{URL: url, SHORT: id}
		return short, s.repo.CreateURL(*short)
	}

	return short, nil

}

func (s *Service) Update(url domain.URLShortener) error {
	return s.repo.Update(url)
}

func (s *Service) Delete(key domain.URLShortener) error {
	return s.repo.Delete(key)
}
