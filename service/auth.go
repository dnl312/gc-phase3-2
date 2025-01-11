package service

import (
	"fmt"
	"gc-phase3-2/config"
	"gc-phase3-2/model"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUser(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	var user model.User
	if err := config.DB.Table("users_g2p3w2").Where("username = ?", req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "user not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error", "detail": err.Error()})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid password"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp" : jwt.TimeFunc().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error", "detail": err.Error()})
	}

	return c.JSON(http.StatusOK, model.LoginResponse{Token: tokenString})
}

func validateRegisterUser(user model.RegisterUser) error {
    validate := validator.New()
    err := validate.Struct(user)
    if err != nil {
        for _, e := range err.(validator.ValidationErrors) {
            switch e.Field() {
            default:
                return fmt.Errorf("%s is %s", e.Field(), e.Tag())
            }
        }
    }
    return nil
}

func RegisterUser(c echo.Context) error {
	var user model.RegisterUser
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	err := validateRegisterUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error", "detail": err.Error()})
	}	

	newUser := model.User{
		ID:       uuid.New().String(),
		Username: user.Username,
		Password: string(hashedPassword),
	}

	if err := config.DB.Table("users_g2p3w2").Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error", "detail": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success register",
		"user": map[string]string{
			"Name":  newUser.Username,
		},
	})
}