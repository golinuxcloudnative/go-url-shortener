package domain

type URLShortener struct {
	URL   string `json:"url"`
	SHORT string `json:"short"`
}

//URL repository interface
type UrlRepository interface {
	GetURL(key string) (*URLShortener, error)
	CreateURL(url URLShortener) error
	Update(url URLShortener) error
	Delete(key URLShortener) error
}

//URL usecase interface
type UrlUseCase interface {
	GetURL(key string) (*URLShortener, error)
	CreateURL(url string) (*URLShortener, error)
	Update(url URLShortener) error
	Delete(key URLShortener) error
}
