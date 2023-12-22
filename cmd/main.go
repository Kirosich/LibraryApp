package main

import (
	"LibraryApp/internal/transport"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	globalContext    = context.Background()
	globalContextMux sync.Mutex
	isLoggedOut      bool
)

func setGlobalContextValue(key, value string) {
	globalContextMux.Lock()
	defer globalContextMux.Unlock()

	globalContext = context.WithValue(globalContext, key, value)
}

func getGlobalContextValue(key string) interface{} {
	globalContextMux.Lock()
	defer globalContextMux.Unlock()

	return globalContext.Value(key)
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {
	setGlobalContextValue("username", "")
	isLoggedOut = true

	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func basicAuthMiddleware(r *mux.Router, db *sql.DB) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			user, pass, ok := req.BasicAuth()

			if !ok || !checkCredentials(user, pass, db) || isLoggedOut {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				isLoggedOut = false
				return
			}

			// Сохраняем значение "username" в глобальном контексте
			setGlobalContextValue("username", user)

			next.ServeHTTP(w, req)
		})
	}
}

func checkCredentials(username, password string, db *sql.DB) bool {
	var validUser string
	var validPassword string

	query := `SELECT username, password FROM users WHERE id = 1`

	row := db.QueryRow(query)

	err := row.Scan(&validUser, &validPassword)
	if err != nil {
		fmt.Println("Uncorrect user or password")
	}

	return username == validUser && password == validPassword
}
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем значение "username" из глобального контекста
		username := getGlobalContextValue("username")

		if username == nil || username.(string) == "" {
			fmt.Println("Redirecting to /login")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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

	router.HandleFunc("/logout", logoutHandler).Methods("GET")

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

	router.Use(basicAuthMiddleware(router, db))
	router.Use(authMiddleware)
	// Инициализируем FileServer, он будет обрабатывать
	// HTTP-запросы к статическим файлам из папки "./ui/static".
	fileServer := http.FileServer(http.Dir("/home/kirosi/Desktop/code/LibraryApp/web/ui/static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	fmt.Println("Server is listening...")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error")
	}
}
