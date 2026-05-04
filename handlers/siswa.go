package handlers

import (
	"html/template"
	"io"
	"net/http"
	"os"
)

// ==================== HOME ====================
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("Views/Home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// ==================== GET SISWA ====================
func GetSiswa(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("SUPABASE_URL") + "/rest/v1/siswa"
	key := os.Getenv("SUPABASE_KEY")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Add("apikey", key)
	req.Header.Add("Authorization", "Bearer "+key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}

// ==================== LOGIN ====================
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("Views/Login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	// POST (login)
	username := r.FormValue("username")
	password := r.FormValue("password")

	// dummy login
	if username == "admin" && password == "123" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.Write([]byte("Login gagal"))
}
