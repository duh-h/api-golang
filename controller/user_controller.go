package controller

import (
	"api-golang/model"
	"api-golang/service"
	"api-golang/usercase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsercase usercase.UserUsecase
}

func NewUserController(usecase usercase.UserUsecase) UserController {
	return UserController{
		UserUsercase: usecase,
	}
}

func (p *UserController) CreateUser(ctx *gin.Context) {

	var user model.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user.Password = service.SHA256Encoder(user.Password)

	insertedUser, err := p.UserUsercase.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedUser)

}
