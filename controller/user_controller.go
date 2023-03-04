package controller

import (
	"goapi/controller/dto"
	"goapi/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{UserService: userService}
}

func (ctr *userController) GetUser(c *gin.Context) {
	email := c.MustGet("email").(string)
	user, err := ctr.UserService.FindByEmail(email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (ctr *userController) UpdateUser(c *gin.Context) {
	email := c.MustGet("email").(string)
	user, err := ctr.UserService.FindByEmail(email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var updateUserRequest dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.Name = updateUserRequest.Name
	user, err = ctr.UserService.Update(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (controller *userController) DeleteUser(c *gin.Context) {
	email := c.MustGet("email").(string)
	err := controller.UserService.Delete(email)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user deleted"})
}
