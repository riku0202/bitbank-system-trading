package bitbank

import (
	"encoding/json"
	"log"
)

type Assets struct {
	Success int `json:"success"`
	Data    struct {
		Assets []struct {
			Asset           string `json:"asset"`
			FreeAmount      string `json:"free_amount"`
			AmountPrecision int    `json:"amount_precision"`
			OnhandAmount    string `json:"onhand_amount"`
			LockedAmount    string `json:"locked_amount"`
			WithdrawalFee   string `json:"withdrawal_fee"`
			StopDeposit     bool   `json:"stop_deposit"`
			StopWithdrawal  bool   `json:"stop_withdrawal"`
		} `json:"assets"`
	} `json:"data"`
}

func (api *APIClient) GetAssets() ([]Assets, error) {
	url := "/v1/user/assets"
	resp, err := api.doRequest("GET", url, map[string]string{}, nil)
	if err != nil {
		return nil, err
	}

	var assets []Assets
	log.Println(string(resp))
	err = json.Unmarshal(resp, &assets)
	if err != nil {
		log.Printf("Fail to unmarshal: %v", err)
		return nil, err
	}

	return assets, nil
}