package models

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Services struct {
	Order OrderService
	Menu  MenuService
	db    *sqlx.DB
}

type ServicesConfig func(*Services) error

func NewServices(cfgs ...ServicesConfig) (*Services, error) {
	var s Services
	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil
}

func WithSqlite(path string) ServicesConfig {
	return func(s *Services) error {
		db, err := sqlx.Open("sqlite3", path)
		if err != nil {
			log.Fatalf("Failed to open DB: %v\n", err)
		}

		s.db = db
		return nil
	}
}

func WithMenu() ServicesConfig {
	return func(s *Services) error {
		s.Menu = NewMenuService(s.db)
		return nil
	}
}

func WithOrder() ServicesConfig {
	return func(s *Services) error {
		s.Order = NewOrderService(s.db)
		return nil
	}
}

func (s *Services) Close() error {
	return s.db.Close()
}
