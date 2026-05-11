package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := "postgresql://postgres:ISI_PASSWORD_KAMU@db.ISI_HOST.supabase.co:5432/postgres?sslmode=require"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal buka koneksi:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Gagal konek ke Supabase:", err)
	}

	log.Println("✅ Berhasil connect ke Supabase!")
}