package handlers

import (
	"html/template"
	"io"
	"net/http"
	"os"
)

// ==================== HELPER ====================
func fetchData(table string, w http.ResponseWriter) {
	url := os.Getenv("SUPABASE_URL") + "/rest/v1/" + table
	key := os.Getenv("SUPABASE_KEY")

	if key == "" {
		http.Error(w, "API KEY kosong (.env tidak terbaca)", 500)
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	req.Header.Add("apikey", key)
	req.Header.Add("Authorization", "Bearer "+key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}

// ==================== HOME ====================
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("Views/Home.html")
	tmpl.Execute(w, nil)
}

// ==================== DATA ====================
func GetSiswa(w http.ResponseWriter, r *http.Request) {
	fetchData("siswa", w)
}

func GetPegawai(w http.ResponseWriter, r *http.Request) {
	fetchData("pegawai", w)
}

func GetMapel(w http.ResponseWriter, r *http.Request) {
	fetchData("mapel", w)
}

func GetKelas(w http.ResponseWriter, r *http.Request) {
	fetchData("kelas", w)
}

// ==================== LOGIN ====================
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("Views/Login.html")
		tmpl.Execute(w, nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "123" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.Write([]byte("Login gagal"))
}
