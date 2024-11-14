package models

type Route struct {
	Path        string `yaml:"path"`
	Method      string `yaml:"method"`
	Description string `yaml:"description"`
	Timeout     string `yaml:"timeout"`
	Retries     int    `yaml:"retries"`
}

type Service struct {
	Name   string  `yaml:"name"`
	URL    string  `yaml:"url"`
	Instance []int `yaml:"instance"`
	Routes []Route `yaml:"routes"`
}

type GatewayConfig struct {
	Name         string    `yaml:"name"`
	Version      string    `yaml:"version"`
	Description  string    `yaml:"description"`
	DefaultRoute string    `yaml:"defaultRoute"`
	GateWayInfo string 		`yaml:"gatewayInfo"`
	LoadBalancing string `yaml:"loadBalancing"`
	MainApp string `yaml:"mainApplication"`
	Services     []Service `yaml:"services"`
}
