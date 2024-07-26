package repository

import (
	"api-golang/model"
	"database/sql"
)

type LoginRepository struct {
	connection *sql.DB
}

func NewLoginRepository(connection *sql.DB) LoginRepository {
	return LoginRepository{
		connection: connection,
	}
}

func (pr *LoginRepository) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	query, err := pr.connection.Prepare("SELECT id, user_name, email, password FROM users WHERE email = $1")
	if err != nil {
		return user, err
	}
	defer query.Close()

	err = query.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
