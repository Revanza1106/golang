package main

import (
	"fmt"
	"go-postgres-crud/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fs := http.FileServer(http.Dir("build"))
	http.Handle("/", fs)
	fmt.Println("Server is available at http://localhost:8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}
