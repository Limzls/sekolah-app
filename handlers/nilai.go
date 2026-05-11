package handlers

import (
	"io"
	"net/http"
	"os"
)

func GetNilai(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("SUPABASE_URL") + "/rest/v1/nilai"
	key := os.Getenv("SUPABASE_KEY")

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