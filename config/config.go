package config

type Config struct {
	DB *DB
}

type DB struct {
	Driver   string
	Address  string
	Username string
	Password string
	Database string
	Debug    bool
}
