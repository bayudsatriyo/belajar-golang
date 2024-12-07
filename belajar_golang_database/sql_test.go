package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)


func TestExecSql(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(email, balance, rating, birth_date, married) VALUES('yanto@gmail.com', '1000000', 5.0, '2001-01-20', 1)"

	// exec untuk mengeksekusi creation data (bukan query)
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	t.Log("Data berhasil disimpan")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	script := "SELECT * FROM customer"
	// queryContext untuk mengeksekusi query data
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id, "name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	script := "SELECT email, balance, rating, birth_date, married, created_at FROM customer"
	// queryContext untuk mengeksekusi query data
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var email string
		var balance int32
		var rating float64
		var birth_date, created_at time.Time
		var married bool
		err := rows.Scan(&email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("email:", email, "balance:", balance, "rating:", rating, "birth_date:", birth_date, "married:", married, "created_at:", created_at)
	}
}

func TestQuerySqlNullableComplex(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	script := "SELECT email, balance, rating, birth_date, married, created_at FROM customer"
	// queryContext untuk mengeksekusi query data
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var email string
		var balance int32
		var rating float64
		var created_at time.Time
		var birth_date sql.NullTime
		var married bool
		err := rows.Scan(&email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("email:", email, "balance:", balance, "rating:", rating, "married:", married, "created_at:", created_at)

		if birth_date.Valid {
			fmt.Println("birth_date:", birth_date.Time)
		}
	}
}

func TestQueryLogin(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	username := "yanti"
	password := "rahasia"

	script := "SELECT username, password FROM users WHERE username = ? AND password = ?"
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username, password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println("username:", username, "password:", password)
		fmt.Println("Data ditemukan")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func TestLastInsertId (t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	email := "yanto@gmail.com"
	comment := "halo"

	script := "INSERT INTO comment(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("success insert new coment with id", insertId)
}

func TestPrepareStatemnet(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comment(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := "bayu" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke-" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("success insert new coment with id", lastInsertId)
	}
}

func TestTransaction(t *testing.T)  {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comment(email, comment) VALUES(?, ?)"

	// do transaction
	for i := 11; i < 20; i++ {
		email := "bayu" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke-" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)

		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("success insert new coment with id", lastInsertId)
	}


	err = tx.Commit()
	// jika ingin rollback
	err = tx.Rollback()

	if err != nil {
		panic(err)
	}

}