package main

import (
	"LibraryApp/internal/transport"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	// Replace the connection string with your PostgreSQL server details
	connStr := "user=postgres dbname=LibraryApp password=272731 sslmode=disable"

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		transport.Home(w, r, db)
	})
	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		transport.Books(w, r, db)
	})
	router.HandleFunc("/changeAmount", func(w http.ResponseWriter, r *http.Request) {
		transport.BooksAmount(w, r, db)
	})
	router.HandleFunc("/books/add", func(w http.ResponseWriter, r *http.Request) {
		transport.AddBook(w, r, db)
	})
	router.HandleFunc("/books/add/ready", func(w http.ResponseWriter, r *http.Request) {
		transport.AddBookReady(w, r, db)
	})
	router.HandleFunc("/issuetable", func(w http.ResponseWriter, r *http.Request) {
		transport.Issuetable(w, r, db)
	})
	router.HandleFunc("/issuetable/expired", func(w http.ResponseWriter, r *http.Request) {
		transport.IssuetableExpired(w, r, db)
	})
	router.HandleFunc("/issuetable/newmember", func(w http.ResponseWriter, r *http.Request) {
		transport.IssueNewMember(w, r, db)
	})
	router.HandleFunc("/issuetable/newmember/ready", func(w http.ResponseWriter, r *http.Request) {
		transport.IssueNewMemberReady(w, r, db)
	})
	router.HandleFunc("/issuetable/allmembers", func(w http.ResponseWriter, r *http.Request) {
		transport.IssueAllMembers(w, r, db)
	})
	router.HandleFunc("/issuetable/givebook", func(w http.ResponseWriter, r *http.Request) {
		transport.IssueGiveBook(w, r, db)
	})
	router.HandleFunc("/issuetable/givebook/ready", func(w http.ResponseWriter, r *http.Request) {
		transport.IssueGiveBookReady(w, r, db)
	})

	// Инициализируем FileServer, он будет обрабатывать
	// HTTP-запросы к статическим файлам из папки "./ui/static".
	// Обратите внимание, что переданный в функцию http.Dir путь
	// является относительным корневой папке проекта
	fileServer := http.FileServer(http.Dir("/home/kirosi/Desktop/code/LibraryApp/web/ui/static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	fmt.Println("Server is listening...")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error")
	}
}
