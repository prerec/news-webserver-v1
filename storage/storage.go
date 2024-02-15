package storage

import (
	"database/sql"
	"log"
	
	_ "github.com/lib/pq"
)

// Instance of storage
type Storage struct {
	config                 *Config
	db                     *sql.DB
	newsResponseRepository *NewsResponseRepository
}

// Storage constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open установка соединения с БД
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connection successfully")
	return nil
}

// Close закрывает соединение с БД
func (storage *Storage) Close() {
	storage.db.Close()
}

// Public Repo for News
func (s *Storage) News() *NewsResponseRepository {
	if s.newsResponseRepository != nil {
		return s.newsResponseRepository
	}
	s.newsResponseRepository = &NewsResponseRepository{
		storage: s,
	}
	return s.newsResponseRepository
}
