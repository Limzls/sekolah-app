package handlers

import (
	"io"
	"net/http"
	"os"
)

func fetchData(table string, w http.ResponseWriter) {

	url := os.Getenv("SUPABASE_URL") + "/rest/v1/" + table
	key := os.Getenv("SUPABASE_KEY")

	// cek API key
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