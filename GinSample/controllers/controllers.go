package controllers

import (
	"GoSamples/GinSample/domain"
	"GoSamples/GinSample/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var usermanager = services.UserManager{}

type UserController struct {
}

func (t *UserController) FindAll(ctx *gin.Context) {
	users := []domain.User{
		{Id: 1, Name: "test1", Account: "test1"},
		{Id: 1, Name: "test2", Account: "test2"},
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (t *UserController) Create(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Fatal("binding param error")
	}

	if err := usermanager.Create(&user); err != nil {
		ctx.JSON(400, err)
	}

	ctx.JSON(200, "Success")
}
