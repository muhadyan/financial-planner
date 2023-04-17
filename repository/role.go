package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	db "github.com/muhadyan/financial-planner/database"
	"github.com/muhadyan/financial-planner/model"
)

type RoleRepository interface {
	GetRolesByIDs(ids []int) ([]model.Role, error)
	GetRole(params *model.Role) (*model.Role, error)
	InsertRole(role *model.Role) (*model.Role, error)
}

type RoleRepositoryCtx struct{}

func (c *RoleRepositoryCtx) GetRolesByIDs(ids []int) ([]model.Role, error) {
	db := db.DbManager()
	roles := []model.Role{}

	db = db.Where("id IN (?)", ids)

	err := db.Find(&roles).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return roles, nil
}

func (c *RoleRepositoryCtx) GetRole(params *model.Role) (*model.Role, error) {
	db := db.DbManager()
	role := model.Role{}

	if params.Name != "" {
		db = db.Where("name = ?", params.Name)
	}

	err := db.First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &role, nil
}

func (c *RoleRepositoryCtx) InsertRole(role *model.Role) (*model.Role, error) {
	db := db.DbManager()
	err := db.Create(role).Error
	if err != nil {
		return nil, err
	}

	return role, nil
}
