package repository

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	db "github.com/muhadyan/financial-planner/database"
	"github.com/muhadyan/financial-planner/model"
	"github.com/muhadyan/financial-planner/utils"
)

type UserGoldRepository interface {
	InsertUserGold(userGold *model.UserGold) (*model.UserGold, error)
	GetUserGolds(params *model.UserGoldViewParams) ([]model.UserGold, error)
	GetAllUserGolds(params *model.UserGold) ([]model.UserGold, error)
	GetUserGold(params *model.UserGold) (*model.UserGold, error)
	UpdateUserGold(userGold *model.UserGold) (*model.UserGold, error)
	DeleteUserGold(id int) error
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

func (c *UserGoldRepositoryCtx) GetAllUserGolds(params *model.UserGold) ([]model.UserGold, error) {
	db := db.DbManager()
	userGolds := []model.UserGold{}

	if params.UserID != 0 {
		db = db.Where("user_id = ?", params.UserID)
	}

	db = db.Order(`buy_date ASC`)

	err := db.Find(&userGolds).Error
	if err != nil {
		return nil, err
	}

	return userGolds, nil
}

func (c *UserGoldRepositoryCtx) GetUserGold(params *model.UserGold) (*model.UserGold, error) {
	db := db.DbManager()
	userGold := model.UserGold{}

	if params.ID != 0 {
		db = db.Where("id = ?", params.ID)
	}

	err := db.First(&userGold).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &userGold, nil
}

func (c *UserGoldRepositoryCtx) UpdateUserGold(userGold *model.UserGold) (*model.UserGold, error) {
	db := db.DbManager().Model(&model.UserGold{})
	update := map[string]interface{}{}

	if userGold.Weight != 0 {
		update["weight"] = userGold.Weight
	}

	if userGold.BuyPrice != 0 {
		update["buy_price"] = userGold.BuyPrice
	}

	if userGold.SellPrice != nil {
		update["sell_price"] = userGold.SellPrice
	}

	if userGold.BuyDate.String() != utils.EmptyDate {
		update["buy_date"] = userGold.BuyDate
	}

	if userGold.SellDate != nil {
		update["sell_date"] = userGold.SellDate
	}

	update["updated_at"] = time.Now()

	err := db.Where("id = ?", userGold.ID).Updates(update).Error
	if err != nil {
		return nil, err
	}

	return userGold, nil
}

func (c *UserGoldRepositoryCtx) DeleteUserGold(id int) error {
	db := db.DbManager()
	err := db.Delete(&model.UserGold{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
