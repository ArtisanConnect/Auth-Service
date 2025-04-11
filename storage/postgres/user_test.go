package postgres

import (
	pb "auth/genproto/auth"
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)

func ConnectUser() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "545512", "artisanconnect")

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestRegister(t *testing.T) {
	db, err := ConnectUser()
	if err != nil {
		t.Errorf("ConnectUser() error = %v", err)
	}
	defer db.Close()

	res := pb.RegisterRequest{
		Username: "test",
		Password: "test",
		Email:    "test@test.com",
		FullName: "test",
		UserType: "test",
	}

	user := NewUserRepo(db)

	req, err := user.Register(&res)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	fmt.Println(req)
}
