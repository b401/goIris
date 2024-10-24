package iris

import "sync"

type Config struct {
	BaseUrl   string
	AuthToken string
}

var (
	once     sync.Once
	instance *Config
)

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}
