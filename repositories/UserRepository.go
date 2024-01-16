package repositories

import (
	"errors"
	"github.com/dev-hyunsang/clone-stackbuck-backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	failedCreateMsg  = "failed to Create User in Users Table"
	failedGetUserMsg = "failed to Get User Data in Users Table"
)

func CreateUser(db *gorm.DB, user *models.Users) (int64, error) {
	tx := db.Table("users").Create(&user)
	if tx.Error != nil {
		return 0, errors.New(failedCreateMsg)
	}

	return tx.RowsAffected, nil
}

func GetuserByEmail(db *gorm.DB, email string) (*models.Users, error) {
	user := new(models.Users)

	tx := db.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return nil, errors.New(failedGetUserMsg)
	}

	return user, nil
}

func GetUserByPk(db *gorm.DB, pk int) (*models.Users, error) {
	user := new(models.Users)

	tx := db.Where("pk = ?", pk).First(&user)
	if tx.Error != nil {
		return nil, errors.New(failedGetUserMsg)
	}

	return user, nil
}

func GetUserByUUID(db *gorm.DB, userUUID uuid.UUID) (*models.Users, error) {
	users := new(models.Users)

	tx := db.Where("uuid = ?", userUUID).First(&users)
	if tx.Error != nil {
		return nil, errors.New(failedGetUserMsg)
	}

	return users, nil
}
