package main

import (
	"log"
	"net/http"

	"sekolahapp/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// ==================== LOAD .ENV ====================
	err := godotenv.Load()
	if err != nil {
		log.Println("Gagal load .env")
	}

	// ==================== ROUTER ====================
	r := mux.NewRouter()

	// ==================== AUTH ====================
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")

	// ==================== DASHBOARD ====================
	r.HandleFunc("/admin", handlers.AdminDashboard).Methods("GET")
	r.HandleFunc("/guru", handlers.GuruDashboard).Methods("GET")
	r.HandleFunc("/siswa-dashboard", handlers.SiswaDashboard).Methods("GET")

	// ==================== DATA ====================

	r.HandleFunc("/pegawai", handlers.GetPegawai).Methods("GET")
	r.HandleFunc("/kelas", handlers.GetKelas).Methods("GET")
	r.HandleFunc("/mapel", handlers.GetMapel).Methods("GET")

	// ==================== DATA ROLE ====================
	r.HandleFunc("/data-admin", handlers.GetAdmin).Methods("GET")
	r.HandleFunc("/data-guru", handlers.GetGuru).Methods("GET")
	r.HandleFunc("/siswa", handlers.GetSiswa).Methods("GET")

	// ==================== SERVER ====================
	log.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", r)
}
