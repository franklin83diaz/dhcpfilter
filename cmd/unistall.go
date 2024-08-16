package cmd

import (
	"dhcpfilter/pkg"
	"fmt"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall the DHCP Filter service",
	Long:  `Uninstall the DHCP Filter service`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Uninstall()
		fmt.Println("Service uninstalled successfully")
	},
}
