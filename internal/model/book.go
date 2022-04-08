package model

import "time"

type Book struct {
	ID         string `json:"id"`
	NameAuthor string `json:"name_author"`
	NameBook   string `json:"name_book"`
	Genre      string `json:"genre"`
	Year       string `json:"publication_year"`
	NumPages   int    `json:"num_pages"`
	BBK        string `json:"bbk"`

	Desc    string    `json:"description_book,omitempty"`
	AddDate time.Time `json:"add_date,omitempty"`
	IsHere  bool      `json:"is_here"`
	// ReturnDate time.Time `json:"return_date,omitempty"`
	// TakeDate   time.Time `json:"take_date,omitempty"`
}