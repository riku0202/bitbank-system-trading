package main

import (
	"fmt"

	"github.com/riku0202/bitbank-system-trading/bitbank"
	"github.com/riku0202/bitbank-system-trading/config"
)

func main() {
	apiClient := bitbank.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetAssets())

}
