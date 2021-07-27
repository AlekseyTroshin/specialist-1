package apiserver

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr`
	LogLevel string `toml:"log_level"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config {
		BindAddr: ":8088",
		LogLevel: "debug",  
	}
}