package main

import (
	"LibraryApp/internal/transport"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", transport.Home)

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error")
	}
}
