package transport

import (
	"LibraryApp/internal/database"
	"LibraryApp/internal/models"
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type ViewData struct {
	Title      string
	Authors    []*models.Authors
	Books      []*models.Books
	BookAuthor string
	Issues     []*models.Issues
}

func handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func render(w http.ResponseWriter, app *database.ConnectionBase, files []string) {

	//Инициализация всех авторов
	authors, err := app.AuthorsAll()
	if err != nil {
		handleError(w, err)
		return
	}

	//Инициализация всех книг
	books, err := app.BooksAll()
	if err != nil {
		handleError(w, err)
		return
	}

	//Инициализация авторов книг в таблицу книг
	for _, book := range books {
		authorname, err := app.GetAuthorForBook(book.AuthorCode)
		if err != nil {
			handleError(w, err)
			return
		}
		book.AuthorName = authorname
	}

	//Инициализация всех выдач
	issues, err := app.GetAllIssues()
	if err != nil {
		handleError(w, err)
		return
	}

	// Инициализация форматированных дат и полей читателя
	for _, issue := range issues {

		// Блок форматирования даты
		var expired bool
		formatted_issue, err := time.Parse(time.RFC3339, issue.Issue_date)
		if err != nil {
			fmt.Println("Ошибка при парсе даты:", err)
			return
		}
		formatted_issue_in_date := formatted_issue.Format("2006-01-02")
		issue.Formatted_issue = formatted_issue_in_date

		formatted_return, err := time.Parse(time.RFC3339, issue.Return_date)
		if err != nil {
			fmt.Println("Ошибка при парсе даты:", err)
			return
		}
		formatted_return_in_date := formatted_return.Format("2006-01-02")
		issue.Formatted_return = formatted_return_in_date

		difference := time.Now().Sub(formatted_issue)
		if difference > 7*24*time.Hour {
			expired = true
		} else {
			expired = false
		}
		issue.IsExpired = expired

		// Блок получения имени и телефона читателя
		var fullname string
		var phone string
		var book_name string

		fullname, phone, book_name, err = app.GetReaderForIssue(issue.Ticket, issue.BookCode)
		if err != nil {
			fmt.Println("Ошибка при присвоении имени и телефона читателя: ", err)
			return
		}

		issue.Reader_Fullname = fullname
		issue.Reader_Phone = phone
		issue.Book_name = book_name
	}

	data := ViewData{
		Title:   "Главная",
		Authors: authors,
		Books:   books,
		Issues:  issues,
	}
	if err != nil {
		handleError(w, err)
	}

	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, data)
}
