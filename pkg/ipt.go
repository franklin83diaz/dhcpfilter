package pkg

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/coreos/go-iptables/iptables"
)

// Delete a MAC address rule from ipt
func DelIpt(mac string) {

	ipt, err := iptables.New()
	if err != nil {
		log.Fatalf("Error creating instance ipt: %v", err)
	}

	err = ipt.Delete("filter", "dhcpfilter", "-p", "udp", "-m", "mac", "--mac-source", mac, "-j", "ACCEPT")
	if err != nil {

		if strings.Contains(err.Error(), "exit status 4") {
			log.Printf("Permission denied, try running as root")
		}
		if os.Getenv("DEBUG") == "true" {
			log.Fatalf("Error creating : %v", err)
		}
	}

}

// Add a MAC address rule to ipt
func AddIpt(mac string) {

	ipt, err := iptables.New()
	if err != nil {
		log.Fatalf("Error creating instance ipt: %v", err)
	}

	//Append a rule to the end of a chain in the 'filter' table. dest port udp 67
	err = ipt.InsertUnique("filter", "dhcpfilter", 1, "-p", "udp", "-m", "mac", "--mac-source", mac, "-j", "ACCEPT")
	if err != nil {
		//check status error contains "exit status 4"
		if strings.Contains(err.Error(), "exit status 4") {
			log.Printf("Permission denied, try running as root")
		}

		if os.Getenv("DEBUG") == "true" {
			log.Fatalf("Error Add rule: %v", err)
		}
	}

}

// DropAll drops all packets to port 67
func DropAll() {

	ipt, err := iptables.New()
	if err != nil {
		log.Fatalf("Error creando instancia de iptables: %v", err)
	}
	err = ipt.AppendUnique("filter", "INPUT", "-p", "udp", "--dport", "67", "-j", "DROP")
	if err != nil {
		if strings.Contains(err.Error(), "exit status 4") {
			log.Printf("Permission denied, try running as root")
		}

		if os.Getenv("DEBUG") == "true" {
			log.Fatalf("Error add rule: %v", err)
		}
	}
}

// check if the rule exists
func CheckRuleExists() {

	ipt, err := iptables.New()
	if err != nil {
		log.Fatalf("Error creating instance ipt: %v", err)
	}

	rules, err := ipt.List("filter", "dhcpfilter")
	if err != nil {
		if strings.Contains(err.Error(), "exit status 4") {
			log.Printf("Permission denied, try running as root")
		}
		if os.Getenv("DEBUG") == "true" {
			log.Fatalf("Error creating : %v", err)
		}
	}

	if len(rules) != 0 {
		//red color
		fmt.Println("\033[31m")
		fmt.Println("Alert: Rule dhcpfilter exists, if you running the service as a command line, when you kill this service, the rule will also be deleted and service dhcpfilter will not work, so you must restart the service after killing it.")
		fmt.Println("If you no running the service as a command line, you can ignore this message.")
		fmt.Println("\033[0m")

	}
}
