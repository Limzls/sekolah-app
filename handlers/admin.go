package handlers

import (
	"html/template"
	"net/http"
)

// ==================== DASHBOARD ADMIN ====================
func AdminDashboard(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("Views/admin.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, nil)
}

// ==================== GET ADMIN ====================
func GetAdmin(w http.ResponseWriter, r *http.Request) {

	fetchData("users", w)
}