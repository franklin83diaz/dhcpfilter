package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var pathFile string

// Set the path file
func SetPathFile(path string) {
	pathFile = path
}

// Get the path file
func GetPathFile() string {
	return pathFile
}

// Find a MAC address in the file
//
// Parameters:
// mac: MAC address to find  (Capitalized format)
//
// Return:
// Return true if the MAC address is found
// Return false if the MAC address is not found
func FindMAC(mac string) bool {
	file, err := os.Open(pathFile)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.ToUpper(scanner.Text()) == mac {
			return true
		}
	}

	return false
}

// Add a MAC address to the file
//
// Parameters:
// mac: MAC address to add  (Capitalized format)
//
// Return:
// Return nil if the MAC address is added successfully
// Return error if there is an error adding the MAC address
func AddMAC(mac string) error {

	mac = strings.ToUpper(mac)

	// Check if the MAC address already exists
	if FindMAC(mac) {
		fmt.Println("MAC address already exists")
		return nil
	}

	// Open the file
	file, err := os.OpenFile(pathFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()

	// Write the MAC address to the file
	_, err = file.WriteString(mac + "\n")
	if err != nil {
		fmt.Println(err)
		return err
	}

	//Do in the service
	AddIpt(mac)

	return nil

}

// Remove a MAC address from the file
//
// Parameters:
// mac: MAC address to remove  (Capitalized format)
//
// Return:
// Return nil if the MAC address is removed successfully
// Return error if there is an error removing the MAC address
func RemoveMAC(mac string) error {

	mac = strings.ToUpper(mac)

	// Open the file
	file, err := os.Open(pathFile)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()

	var lines []string

	// Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the MAC address from the file except the MAC address to remove
		if strings.ToUpper(scanner.Text()) != mac { // Use uppercase for fixed case the file written in lowercase manually in lowercase.
			lines = append(lines, scanner.Text())
		}
	}

	// Open the file for writing
	file, err = os.OpenFile(pathFile, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	// Write the file
	for _, line := range lines {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	//Do in the service
	DelIpt(mac)

	return nil
}

// List all MAC addresses in the file
//
// Return:
// Return a list of MAC addresses
// Return error if there is an error reading the file
func ListMAC() (listMAC []string, err error) {
	file, err := os.Open(pathFile)
	if err != nil {
		fmt.Println(err)
		return listMAC, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		listMAC = append(listMAC, scanner.Text())
	}

	return listMAC, nil
}
