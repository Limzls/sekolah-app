package handlers

import "net/http"

func GetKelas(w http.ResponseWriter, r *http.Request) {
	fetchData("kelas", w)
}
