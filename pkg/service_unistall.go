package pkg

import (
	"log"
	"os"
	"os/exec"
)

// This function Uninstall the service in the system linux systemctl
func Uninstall() {
	// Stop the service
	cmd := exec.Command("sudo", "systemctl", "stop", "dhcpfilter")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(output))
		log.Fatalf("Error stopping service: %v", err)
	}

	// Disable the service
	cmd = exec.Command("sudo", "systemctl", "disable", "dhcpfilter")
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Println(string(output))
		log.Fatalf("Error disabling service: %v", err)
	}

	// Remove the service unit file
	serviceFilePath := "/etc/systemd/system/dhcpfilter.service"
	err = os.Remove(serviceFilePath)
	if err != nil {
		log.Fatalf("Error removing service file: %v", err)
	}

	// Remove the binary
	err = os.Remove("/usr/local/bin/dhcpfilter")
	if err != nil {
		log.Fatalf("Error removing binary: %v", err)
	}

	// Remove the directory and file /var/dhcpfilter/mac_allow_list
	err = os.RemoveAll("/var/dhcpfilter")
	if err != nil {
		log.Fatalf("Error removing /var/dhcpfilter: %v", err)
	}

	log.Println("Service uninstalled successfully")
}
