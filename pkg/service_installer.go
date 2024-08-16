package pkg

import (
	"log"
	"os"
	"os/exec"
)

// This function Install the service in the system linux systemctl

func Install() {

	//Copy sample binary to /usr/local/bin
	CopyFile("dhcpfilter", "/usr/local/bin/dhcpfilter")
	os.Chmod("/usr/local/bin/dhcpfilter", 0755)

	serviceContent := `[Unit]
	Description= DHCP Filter
	After=network.target
	
	[Service]
	ExecStart=/usr/local/bin/dhcpfilter service
	Restart=always
	User=nobody
	Group=nogroup
	
	[Install]
	WantedBy=multi-user.target
	`

	// Write the service unit file
	serviceFilePath := "/etc/systemd/system/dhcpfilter.service"
	err := os.WriteFile(serviceFilePath, []byte(serviceContent), 0644)
	if err != nil {
		log.Fatalf("Error writing service file: %v", err)
	}

	// Reload systemd to recognize the new service
	cmd := exec.Command("sudo", "systemctl", "daemon-reload")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(output))
		log.Fatalf("Error reloading systemd: %v", err)
	}

	// Enable the service
	cmd = exec.Command("sudo", "systemctl", "enable", "dhcpfilter")
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Println(string(output))
		log.Fatalf("Error enabling service: %v", err)
	}

}
