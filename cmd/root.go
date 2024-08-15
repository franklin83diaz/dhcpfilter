package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dhcpfilter",
	Short: "run dhcpfilter service",
	Long:  `dhcpfilter is a service that filters DHCP requests.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Run the service
	},
}

func init() {
	//commands
	rootCmd.AddCommand(versionCmd)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
