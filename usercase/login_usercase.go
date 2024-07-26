package usercase

import (
	"api-golang/model"
	"api-golang/repository"
	"api-golang/service"
	"fmt"
)

type LoginUsecase struct {
	repository repository.LoginRepository
	jwtService service.JWTService
}

func NewLoginUsecase(repo repository.LoginRepository, jwtService *service.JWTService) LoginUsecase {
	return LoginUsecase{
		repository: repo,
		jwtService: *jwtService,
	}
}

func (u *LoginUsecase) Login(loginRequest model.Login) (string, error) {
	user, err := u.repository.FindUserByEmail(loginRequest.Email)
	if err != nil {
		return "", err
	}

	if user.Password != service.SHA256Encoder(loginRequest.Password) {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := u.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
