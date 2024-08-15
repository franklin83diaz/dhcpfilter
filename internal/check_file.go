package internal

import (
	"dhcpfilter/pkg"
	"fmt"
	"os"
)

// CheckFileMac checks if the file exists
func CheckFileMac() bool {

	pathFile := pkg.GetPathFile()

	// Check if the file exists
	if _, err := os.Stat(pathFile); os.IsNotExist(err) {
		fmt.Println("File ", pathFile, " does not exist")
		return false
	}

	return true

}

// CreateTempFileMac creates a temporary file for testing purposes
func CreateTempFileMac() {
	pathFile := "mac_allow_list"
	// Create the file but no truncate if it already exists
	_, err := os.OpenFile(pathFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}

	pkg.SetPathFile(pathFile)
	// color yellow
	fmt.Println("\033[33m")
	fmt.Println("warning: You are using a temporary file, please reinstall the DHCP Filter or Create the file /var/dhcpfilter/mac_allow_list")
	fmt.Println("\033[0m")
}
