package main

import (
	"fmt"

	"github.com/halflifeviper/keycloak/cmd"
	"github.com/halflifeviper/keycloak/config"
)

func main() {
	cf, err := config.Get_Config()
	if err != nil {
		return
	}

	fmt.Println(cf.ECS.KeyCloak.Family)

	cmd.Execute()
}
