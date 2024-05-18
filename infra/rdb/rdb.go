package rdb

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RDB interface {
	NewSession() *gorm.DB
}

type rdb struct {
	config Config
	mu     sync.RWMutex
	db     *gorm.DB
}

type Config struct {
	Driver   string
	Address  string
	Username string
	Password string
	Database string
	Debug    bool
}

func NewRDB(config Config) RDB {
	var (
		db       *gorm.DB
		err      error
		duration time.Duration = 3 * time.Second
	)

	for {
		db, err = connect(config)
		if err != nil {
			log.Printf("failed to connect to database: %v\n", err)
			time.Sleep(duration)
			continue
		}

		break
	}

	return &rdb{
		config: config,
		mu:     sync.RWMutex{},
		db:     db,
	}
}

func connect(config Config) (*gorm.DB, error) {
	var (
		dialector gorm.Dialector
		dialect   string
		host      string
		port      string
	)

	dialect = config.Driver
	add := strings.Split(config.Address, ":")
	host = add[0]
	if len(add) == 2 {
		port = add[1]
	}

	switch dialect {
	case "postgres":
		if port == "" {
			port = "5432"
		}

		args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, config.Username, config.Password, config.Database)
		dialector = postgres.Open(args)
	case "mysql":
		if port == "" {
			port = "3306"
		}

		args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci", config.Username, config.Password, config.Address, port, config.Database)
		dialector = mysql.Open(args)
	default:
		return nil, fmt.Errorf("invalid database dialect: %s", dialect)
	}

	db, err := gorm.Open(dialector, &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, err
	}

	if config.Debug {
		db = db.Debug()
	}

	return db, nil
}

func (r *rdb) NewSession() *gorm.DB {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.db.Session(&gorm.Session{})
}
