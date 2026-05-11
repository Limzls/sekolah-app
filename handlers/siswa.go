package handlers

import (
	"html/template"
	"net/http"
)

// ==================== DASHBOARD SISWA ====================
func SiswaDashboard(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("Views/Home.html")
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

// ==================== GET SISWA ====================
func GetSiswa(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Data siswa"))
}
