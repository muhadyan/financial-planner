package repository

import (
	db "github.com/muhadyan/financial-planner/database"
	"github.com/muhadyan/financial-planner/model"
)

type UserGoldRepository interface {
	InsertUserGold(userGold *model.UserGold) (*model.UserGold, error)
}

type UserGoldRepositoryCtx struct{}

func (c *UserGoldRepositoryCtx) InsertUserGold(userGold *model.UserGold) (*model.UserGold, error) {
	db := db.DbManager()
	err := db.Create(userGold).Error
	if err != nil {
		return nil, err
	}

	return userGold, nil
}
