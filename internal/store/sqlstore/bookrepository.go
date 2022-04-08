package sqlstore

import (
	"database/sql"
	"fmt"

	"github.com/KoLLlaka/libraryAPI/internal/model"
	"github.com/KoLLlaka/libraryAPI/internal/store"
	"github.com/google/uuid"
)

const (
	booksTable = "books"
)

type BookRepository struct {
	store *Store
}

func (r *BookRepository) Add(book *model.Book) error {
	id := uuid.New()

	// INSERT INTO books (id, name, bbk, is_here) VALUES (UUID_TO_BIN(UUID()), "BOOK1", 123, true)
	stmt := fmt.Sprintf(`
	INSERT INTO
		%s
		(id, name_author, name_book, genre, publication_year, num_pages, bbk, description_book, is_here)
	VALUES
		(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, booksTable)

	book.ID = id.String()

	return r.store.db.QueryRow(
		stmt,
		id,
		book.NameAuthor,
		book.NameBook,
		book.Genre,
		book.Year,
		book.NumPages,
		book.BBK,
		book.Desc,
		book.IsHere,
	).Err()
}

func (r *BookRepository) Find(num int) ([]*model.Book, error) {
	books := []*model.Book{}

	// SELECT id, name, bbk, add_date, is_here FROM books ORDER BY add_date DESC LIMIT 0, 2;
	stmt := fmt.Sprintf(`
	SELECT 
		id, name_author, name_book, genre, publication_year, num_pages, bbk, description_book, add_date, is_here
	FROM 
		%s
	ORDER BY add_date DESC
	LIMIT 0, ?;
	`, booksTable)

	rows, err := r.store.db.Query(stmt, num)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		book := &model.Book{}

		if err := rows.Scan(
			&book.ID,
			&book.NameAuthor,
			&book.NameBook,
			&book.Genre,
			&book.Year,
			&book.NumPages,
			&book.BBK,
			&book.Desc,
			&book.AddDate,
			&book.IsHere,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (r *BookRepository) FindFrom(cat, catVal string, num int) ([]*model.Book, error) {
	books := []*model.Book{}

	// SELECT id, name, bbk, add_date, is_here FROM books WHERE bbk = 123 ORDER BY add_date DESC LIMIT 0, 2;
	stmt := fmt.Sprintf(`
	SELECT 
		id, name_author, name_book, genre, publication_year, num_pages, bbk, description_book, add_date, is_here
	FROM 
		%s
	WHERE
		%s = ?
	ORDER BY add_date DESC
	LIMIT
		0, ?;
	`, booksTable, cat)

	rows, err := r.store.db.Query(stmt, catVal, num)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		book := &model.Book{}

		if err := rows.Scan(
			&book.ID,
			&book.NameAuthor,
			&book.NameBook,
			&book.Genre,
			&book.Year,
			&book.NumPages,
			&book.BBK,
			&book.Desc,
			&book.AddDate,
			&book.IsHere,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (r *BookRepository) FindByID(bookID string) (*model.Book, error) {
	book := &model.Book{}

	// SELECT id, name, bbk, add_date, is_here FROM books WHERE id = "5efefebc-4d10-4217-96dd-3ccb7db32f82"
	stmt := fmt.Sprintf(`
	SELECT 
		id, name_author, name_book, genre, publication_year, num_pages, bbk, description_book, add_date, is_here
	FROM 
		%s 
	WHERE 
		id = ?
	`, booksTable)

	if err := r.store.db.QueryRow(stmt, bookID).Scan(
		&book.ID,
		&book.NameAuthor,
		&book.NameBook,
		&book.Genre,
		&book.Year,
		&book.NumPages,
		&book.BBK,
		&book.Desc,
		&book.AddDate,
		&book.IsHere,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return book, nil
}
