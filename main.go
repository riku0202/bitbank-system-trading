package main

import (
	"fmt"

	"github.com/riku0202/bitbank-system-trading/bitbank"
	"github.com/riku0202/bitbank-system-trading/config"
)

func main() {
	apiClient := bitbank.New(config.Config.ApiKey, config.Config.ApiSecret)
	result, _ := apiClient.GetTicker("btc_jpy")
	fmt.Printf("(%%+v) %+v\n", result)

}
