package middlewares

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// ==================== SESSION ====================
var Store = sessions.NewCookieStore([]byte("secret-key"))

// ==================== ADMIN ONLY ====================
func AdminOnly(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := Store.Get(r, "session")

		role, ok := session.Values["role"].(string)

		// cek apakah role admin
		if !ok || role != "admin" {

			http.Error(w, "Akses ditolak", http.StatusForbidden)
			return
		}

		// lanjut ke halaman admin
		next.ServeHTTP(w, r)
	}
}

// ==================== GURU ONLY ====================
func GuruOnly(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := Store.Get(r, "session")

		role, ok := session.Values["role"].(string)

		// cek apakah role guru
		if !ok || role != "guru" {

			http.Error(w, "Akses ditolak", http.StatusForbidden)
			return
		}

		// lanjut ke halaman guru
		next.ServeHTTP(w, r)
	}
}

// ==================== SISWA ONLY ====================
func SiswaOnly(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := Store.Get(r, "session")

		role, ok := session.Values["role"].(string)

		// cek apakah role siswa
		if !ok || role != "siswa" {

			http.Error(w, "Akses ditolak", http.StatusForbidden)
			return
		}

		// lanjut ke halaman siswa
		next.ServeHTTP(w, r)
	}
}

// ==================== WALI ONLY ====================
func WaliOnly(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := Store.Get(r, "session")

		role, ok := session.Values["role"].(string)

		// cek apakah role wali
		if !ok || role != "wali" {

			http.Error(w, "Akses ditolak", http.StatusForbidden)
			return
		}

		// lanjut ke halaman wali
		next.ServeHTTP(w, r)
	}
}