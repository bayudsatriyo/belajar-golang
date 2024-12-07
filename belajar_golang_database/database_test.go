package belajargolangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {
	
}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/belajar_golang")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	
	// gunakan dbnya di sini
}