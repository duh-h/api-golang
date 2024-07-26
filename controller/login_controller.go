package controller

import (
	"api-golang/model"
	"api-golang/usercase"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase usercase.LoginUsecase
}

func NewLoginController(usecase usercase.LoginUsecase) *LoginController {
	return &LoginController{
		LoginUsecase: usecase,
	}
}

func (p *LoginController) Login(c *gin.Context) {
	var loginRequest model.Login
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	token, err := p.LoginUsecase.Login(loginRequest)
	if err != nil {
		status := 500
		if err.Error() == "invalid credentials" || err.Error() == "cannot find user" {
			status = 401
		}

		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
