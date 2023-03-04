package controller

import (
	"goapi/auth"
	"goapi/controller/dto"
	"goapi/model"
	"goapi/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
}

type authController struct {
	UserService service.UserService
}

func NewAuthController(userService service.UserService) AuthController {
	return &authController{UserService: userService}
}

func (ctr *authController) SignIn(c *gin.Context) {
	var signInRequest dto.SignInRequest
	if err := c.Bind(&signInRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := ctr.UserService.FindByEmail(signInRequest.Email)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := user.ComparePassword(signInRequest.Password); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}

func (controller *authController) SignUp(c *gin.Context) {
	var signUpRequest dto.SignUpRequest
	if err := c.ShouldBindJSON(&signUpRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := model.User{
		Email:    signUpRequest.Email,
		Password: signUpRequest.Password,
	}

	createdUser, err := controller.UserService.Create(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.GenerateJWT(createdUser.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}
