package transport

import (
	"LibraryApp/internal/database"
	"net/http"
	"text/template"
)

func handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func render(w http.ResponseWriter, app *database.ConnectionBase, files []string) {
	authors, err := app.AuthorsAll()
	if err != nil {
		handleError(w, err)
		return
	}

	books, err := app.BooksAll()
	if err != nil {
		handleError(w, err)
		return
	}

	for _, book := range books {
		authorname, err := app.GetAuthorForBook(book.AuthorCode)
		if err != nil {
			handleError(w, err)
			return
		}
		book.AuthorName = authorname
	}

	data := ViewData{
		Title:   "Главная",
		Authors: authors,
		Books:   books,
	}
	if err != nil {
		handleError(w, err)
	}

	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, data)
}
