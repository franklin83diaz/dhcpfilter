package cmd

import (
	"dhcpfilter/pkg"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all MAC address to the allow list",
	Long: `List a MAC address to the allow list
	
	Example:
	dhcpfilter list
	`,
	Run: func(cmd *cobra.Command, args []string) {

		list, err := pkg.ListMAC()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("MAC address list:")
		for _, mac := range list {
			fmt.Println("  ", mac)
		}

	},
}
