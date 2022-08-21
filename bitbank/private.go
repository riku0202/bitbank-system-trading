package bitbank

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const privateUrl = "https://api.bitbank.cc"

func (api APIClient) privateHeader(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	message := timestamp + endpoint
	mac := hmac.New(sha256.New, []byte(api.secret))
	mac.Write([]byte(message))
	sign := hex.EncodeToString(mac.Sum(nil))

	return map[string]string{
		"ACCESS-KEY":   api.key,
		"ACCESS-NONCE": timestamp,
		// GETの場合: 「ACCESS-NONCE、リクエストのパス、クエリパラメータ」 を連結させたもの
		// POSTの場合: 「ACCESS-NONCE、リクエストボディのJson文字列」 を連結させたもの
		"ACCESS-SIGNATURE": sign,
		"CONTENT-TYPE":     "application/json",
	}
}

func (api *APIClient) doPrivateRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	baseUrl, err := url.Parse(privateUrl)
	if err != nil {
		return
	}

	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return
	}

	endpoint := baseUrl.ResolveReference(apiURL).String()

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()
	for key, value := range api.privateHeader(method, req.URL.RequestURI(), data) {
		req.Header.Add(key, value)
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

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

	resp, err := api.doPrivateRequest("GET", url, map[string]string{}, nil)
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

type Order struct {
	Success int `json:"success"`
	Data    struct {
		OrderID         int    `json:"order_id"`
		Pair            string `json:"pair"`
		Side            string `json:"side"`
		Type            string `json:"type"`
		StartAmount     string `json:"start_amount"`
		RemainingAmount string `json:"remaining_amount"`
		ExecutedAmount  string `json:"executed_amount"`
		Price           string `json:"price"`
		PostOnly        bool   `json:"post_only"`
		AveragePrice    string `json:"average_price"`
		OrderedAt       int    `json:"ordered_at"`
		ExpireAt        int    `json:"expire_at"`
		TriggeredAt     int    `json:"triggered_at"`
		TrigerPrice     string `json:"triger_price"`
		Status          string `json:"status"`
	} `json:"data"`
}

func (api *APIClient) GetOrder() (*Order, error) {
	url := "/v1/user/spot/order"

	resp, err := api.doPrivateRequest("GET", url, map[string]string{}, nil)
	if err != nil {
		return nil, err
	}

	var order *Order
	log.Println(string(resp))
	err = json.Unmarshal(resp, &order)
	if err != nil {
		log.Printf("Fail to unmarshal: %v", err)
		return nil, err
	}

	return order, nil
}

type ActiveOrders struct {
	Success int `json:"success"`
	Data    struct {
		Orders []struct {
			OrderID         int    `json:"order_id"`
			Pair            string `json:"pair"`
			Side            string `json:"side"`
			Type            string `json:"type"`
			StartAmount     string `json:"start_amount"`
			RemainingAmount string `json:"remaining_amount"`
			ExecutedAmount  string `json:"executed_amount"`
			Price           string `json:"price"`
			PostOnly        bool   `json:"post_only"`
			AveragePrice    string `json:"average_price"`
			OrderedAt       int    `json:"ordered_at"`
			ExpireAt        int    `json:"expire_at"`
			TriggeredAt     int    `json:"triggered_at"`
			TriggerPrice    string `json:"trigger_price"`
			Status          string `json:"status"`
		} `json:"orders"`
	} `json:"data"`
}

func (api *APIClient) GetActiveOrders() (*ActiveOrders, error) {
	url := "/v1/user/spot/active_orders"

	resp, err := api.doPrivateRequest("GET", url, map[string]string{"pair": "btc_jpy"}, nil)
	if err != nil {
		return nil, err
	}

	var activeOrders *ActiveOrders

	log.Println(string(resp))
	err = json.Unmarshal(resp, &activeOrders)
	if err != nil {
		log.Printf("Fail to unmarshal: %v", err)
		return nil, err
	}

	return activeOrders, nil
}
