package usecaseHealthz

import "github.com/golinuxcloudnative/go-url-shortener/domain"

//Service healthz usecase
type Service struct {
	repo domain.HealthzRepository
}

//NewService create a healthz service
func NewService(r domain.HealthzRepository) domain.HealthzUseCase {
	return &Service{repo: r}
}

func (s *Service) Ping() (string, error) {
	return s.repo.Ping()
}
