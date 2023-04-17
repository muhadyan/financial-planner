package repository

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	db "github.com/muhadyan/financial-planner/database"
	"github.com/muhadyan/financial-planner/model"
)

type UserRepository interface {
	GetUser(params *model.User) (*model.User, error)
	InsertUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
}

type UserRepositoryCtx struct{}

func (c *UserRepositoryCtx) GetUser(params *model.User) (*model.User, error) {
	db := db.DbManager()
	user := model.User{}

	if params.ID != 0 {
		db = db.Where("id = ?", params.ID)
	}

	if params.Username != "" {
		db = db.Where("username = ?", params.Username)
	}

	if params.Email != "" {
		db = db.Where("email = ?", params.Email)
	}

	err := db.First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (c *UserRepositoryCtx) InsertUser(user *model.User) (*model.User, error) {
	db := db.DbManager()
	err := db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *UserRepositoryCtx) UpdateUser(user *model.User) (*model.User, error) {
	db := db.DbManager().Model(&model.User{})
	update := map[string]interface{}{}

	if user.IsActive {
		update["is_active"] = true
	}

	if !user.IsActive {
		update["is_active"] = false
	}

	if user.Token != "" {
		update["token"] = user.Token
	}

	update["updated_at"] = time.Now()

	err := db.Where("id = ?", user.ID).Updates(update).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
