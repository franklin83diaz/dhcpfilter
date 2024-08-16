package cmd

import (
	"dhcpfilter/pkg"
	"fmt"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the DHCP Filter service",
	Long: `Install the DHCP Filter service
	
	Example:
	dhcpfilter install

	check the service status:
	systemctl status dhcpfilter

	`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Install()
		fmt.Println("Service installed successfully")
	},
}
