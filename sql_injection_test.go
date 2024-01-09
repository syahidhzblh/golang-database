package golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin';#"
	password := "admin"

	// For Prevent from SQL Injection
	// Replace varibale username & password with (?) for adding as parameter on QueryContext
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1;"

	fmt.Println(script)

	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		rows.Scan(&username)
		fmt.Println("Success Login", username)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Login Failed")
	}
}
