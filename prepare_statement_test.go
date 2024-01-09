package golang_database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email,comment) VALUES (?,?);"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "email" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Success Insert New Comment with ID", insertId)
	}

}
