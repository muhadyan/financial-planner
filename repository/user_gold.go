package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	db "github.com/muhadyan/financial-planner/database"
	"github.com/muhadyan/financial-planner/model"
)

type UserGoldRepository interface {
	InsertUserGold(userGold *model.UserGold) (*model.UserGold, error)
	GetUserGolds(params *model.UserGoldViewParams) ([]model.UserGold, error)
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

func (c *UserGoldRepositoryCtx) GetUserGolds(params *model.UserGoldViewParams) ([]model.UserGold, error) {
	db := db.DbManager()
	userGolds := []model.UserGold{}

	if params.UserID != 0 {
		db = db.Where("user_id = ?", params.UserID)
	}

	if params.NullSellPrice {
		db = db.Where("sell_price IS NULL")
	}

	db = db.Order(`buy_date ASC`)

	params.BasedFilter = params.BasedFilter.DefaultQuery()
	db = db.Offset(params.BasedFilter.Offset).Limit(params.BasedFilter.Limit)

	err := db.Find(&userGolds).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return userGolds, nil
}
