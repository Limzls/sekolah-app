package main

import (
	"log"
	"net/http"

	"sekolahapp/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// routing utama
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")

	// 🔥 TAMBAHAN PENTING
	r.HandleFunc("/siswa", handlers.GetSiswa).Methods("GET")

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
