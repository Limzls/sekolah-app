package handlers

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"os"

)

// ==================== STRUCT USER ====================
type User struct {
	IDUser   int    `json:"id_user"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// ==================== HOME ====================
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("Views/Home.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, nil)
}

// ==================== LOGIN ====================
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// ==================== TAMPILKAN HALAMAN LOGIN ====================
	if r.Method == "GET" {

		tmpl, err := template.ParseFiles("Views/Login.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	// ==================== AMBIL INPUT FORM ====================
	username := r.FormValue("username")
	password := r.FormValue("password")

	// debug terminal
	println("USERNAME :", username)
	println("PASSWORD :", password)

	// validasi kosong
	if username == "" || password == "" {
		http.Error(w, "Username dan Password wajib diisi", 400)
		return
	}

	// ==================== QUERY SUPABASE ====================
	urlAPI := os.Getenv("SUPABASE_URL") +
		"/rest/v1/users?select=*" +
		"&username=eq." + url.QueryEscape(username) +
		"&password=eq." + url.QueryEscape(password)

	key := os.Getenv("SUPABASE_KEY")

	req, err := http.NewRequest("GET", urlAPI, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	req.Header.Add("apikey", key)
	req.Header.Add("Authorization", "Bearer "+key)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer resp.Body.Close()

	// ==================== BACA RESPONSE ====================
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// debug response
	println("RESPONSE SUPABASE:")
	println(string(body))

	// cek status response
	if resp.StatusCode != 200 {
		http.Error(w, "Login gagal", 401)
		return
	}

	// ==================== UBAH JSON KE STRUCT ====================
	var users []User

	err = json.Unmarshal(body, &users)
	if err != nil {

		println("ERROR JSON:")
		println(err.Error())

		http.Error(w, "Format data salah", 500)
		return
	}

	// ==================== CEK USER ====================
	if len(users) == 0 {
		http.Error(w, "Username atau Password salah", 401)
		return
	}

	user := users[0]

	// ==================== ROLE ADMIN ====================
	if user.Role == "admin" {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	// ==================== ROLE GURU ====================
	if user.Role == "guru" {
		http.Redirect(w, r, "/guru", http.StatusSeeOther)
		return
	}

	// ==================== ROLE SISWA ====================
	if user.Role == "siswa" {
		http.Redirect(w, r, "/siswa-dashboard", http.StatusSeeOther)
		return
	}

	// ==================== ROLE WALI ====================
	if user.Role == "wali" {
		http.Redirect(w, r, "/wali-dashboard", http.StatusSeeOther)
		return
	}

	// ==================== DEFAULT ====================
	http.Error(w, "Role tidak dikenali", 401)
}