package controller

import (
	"adoptme/app/entity"
	"adoptme/app/model"
	"adoptme/config"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	request := entity.RegisterRequest{}
	ctx.Bind(&request)

	password := sha256.Sum256([]byte(request.Password))

	user := model.User{
		Fullname: request.Fullname,
		Email:    request.Email,
		Password: hex.EncodeToString(password[:]),
		Address:  request.Address,
	}

	result := config.DB.Create(&user)

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "gagal register",
			"data":    result.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "sukses register",
		"data":    &request,
	})
}

func LoginUser(ctx *gin.Context) {
	request := entity.LoginRequest{}
	ctx.Bind(&request)

	user := model.User{}

	password := sha256.Sum256([]byte(request.Password))
	passwordString := hex.EncodeToString(password[:])

	result := config.DB.Where("email = ? and password = ?", request.Email, passwordString).First(&user)

	if result.RowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "login sukses",
			"data":    result.RowsAffected,
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"status":  false,
		"message": "login gagal",
		"data":    result.RowsAffected,
	})

}
