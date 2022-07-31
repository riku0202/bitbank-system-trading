package main

import (
	"fmt"

	"github.com/riku0202/bitbank-system-trading/config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
