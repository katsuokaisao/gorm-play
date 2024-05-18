package config

type Config struct {
	DB *DB
}

type DB struct {
	Driver   string `env:"DB_DRIVER"`
	Address  string `env:"DB_ADDRESS"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Database string `env:"DB_DATABASE"`
	Debug    bool   `env:"DB_DEBUG"`
}
