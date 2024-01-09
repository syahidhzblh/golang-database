package repository

import (
	"context"
	"fmt"
	golang_database "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindComment(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())
	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())
	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
