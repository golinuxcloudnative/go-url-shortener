package domain

type Healthz struct {
	Status   string `json:"status"`
	Version  string `json:"version"`
	Database string `json:"db"`
}

//URL repository interface
type HealthzRepository interface {
	Ping() (string, error)
}

//URL usecase interface
type HealthzUseCase interface {
	Ping() (string, error)
}
