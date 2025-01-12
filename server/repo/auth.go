package repo

import (
	"fmt"
	"log"
	"os"
	"server/config"
	"server/model"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		Db: db,
	}
}

func (u UserRepository) LoginUser(user model.User) (string, error) {
	tokenString := ""

	var userGet model.User
	if err := config.DB.Table("users_g2p3w2").Where("username = ?", user.Username).First(&userGet).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return tokenString, status.Errorf(codes.NotFound, "user %s not found %s", user.Username, user.Password)
		}
		return tokenString, status.Error(codes.Internal, err.Error())
	}

	errString := fmt.Sprintf("invalid username or password for user %s password %s", user.Username, user.Password)
	if err := bcrypt.CompareHashAndPassword( []byte(userGet.Password), []byte(user.Password)); err != nil {
		return errString, status.Errorf(codes.Unauthenticated, err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userGet.ID,
		"exp" : jwt.TimeFunc().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return tokenString, status.Error(codes.Internal, err.Error())
	}

	return tokenString, nil
}

func validateRegisterUser(user model.RegisterUser) error {
	log.Printf("Validating user: %+v", user)
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

func (u UserRepository) RegisterUser(user model.RegisterUser) error {
	// var user model.RegisterUser
	// if err := c.Bind(&user); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	// }

	err := validateRegisterUser(user)
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}	

	newUser := model.User{
		ID:       uuid.New().String(),
		Username: user.Username,
		Password: string(hashedPassword),
	}

	if err := config.DB.Table("users_g2p3w2").Create(&newUser).Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}