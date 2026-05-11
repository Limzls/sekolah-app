package handlers

import "net/http"

func GetMapel(w http.ResponseWriter, r *http.Request) {
	fetchData("mapel", w)
}