package usercase

import (
	"api-golang/model"
	"api-golang/repository"
)

type UserUsecase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUsecase {
	return UserUsecase{
		repository: repo,
	}
}

func (pu *UserUsecase) CreateUser(user model.User) (model.User, error) {

	userId, err := pu.repository.CreateUser(user)

	if err != nil {
		return model.User{}, err
	}

	user.ID = userId

	return user, nil

}
