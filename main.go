package main

import (
	"fmt"
	"your_project/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println("Config yüklendi:", cfg.AppPort)
}
