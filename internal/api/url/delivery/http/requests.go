package http

type UrlCreateRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type UrlResolveRequest struct {
	URL string `json:"url" validate:"required"`
}
