package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"go-key/config"
	"go-key/keycloak"
)

// destroyInstanceCmd represents the destroyInstance command
var destroyInstanceCmd = &cobra.Command{
	Use:   "destroyInstance",
	Short: "destroy instance command",
	Long: `Destroy instance command lint`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("destroyInstance called")
		cf, err := config.Get_Config()
		if err != nil {
			return
		}

		if err := keycloak.DestroyKeyCloak(cf); err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(destroyInstanceCmd)
}
