package store

import "github.com/KoLLlaka/libraryAPI/internal/model"

type BookRepository interface {
	Add(*model.Book) error
	Find(int) ([]*model.Book, error)
	FindFrom(string, string, int) ([]*model.Book, error)
	FindByID(string) (*model.Book, error)
}
