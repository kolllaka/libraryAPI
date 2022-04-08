package store

type Store interface {
	Book() BookRepository
}
