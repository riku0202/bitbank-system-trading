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

const result := {
	"success": 1,
	"data": {
		"assets": [
			{
				"asset": "jpy",
				"amount_precision": 4,
				"onhand_amount": "791.9486",
				"locked_amount": "0.0000",
				"free_amount": "791.9486",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": {
					"threshold": "30000.0000",
					"under": "550.0000",
					"over": "770.0000"
				}
			},
			{
				"asset": "btc",
				"amount_precision": 8,
				"onhand_amount": "0.03820000",
				"locked_amount": "0.00000000",
				"free_amount": "0.03820000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "0.00060000"
			},
			{
				"asset": "ltc",
				"amount_precision": 8,
				"onhand_amount": "0.00000000",
				"locked_amount": "0.00000000",
				"free_amount": "0.00000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "0.00100000"
			},
			{
				"asset": "xrp",
				"amount_precision": 6,
				"onhand_amount": "0.000000",
				"locked_amount": "0.000000",
				"free_amount": "0.000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "0.150000"
			},
			{
				"asset": "eth",
				"amount_precision": 8,
				"onhand_amount": "0.00000000",
				"locked_amount": "0.00000000",
				"free_amount": "0.00000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "0.00500000"
			},
			{
				"asset": "mona",
				"amount_precision": 8,
				"onhand_amount": "0.00000000",
				"locked_amount": "0.00000000",
				"free_amount": "0.00000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "0.00100000"
			},
			{
				"asset": "bcc",
				"amount_precision": 8,
				"onhand_amount": "0.00000000",
				"locked_amount": "0.00000000",
				"free_amount": "0.00000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "0.00100000"
			},
			{
				"asset": "xlm",
				"amount_precision": 7,
				"onhand_amount": "0.0000000",
				"locked_amount": "0.0000000",
				"free_amount": "0.0000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "0.0100000"
			},
			{
				"asset": "qtum",
				"amount_precision": 8,
				"onhand_amount": "0.00000000",
				"locked_amount": "0.00000000",
				"free_amount": "0.00000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "0.01000000"
			},
			{
				"asset": "bat",
				"amount_precision": 8,
				"onhand_amount": "0.00000000",
				"locked_amount": "0.00000000",
				"free_amount": "0.00000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "45.00000000"
			},
			{
				"asset": "omg",
				"amount_precision": 8,
				"onhand_amount": "0.00000000",
				"locked_amount": "0.00000000",
				"free_amount": "0.00000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "5.00000000"
			},
			{
				"asset": "xym",
				"amount_precision": 6,
				"onhand_amount": "0.000000",
				"locked_amount": "0.000000",
				"free_amount": "0.000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "2.000000"
			},
			{
				"asset": "link",
				"amount_precision": 8,
				"onhand_amount": "0.00000000",
				"locked_amount": "0.00000000",
				"free_amount": "0.00000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "1.10000000"
			},
			{
				"asset": "mkr",
				"amount_precision": 8,
				"onhand_amount": "0.00000000",
				"locked_amount": "0.00000000",
				"free_amount": "0.00000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "0.02000000"
			},
			{
				"asset": "boba",
				"amount_precision": 8,
				"onhand_amount": "0.00000000",
				"locked_amount": "0.00000000",
				"free_amount": "0.00000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "17.00000000"
			},
			{
				"asset": "enj",
				"amount_precision": 8,
				"onhand_amount": "0.00000000",
				"locked_amount": "0.00000000",
				"free_amount": "0.00000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "14.00000000"
			},
			{
				"asset": "matic",
				"amount_precision": 8,
				"onhand_amount": "0.00000000",
				"locked_amount": "0.00000000",
				"free_amount": "0.00000000",
				"stop_deposit": false,
				"stop_withdrawal": false,
				"withdrawal_fee": "19.00000000"
			}
		]
	}
}