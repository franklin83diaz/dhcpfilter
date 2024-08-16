package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show DHCP Filter version",
	Long:  `Show DHCP Filter version and exit`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DHCP Filter v0.0.1")
	},
}
