package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('hisbul', 'Hisbul');"
	_, err := db.ExecContext(ctx, script) //ExecContext can use for INSERT, UPDATE, DELETE (DML)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer;"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
	}
}

func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer;"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id, name   string
			email      sql.NullString
			balance    int32
			rating     float64
			created_at time.Time
			birth_date sql.NullTime
			married    bool
		)

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("===========")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid { //Check for Null Data
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birth_date.Valid {
			fmt.Println("Birth Date:", birth_date.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", created_at)
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "syahid"
	password := "syahid"

	script := "INSERT INTO user(username,password) VALUES (?,?)"
	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert new user")
}
