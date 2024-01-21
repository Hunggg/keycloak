package cmd

import (
	"github.com/halflifeviper/keycloak/keycloakservice"
	"log"

	"github.com/spf13/cobra"

	"github.com/halflifeviper/keycloak/config"
)

// createInstanceCmd represents the createInstance command
var createInstanceCmd = &cobra.Command{
	Use:   "createInstance",
	Short: "create instance command",
	Long:  `Create all instance`,
	Run: func(cmd *cobra.Command, args []string) {
		cf, err := config.Get_Config()
		if err != nil {
			return
		}

		if err := keycloakservice.CreateKeyCloak(cf); err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createInstanceCmd)
}
