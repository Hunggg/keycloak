package main

import (
	"fmt"
	"go-key/cmd"
	"go-key/config"
)

func main() {
	cf, err := config.Get_Config()
	if err != nil {
		return
	}

	fmt.Println(cf.ECS.KeyCloak.Family)

	cmd.Execute()
}