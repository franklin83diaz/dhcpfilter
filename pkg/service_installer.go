package pkg

import (
	"log"
	"os"
	"os/exec"
)

// This function Install the service in the system linux systemctl

func Install() {

	//Change iptables iptables-nft to iptables-legacy
	_ = exec.Command("sudo", "update-alternatives", "--set", "iptables", "/usr/sbin/iptables-legacy")

	//Create the directory and file /var/dhcpfilter/mac_allow_list if not exists
	if _, err := os.Stat("/var/dhcpfilter"); os.IsNotExist(err) {
		os.Mkdir("/var/dhcpfilter", 0755)
	}

	if _, err := os.Stat("/var/dhcpfilter/mac_allow_list"); os.IsNotExist(err) {
		os.Create("/var/dhcpfilter/mac_allow_list")
		log.Println("File /var/dhcpfilter/mac_allow_list created")
	}

	//Copy sample binary to /usr/local/bin
	CopyFile("dhcpfilter", "/usr/local/bin/dhcpfilter")
	os.Chmod("/usr/local/bin/dhcpfilter", 0755)

	serviceContent := `[Unit]
	Description= DHCP Filter
	After=network.target
	
	[Service]
	ExecStart=/usr/local/bin/dhcpfilter service
	Restart=always
	User=root
	Group=root
	
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
