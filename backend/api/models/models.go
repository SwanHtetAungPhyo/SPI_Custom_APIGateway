package models

type Route struct {
	Path           []string  `yaml:"path" json:"path,omitempty"`
	Method         []string  `yaml:"method" json:"method,omitempty"`
	Description    string    `yaml:"description" json:"description,omitempty"`
	Timeout        string    `yaml:"timeout" json:"timeout,omitempty"`
	Retries        int       `yaml:"retries" json:"retries,omitempty"`
	GeneratedRoute *[]string `yaml:"generatedRoute" json:"generated-route,omitempty"`
}

type Service struct {
	Name     string   `yaml:"name"`
	URL      string   `yaml:"url"`
	Leader   string   `yaml:"leader"`
	Instance []int    `yaml:"instance"`
	Routes   [1]Route `yaml:"routes"`
}

type GatewayConfig struct {
	Name          string    `yaml:"name"`
	Version       string    `yaml:"version"`
	Description   string    `yaml:"description"`
	DefaultRoute  string    `yaml:"defaultRoute"`
	GateWayInfo   string    `yaml:"gatewayInfo"`
	UserName      string    `yaml:"userName"`
	Password      string    `yaml:"password"`
	JwtKey        string    `yaml:"jwtKey"`
	BlackSpace    string    `yaml:"blackSpaceIP"`
	LoadBalancing string    `yaml:"loadBalancing"`
	MainApp       string    `yaml:"mainApplication"`
	Services      []Service `yaml:"services"`
}
