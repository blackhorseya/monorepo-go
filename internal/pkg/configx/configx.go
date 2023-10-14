package configx

// Config defines the config struct.
type Config struct {
	HTTP HTTP `json:"http" yaml:"http"`
	GRPC GRPC `json:"grpc" yaml:"grpc"`
}

// HTTP defines the http config struct.
type HTTP struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

// GRPC defines the grpc config struct.
type GRPC struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}
