package main

import (
    "log"
    "net/http"

    //"sekolahapp/config"
    "sekolahapp/handlers"

    "github.com/gorilla/mux"
)

func main() {
	//config.ConnectDB()

	r := mux.NewRouter() // jika run tidak bisa nonaktifkan keamanan //
    
    r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
    r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
