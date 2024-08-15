package cmd

import (
	"dhcpfilter/internal"
	"dhcpfilter/service"
	"fmt"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "is the service that filters DHCP requests",
	Long: `systemctl start dhcpfilter 
	`,
	Run: func(cmd *cobra.Command, args []string) {

		internal.Init()
		service.ServiceRun()

		fmt.Println("Service dhcpfilter started")
	},
}
