package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	db "github.com/muhadyan/financial-planner/database"
	"github.com/muhadyan/financial-planner/model"
)

type UserRoleRepository interface {
	GetUserRoles(params *model.UserRole) ([]model.UserRole, error)
	InsertUserRole(userRole *model.UserRole) (*model.UserRole, error)
}

type UserRoleRepositoryCtx struct{}

func (c *UserRoleRepositoryCtx) GetUserRoles(params *model.UserRole) ([]model.UserRole, error) {
	db := db.DbManager()
	userRoles := []model.UserRole{}

	if params.UserID != 0 {
		db = db.Where("user_id = ?", params.UserID)
	}

	err := db.Find(&userRoles).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return userRoles, nil
}

func (c *UserRoleRepositoryCtx) InsertUserRole(userRole *model.UserRole) (*model.UserRole, error) {
	db := db.DbManager()
	err := db.Create(userRole).Error
	if err != nil {
		return nil, err
	}

	return userRole, nil
}
