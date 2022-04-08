package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/KoLLlaka/libraryAPI/internal/model"
	"github.com/KoLLlaka/libraryAPI/internal/store"
	"github.com/KoLLlaka/libraryAPI/internal/templates"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	numBooks = 20
)

type server struct {
	router        *mux.Router
	logger        *logrus.Logger
	store         store.Store
	templateCashe *templates.TemplateCashe
}

func newServer(store store.Store) *server {
	templateCashe, _ := templates.NewTemplateCashe()

	s := &server{
		router:        mux.NewRouter(),
		logger:        logrus.New(),
		store:         store,
		templateCashe: templateCashe,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	fs := http.FileServer(http.Dir("./static"))
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	s.router.HandleFunc("/books", s.handleBookCreate()).Methods("POST")
	s.router.HandleFunc("/books", s.handleShowBooks()).Methods("GET")
	//s.router.HandleFunc("/books/{bbk}/", s.handleShowBooksFromBBK()).Methods("GET")
	s.router.HandleFunc("/books/{key}", s.handleShowBook()).Methods("GET")

}

func (s *server) handleBookCreate() http.HandlerFunc {
	type request struct {
		NameAuthor string `json:"name_author"`
		NameBook   string `json:"name_book"`
		Genre      string `json:"genre"`
		Year       string `json:"publication_year"`
		NumPages   int    `json:"num_pages"`
		BBK        string `json:"bbk"`
		Desc       string `json:"description_book,omitempty"`
		IsHere     bool   `json:"is_here"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		book := &model.Book{
			NameAuthor: req.NameAuthor,
			NameBook:   req.NameBook,
			Genre:      req.Genre,
			Year:       req.Year,
			NumPages:   req.NumPages,
			BBK:        req.BBK,
			Desc:       req.Desc,
			IsHere:     req.IsHere,
		}

		if err := s.store.Book().Add(book); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, book)
	}
}

func (s *server) handleShowBooks() http.HandlerFunc {
	type data struct {
		Title string
		Books []*model.Book
	}

	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		books := []*model.Book{}

		books, err := s.store.Book().Find(numBooks)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		data := data{
			Title: "Library",
			Books: books,
		}

		s.templateCashe.Render("index.html").Execute(w, &data)
		//s.respond(w, r, http.StatusOK, books)
	}
}

func (s *server) handleShowBooksFromBBK() http.HandlerFunc {
	type data struct {
		Title string
		Books []*model.Book
	}

	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		books := []*model.Book{}
		bbk := mux.Vars(r)["bbk"]

		books, err := s.store.Book().FindFrom("bbk", bbk, numBooks)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		data := data{
			Title: bbk,
			Books: books,
		}

		s.templateCashe.Render("bbk.html").Execute(w, &data)
		//s.respond(w, r, http.StatusOK, books)
	}
}

func (s *server) handleShowBook() http.HandlerFunc {
	type data struct {
		Title string
		Book  *model.Book
	}

	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		book := &model.Book{}

		book, err := s.store.Book().FindByID(mux.Vars(r)["key"])
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		data := data{
			Title: book.NameAuthor,
			Book:  book,
		}

		s.templateCashe.Render("book.html").Execute(w, &data)
		//s.respond(w, r, http.StatusOK, book)
	}
}
