package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db  *sql.DB
	DSN string
}

func (s *Store) Open() error {
	db, err := sql.Open("mysql", s.DSN)
	if err != nil {
		return err
	}
	s.db = db
	//defer s.db.Close()
	if err := s.db.Ping(); err != nil {
		return err
	}
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) CreateRoute(urlObj *URLModel) error {
	stmt := "INSERT INTO URL (long_url, short_url) VALUES (?, ?)"
	_, err := s.db.Exec(stmt, urlObj.LongURL, urlObj.ShortURL)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetLongURL(urlObj *URLModel) error {
	stmt := "SELECT long_url FROM url WHERE short_url = ?"
	err := s.db.QueryRow(stmt, urlObj.ShortURL).Scan(&urlObj.LongURL)
	if err != nil {
		return err
	}
	return nil
}
