package main

import (
	"log"
	"net/http"

	"sekolahapp/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	
)

func main() {
	// 🔥 load file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Gagal load .env")
	}

	r := mux.NewRouter()

	// routing utama
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/siswa", handlers.GetSiswa).Methods("GET")
	r.HandleFunc("/pegawai", handlers.GetPegawai).Methods("GET")
    r.HandleFunc("/kelas", handlers.GetKelas).Methods("GET")
    r.HandleFunc("/mapel", handlers.GetMapel).Methods("GET")

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
