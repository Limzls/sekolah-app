package handlers

import (
	"html/template"
	"net/http"
)

// ==================== DASHBOARD GURU ====================
func GuruDashboard(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("Views/guru.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, nil)
}

// ==================== GET GURU ====================
func GetGuru(w http.ResponseWriter, r *http.Request) {

	fetchData("guru", w)
}