package user

import (
	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUserController interface {
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	AddAdmin(ctx *gin.Context)
}

func NewUserController(service IUserService) IUserController {

	return &userController{
		service: service,
	}
}

type userController struct {
	service IUserService
}

func (c *userController) GetUser(ctx *gin.Context) {
	//email := ctx.Param("email")
	//res, err := c.service.Get(email)

	var request struct {
		Email string `json:"email"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	res, err := c.service.Get(request.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user info:": res})

}

func (c *userController) DeleteUser(ctx *gin.Context) {
	//email := ctx.Param("email")
	//res, err := c.service.Delete(email)

	var request struct {
		Email string `json:"email"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	res, err := c.service.Delete(request.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"response": res, "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "nice delete user")

}

func (c *userController) CreateUser(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.service.Create(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"response": res, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "nice create user")
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.service.Update(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"response": res, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "nice user update")
}

func (c *userController) AddAdmin(ctx *gin.Context) {
	//email := ctx.Param("email")
	//res, err := c.service.AddAdmin(email)

	var request struct {
		Email string `json:"email"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	res, err := c.service.AddAdmin(request.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"response": res, "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "nice add admin")
}