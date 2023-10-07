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
