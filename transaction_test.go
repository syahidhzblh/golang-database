package golang_database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// Start Transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email,comment) VALUES (?,?);"
	// Do Transaction
	for i := 0; i < 10; i++ {
		email := "email" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment ke " + strconv.Itoa(i)

		result, err := db.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Success Insert New Comment with ID", insertId)
	}

	// Commit
	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	//Use tx.Rollback() if want to undo changes
}
