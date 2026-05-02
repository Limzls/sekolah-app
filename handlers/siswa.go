package handlers

import (
    "html/template"
    "net/http"

    "sekolahapp/config"
    "sekolahapp/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("Views/Home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func GetSiswa(w http.ResponseWriter, r *http.Request) {
    rows, err := config.DB.Query("SELECT * FROM Siswa")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var siswa []models.Siswa

    for rows.Next() {
        var s models.Siswa
        rows.Scan(&s.NIS, &s.NamaDepan, &s.NamaBelakang, &s.Alamat)
        siswa = append(siswa, s)
    }

    tmpl := template.Must(template.ParseFiles("templates/siswa.html"))
    tmpl.Execute(w, siswa)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("Views/Login.html")
		tmpl.Execute(w, nil)
		return
	}

	// POST (saat login dikirim)
	username := r.FormValue("username")
	password := r.FormValue("password")

	// sementara (dummy login)
	if username == "admin" && password == "123" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.Write([]byte("Login gagal"))
}