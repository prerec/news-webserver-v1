package storage

type Config struct {
	// Строка подключения к БД
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
