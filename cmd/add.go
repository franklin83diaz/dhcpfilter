package cmd

import (
	"dhcpfilter/pkg"
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a MAC address to the allow list",
	Long: `Add a MAC address to the allow list
	
	Example:
	dhcpfilter add 00:11:22:33:44:55
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid number of arguments")
			return
		}

		mac := args[0]
		err := pkg.AddMAC(mac)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("MAC address added successfully")
	},
}
