package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var users []User
var idCounter = 1

func main() {
	r := gin.Default()

	r.POST("/register", func(ctx *gin.Context) {
		var register User
		if err := ctx.ShouldBindBodyWithJSON(&register); err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, Response{
				Message: err.Error(),
				Data:    nil,
				Status:  false,
				Error:   "Internal Server Error",
			})
			return
		}

		if register.Email == "" || register.Password == "" {
			ctx.JSON(http.StatusBadRequest, Response{
				Message: "Email and Pasword are required",
				Data:    nil,
				Status:  false,
				Error:   "Bad Request",
			})
			return
		}
		if !formatEmail(register.Email) {
			ctx.JSON(http.StatusBadRequest, Response{
				Message: "Format Email salah",
				Data:    nil,
				Status:  false,
				Error:   "Bad Request",
			})
			return
		}

		newUser := User{
			Id:       idCounter,
			Email:    register.Email,
			Password: register.Password,
		}

		users = append(users, newUser)
		idCounter++

		log.Println("User sudah terdaftar: ", users)
		ctx.JSON(http.StatusCreated, Response{
			Message: "Register Berhasil!",
			Data:    newUser,
			Status:  true,
			Error:   "",
		})
	})
	r.POST("/login", func(ctx *gin.Context) {
		var login User
		if err := ctx.ShouldBindBodyWithJSON(&login); err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, Response{
				Message: err.Error(),
				Data:    nil,
				Status:  false,
				Error:   "Internal Server Error",
			})
			return
		}
		if login.Email == "" || login.Password == "" {
			ctx.JSON(http.StatusBadRequest, Response{
				Message: "Email and Pasword are required",
				Data:    nil,
				Status:  false,
				Error:   "Bad Request",
			})
			return
		}
		var checkUser User
		var isCheck bool = false
		for _, usr := range users {
			if usr.Email == login.Email && usr.Password == login.Password {
				checkUser = usr
				isCheck = true
				break
			}
		}
		if !isCheck {
			if login.Email == "" || login.Password == "" {
				ctx.JSON(http.StatusUnauthorized, Response{
					Message: "Email atau Password salah",
					Data:    nil,
					Status:  false,
					Error:   "Unauthorized",
				})
				return
			}
		}
		ctx.JSON(http.StatusOK, Response{
			Message: "Login Successfully",
			Data:    checkUser,
			Status:  true,
			Error:   "",
		})
	})

	r.Run("localhost:9000")

}

func formatEmail(Email string) bool {
	pattern := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(Email)
}

type Response struct {
	Message string
	Data    any
	Status  bool
	Error   string
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type Users struct {
// 	Users map[string]User
// }
