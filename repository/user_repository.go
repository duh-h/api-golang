package repository

import (
	"api-golang/model"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (pr *UserRepository) CreateUser(user model.User) (int, error) {
	var id int
	querry, err := pr.connection.Prepare("INSERT INTO users (user_name, email, password) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = querry.QueryRow(user.Name, user.Email, user.Password).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	querry.Close()

	return id, nil
}
