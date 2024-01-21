/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/halflifeviper/keycloak/config"
	keycloak "github.com/halflifeviper/keycloak/keycloakservice"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "root command",
	Long:  `Root command line`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
		cf, err := config.Get_Config()
		if err != nil {
			return
		}
		err = keycloak.CreateKeyCloak(cf)
		fmt.Println(err)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
