package main

import (
	"fmt"

	config "github.com/jalexakos/gator-config/internal/config"
)

func main() {
	configFile, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}
	configFile.SetUser("josh")
	fmt.Println("Config current user:", configFile.CurrentUserName)
	fmt.Println("Config DB URL:", configFile.DbUrl)
}
