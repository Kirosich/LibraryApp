package transport

import (
	"LibraryApp/internal/database"
	"LibraryApp/internal/models"
	"database/sql"
	"fmt"
	"net/http"
)

type ViewData struct {
	Title      string
	Authors    []*models.Authors
	Books      []*models.Books
	BookAuthor string
}

func Home(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	app := &database.ConnectionBase{
		DB: db,
	}

	files := []string{
		"web/ui/html/home.page.tmpl",
		"web/ui/html/base.layout.tmpl",
	}

	render(w, app, files)
}

func Books(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	app := &database.ConnectionBase{
		DB: db,
	}

	files := []string{
		"web/ui/html/books.page.tmpl",
		"web/ui/html/base.layout.tmpl",
	}

	render(w, app, files)
}

func AddBook(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	app := &database.ConnectionBase{
		DB: db,
	}

	files := []string{
		"web/ui/html/booksadd.page.tmpl",
		"web/ui/html/base.layout.tmpl",
	}

	render(w, app, files)
}

func AddBookReady(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	app := &database.ConnectionBase{
		DB: db,
	}
	files := []string{
		"web/ui/html/booksadd.page.tmpl",
		"web/ui/html/base.layout.tmpl",
	}
	if r.Method == http.MethodPost {
		bookName := r.FormValue("book_name")
		authorCode := r.FormValue("author_code")
		yearPub := r.FormValue("year_pub")
		amount := r.FormValue("amount")

		_, err := app.DB.Exec("INSERT INTO books (book_name, authorcode, yearpub, amount) VALUES ($1,$2,$3,$4)", bookName, authorCode, yearPub, amount)
		if err != nil {
			handleError(w, err)
		}

		fmt.Fprintln(w, "Книга успешно добавлена!")
	} else {
		render(w, app, files)
	}
}

func Issuetable(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	app := &database.ConnectionBase{
		DB: db,
	}
	files := []string{
		"web/ui/html/issuetable.page.tmpl",
		"web/ui/html/base.layout.tmpl",
	}

	render(w, app, files)
}
