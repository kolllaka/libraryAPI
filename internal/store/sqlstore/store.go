package sqlstore

import (
	"database/sql"

	"github.com/KoLLlaka/libraryAPI/internal/store"
	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db             *sql.DB
	bookRepository *BookRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Book() store.BookRepository {
	if s.bookRepository != nil {
		return s.bookRepository
	}

	s.bookRepository = &BookRepository{
		store: s,
	}

	return s.bookRepository
}
