package transport

import (
	"LibraryApp/internal/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

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
		"web/ui/html/books-addbutton.partial.tmpl",
	}

	render(w, app, files)
}

func BooksAmount(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	app := &database.ConnectionBase{
		DB: db,
	}
	bookCode := r.URL.Query().Get("bookCode")
	action := r.URL.Query().Get("action")

	var updateQuery string
	switch action {
	case "increment":
		updateQuery = fmt.Sprintf("UPDATE books SET amount = amount + 1 WHERE book_code = '%s'", bookCode)
	case "decrement":
		updateQuery = fmt.Sprintf("UPDATE books SET amount = amount - 1 WHERE book_code = '%s'", bookCode)
	default:
		http.Error(w, "Некорректное действие", http.StatusBadRequest)
		return
	}

	var amount int

	_, err := app.DB.Exec(updateQuery)
	if err != nil {
		handleError(w, err)
	}

	err = app.DB.QueryRow("SELECT amount FROM books WHERE book_code = $1", bookCode).Scan(&amount)
	if err != nil {
		handleError(w, err)
	}

	jsonResponse := map[string]interface{}{"newAmount": amount}
	jsonResponseString, err := json.Marshal(jsonResponse)
	if err != nil {
		log.Println("Ошибка при маршалинге JSON:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponseString)
}

func AddBook(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	app := &database.ConnectionBase{
		DB: db,
	}

	files := []string{
		"web/ui/html/booksadd.page.tmpl",
		"web/ui/html/base.layout.tmpl",
		"web/ui/html/books-addbutton.partial.tmpl",
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
		} else {
			fmt.Fprintln(w, "Книга успешно добавлена!")
		}
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
		"web/ui/html/issuetable-buttons.partial.tmpl",
	}

	render(w, app, files)
}

func IssuetableExpired(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	app := &database.ConnectionBase{
		DB: db,
	}
	files := []string{
		"web/ui/html/issuetable-expired.page.tmpl",
		"web/ui/html/base.layout.tmpl",
		"web/ui/html/issuetable-buttons.partial.tmpl",
	}

	render(w, app, files)
}

func IssueNewMember(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	app := &database.ConnectionBase{
		DB: db,
	}
	files := []string{
		"web/ui/html/issuetable-newmember.page.tmpl",
		"web/ui/html/base.layout.tmpl",
		"web/ui/html/issuetable-buttons.partial.tmpl",
	}
	render(w, app, files)
}

func IssueNewMemberReady(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	app := &database.ConnectionBase{
		DB: db,
	}
	files := []string{
		"web/ui/html/issuetable-newmember.page.tmpl",
		"web/ui/html/base.layout.tmpl",
		"web/ui/html/issuetable-buttons.partial.tmpl",
	}
	if r.Method == http.MethodPost {
		userName := r.FormValue("user_fio")
		userAdress := r.FormValue("user_adress")
		userPhone := r.FormValue("user_phone")

		_, err := app.DB.Exec("INSERT INTO readers (fullname, adress, phone) VALUES ($1,$2,$3)", userName, userAdress, userPhone)
		if err != nil {
			handleError(w, err)
		} else {
			fmt.Fprintln(w, "Пользователь успешно добавлен!")
		}
	} else {
		render(w, app, files)
	}
}

func IssueAllMembers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	app := &database.ConnectionBase{
		DB: db,
	}
	files := []string{
		"web/ui/html/issuetable-allmembers.page.tmpl",
		"web/ui/html/base.layout.tmpl",
		"web/ui/html/issuetable-buttons.partial.tmpl",
	}
	render(w, app, files)
}

func IssueGiveBook(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	app := &database.ConnectionBase{
		DB: db,
	}
	files := []string{
		"web/ui/html/issuetable-givebook.page.tmpl",
		"web/ui/html/base.layout.tmpl",
		"web/ui/html/issuetable-buttons.partial.tmpl",
	}

	render(w, app, files)
}

func IssueGiveBookReady(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	app := &database.ConnectionBase{
		DB: db,
	}
	files := []string{
		"web/ui/html/issuetable-givebook.page.tmpl",
		"web/ui/html/base.layout.tmpl",
		"web/ui/html/issuetable-buttons.partial.tmpl",
	}
	if r.Method == http.MethodPost {
		userTicket := r.FormValue("user_ticket")
		bookCode := r.FormValue("book_code")
		TimeNow := time.Now().Format("2006-01-02 15:04:05")
		TimeAfter := time.Now().Add(time.Hour * 24 * 7).Format("2006-01-02 15:04:05")
		fmt.Println(bookCode)
		_, err := app.DB.Exec("INSERT INTO issues (book_code, issue_date, return_date, ticket) VALUES ($1,$2,$3,$4)", bookCode, TimeNow, TimeAfter, userTicket)
		if err != nil {
			handleError(w, err)
		}
		if err == nil {
			fmt.Fprintf(w, "Книга успешна выдана!")
		}
		//Уменьшаем взятую книгу, если это возможно
		query := "UPDATE books SET amount = amount - 1 WHERE book_code = $1 AND amount > 0"
		_, err = db.Exec(query, bookCode)
		if err != nil {
			handleError(w, err)
		}

	} else {
		render(w, app, files)
	}
}
