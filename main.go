package main

import (
	"fmt"

	"github.com/riku0202/bitbank-system-trading/bitbank"
	"github.com/riku0202/bitbank-system-trading/config"
)

func main() {
	apiClient := bitbank.New(config.Config.ApiKey, config.Config.ApiSecret)
	result, err := apiClient.GetOrder()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("(%%+v) %+v\n", result)

}
