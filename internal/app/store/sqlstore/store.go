package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"
	"hack/internal/app/config"
	"hack/internal/app/store"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Store struct {
	db                 *sql.DB
	fileRepository     *FileRepository
	regionRepository   *RegionRepository
	leadRepository     *LeadRepository
	compaignRepository *CompaignRepository
	telegramRepository *TelegramRepository
}

func New(config config.Config) (*Store, error) {

	db, err := sql.Open("postgres", config.URL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse DATABASE_URL '%s'", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to create connection pool '%s'", err)
	}

	// Устанавливаем параметры
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(time.Minute * time.Duration(config.ConnMaxLifetime))

	if err != nil && err != migrate.ErrNoChange && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("unable to create database '%s'", err)
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) File() store.FileRepository {
	if s.fileRepository != nil {
		return s.fileRepository
	}

	s.fileRepository = &FileRepository{
		store: s,
	}

	return s.fileRepository
}

func (s *Store) Region() store.RegionRepository {
	if s.regionRepository != nil {
		return s.regionRepository
	}

	s.regionRepository = &RegionRepository{
		store: s,
	}

	return s.regionRepository
}

func (s *Store) Lead() store.LeadRepository {
	if s.leadRepository != nil {
		return s.leadRepository
	}

	s.leadRepository = &LeadRepository{
		store: s,
	}

	return s.leadRepository
}

func (s *Store) Compaign() store.CompaignRepository {
	if s.compaignRepository != nil {
		return s.compaignRepository
	}

	s.compaignRepository = &CompaignRepository{
		store: s,
	}

	return s.compaignRepository
}

func (s *Store) Telegram() store.TelegramRepository {
	if s.telegramRepository != nil {
		return s.telegramRepository
	}

	s.telegramRepository = &TelegramRepository{
		store: s,
	}

	return s.telegramRepository
}
