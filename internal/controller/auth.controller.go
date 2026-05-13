package controller

import (
	"log"
	"net/http"

	"github.com/aqilknz/koda-b7-backend/internal/dto"
	"github.com/aqilknz/koda-b7-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService *service.UserService
}

func NewAuthController(us *service.UserService) *AuthController {
	return &AuthController{
		UserService: us,
	}
}

func (ac *AuthController) Register(ctx *gin.Context) {
	var register dto.UserRequest
	if err := ctx.ShouldBindBodyWithJSON(&register); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Message: err.Error(), Data: nil, Status: false, Error: "Internal Server Error",
		})
		return
	}

	if register.Email == "" || register.Password == "" {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Message: "Email and Pasword are required", Data: nil, Status: false, Error: "Bad Request",
		})
		return
	}

	if len(register.Password) < 8 {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Message: "Password minimal harus 8 karakter", Data: nil, Status: false, Error: "Bad Request",
		})
		return
	}

	if !ac.UserService.ValidateEmail(register.Email) {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Message: "Format Email salah", Data: nil, Status: false, Error: "Bad Request",
		})
		return
	}

	newUser, _ := ac.UserService.RegisterUser(register)

	ctx.JSON(http.StatusCreated, dto.Response{
		Message: "Register Berhasil!",
		Data:    newUser,
		Status:  true,
		Error:   "",
	})
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var login dto.UserRequest
	if err := ctx.ShouldBindBodyWithJSON(&login); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Message: err.Error(), Data: nil, Status: false, Error: "Internal Server Error",
		})
		return
	}

	if login.Email == "" || login.Password == "" {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Message: "Email and Pasword are required", Data: nil, Status: false, Error: "Bad Request",
		})
		return
	}
	if len(login.Password) < 8 {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Message: "Password minimal harus 8 karakter", Data: nil, Status: false, Error: "Bad Request",
		})
		return
	}

	if !ac.UserService.ValidateEmail(login.Email) {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Message: "Format Email salah", Data: nil, Status: false, Error: "Bad Request",
		})
		return
	}

	user, found := ac.UserService.LoginUser(login)
	if !found {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Message: "Email atau Password salah", Data: nil, Status: false, Error: "Unauthorized",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "Login Successfully",
		Data:    user,
		Status:  true,
		Error:   "",
	})
}
