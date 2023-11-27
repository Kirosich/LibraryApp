package database

import (
	"LibraryApp/internal/models"
	"database/sql"
)

type ConnectionBase struct {
	DB *sql.DB
}

func (m *ConnectionBase) AuthorsAll() ([]*models.Authors, error) {
	query := `SELECT "authorcode", "fullname" FROM authors`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var authors []*models.Authors

	for rows.Next() {
		s := &models.Authors{}
		err = rows.Scan(&s.AuthorCode, &s.Fullname)
		if err != nil {
			return nil, err
		}
		authors = append(authors, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (m *ConnectionBase) BooksAll() ([]*models.Books, error) {
	query := `SELECT "book_code", "book_name", "authorcode", "yearpub", "amount" FROM books`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*models.Books

	for rows.Next() {
		s := &models.Books{}
		err = rows.Scan(&s.BookCode, &s.Book_name, &s.AuthorCode, &s.Yearpub, &s.Amount)
		if err != nil {
			return nil, err
		}
		books = append(books, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (m *ConnectionBase) GetAuthorForBook(BookAuthorId int) (string, error) {
	query := `
		SELECT fullname
		FROM authors
		JOIN books ON authors.authorcode = books.authorcode
		WHERE books.authorcode = $1
	`

	row := m.DB.QueryRow(query, BookAuthorId)

	var authorName string
	if err := row.Scan(&authorName); err != nil {
		return "", err
	}
	return authorName, nil
}

func (m *ConnectionBase) GetAllIssues() ([]*models.Issues, error) {
	query := `SELECT "issue_code", "book_code", "issue_date", "return_date", "ticket" FROM issues`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var issues []*models.Issues

	for rows.Next() {
		s := &models.Issues{}
		err = rows.Scan(&s.Issue_code, &s.BookCode, &s.Issue_date, &s.Return_date, &s.Ticket)
		if err != nil {
			return nil, err
		}
		issues = append(issues, s)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return issues, nil

}

func (m *ConnectionBase) GetReaderForIssue(Ticket int, book_code int) (string, string, string, error) {
	query := `
		SELECT fullname, phone
		FROM Readers
		WHERE Ticket = $1
	`

	row := m.DB.QueryRow(query, Ticket)

	var Fullname string
	var Phone string

	if err := row.Scan(&Fullname, &Phone); err != nil {
		return "", "", "", err
	}

	query = `
		SELECT book_name
		FROM books
		WHERE book_code = $1
	`

	row = m.DB.QueryRow(query, book_code)

	var Book_Name string
	if err := row.Scan(&Book_Name); err != nil {
		return "", "", "", err
	}
	return Fullname, Phone, Book_Name, nil
}

func (m *ConnectionBase) GetAllReaders() ([]*models.Readers, error) {
	query := `
		SELECT * FROM readers
	`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var readers []*models.Readers

	for rows.Next() {
		s := &models.Readers{}
		err = rows.Scan(&s.Ticket, &s.Fullname, &s.Adress, &s.Phone)
		if err != nil {
			return nil, err
		}
		readers = append(readers, s)

	}

	return readers, nil
}
