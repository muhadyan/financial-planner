package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/muhadyan/financial-planner/config"
	"github.com/muhadyan/financial-planner/model"
	"github.com/muhadyan/financial-planner/utils"
)

type CurrentGoldRepository interface {
	GetCurrentGold() (*model.CurrentGolds, error)
}

type CurrentGoldRepositoryCtx struct{}

func (c *CurrentGoldRepositoryCtx) GetCurrentGold() (*model.CurrentGolds, error) {
	result := new(model.CurrentGolds)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	requestURL := fmt.Sprintf("%s/prices/hargaemas-org", config.GetConfig().CurrentGoldUrl)
	resp, statusCode, err := utils.GetHTTPRequestJSON("GET", requestURL, nil, headers)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, nil
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
