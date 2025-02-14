package config

type ServerConfig struct {
	Listen    int        `json:"listen" validate:"required"`
	Workers   int        `json:"workers"`
	Upstreams []Upstream `json:"upstreams" validate:"required,dive"`
	Headers   []Headers  `json:"headers"`
	Rules     []Rules    `json:"rules" validate:"required"`
}

type Upstream struct {
	ID  string `json:"id" validate:"required"`
	URL string `json:"url" validate:"required,url"`
}

type Headers struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type Rules struct {
	Path      string   `json:"path" validate:"required"`
	Upstreams []string `json:"upstreams" validate:"required"`
}

type RootConfigSchema struct {
	Server ServerConfig
}