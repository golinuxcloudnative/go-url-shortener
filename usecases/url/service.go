package url

import "github.com/golinuxcloudnative/go-url-shortener/domain"

//Service url usecase
type Service struct {
	repo Repository
}

//NewService create a url service
func NewService(r Repository) *Service {
	return &Service{repo: r}
}

//Get an url
func (s *Service) GetURL(key string) (*domain.URLShortener, error) {
	url, err := s.repo.GetURL(key)
	if err != nil {
		return nil, err
	}
	return url, err
}

func (s *Service) CreateURL(url string) (*domain.URLShortener, error) {
	short := &domain.URLShortener{}
	return short, s.repo.CreateURL(*short)
}

func (s *Service) Update(url domain.URLShortener) error {
	return s.repo.Update(url)
}

func (s *Service) Delete(key domain.URLShortener) error {
	return s.repo.Delete(key)
}
