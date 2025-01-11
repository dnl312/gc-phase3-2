package repo

import (
	"server/models"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserRepository struct{
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{DB: db}
}

func (c UserRepository) RegisterUser(data *models.User) error {
	result := c.DB.Create(data)
	if err := result.Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return status.Error(codes.AlreadyExists, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}



