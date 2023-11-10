package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreatAccount(*Account) error
	DelteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disabled"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil

}
