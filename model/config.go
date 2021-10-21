package model

type Config struct {
	Port      string      `yaml:"Port"`
	Endpoints []Endpoints `yaml:"Endpoints"`
}

// Endpoints
type Endpoints struct {
	Method     string `yaml:"Method"`
	Statuscode int    `yaml:"StatusCode"`
	Resource   string `yaml:"Resource"`
	Response   string `yaml:"Response"`
}
