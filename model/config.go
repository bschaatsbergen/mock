package model

type Config struct {
	Endpoints []Endpoints `yaml:"endpoints"`
}

// Endpoints
type Endpoints struct {
	Method     string `yaml:"method"`
	Statuscode int    `yaml:"statuscode"`
	Resource   string `yaml:"resource"`
	Body       string `yaml:"body"`
}
