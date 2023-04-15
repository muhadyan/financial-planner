package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	db "github.com/muhadyan/financial-planner/database"
	"github.com/muhadyan/financial-planner/model"
)

type UserRepository interface {
	GetUser(params *model.User) (*model.User, error)
	InsertUser(params *model.User) (*model.User, error)
}

type UserRepositoryCtx struct{}

func (c *UserRepositoryCtx) GetUser(params *model.User) (*model.User, error) {
	db := db.DbManager()
	users := model.User{}

	if params.Username != "" {
		db = db.Where("username = ?", params.Username)
	}

	if params.Email != "" {
		db = db.Where("email = ?", params.Email)
	}

	err := db.First(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &users, nil
}

func (c *UserRepositoryCtx) InsertUser(user *model.User) (*model.User, error) {
	db := db.DbManager()
	err := db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
