package pkg

import (
	"log"
	"os"
	"strings"

	"github.com/coreos/go-iptables/iptables"
)

// Delete a MAC address rule from iptables
func DelIpt(mac string) {

	ipt, err := iptables.New()
	if err != nil {
		log.Fatalf("Error creating instance ipt: %v", err)
	}

	err = ipt.Delete("filter", "dhcpfilter", "-p", "udp", "-m", "mac", "--mac-source", mac, "-j", "ACCEPT")
	if err != nil {
		log.Fatalf("Error eliminando regla de iptables: %v", err)
	}

}

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

	// TODO: Check incert rule in ipt
	//
	//Move the rule to drop all to the end of the chain
	////////////////////////////////////////////////////////////////////

	//Delete the rule to drop all the packets that don't match the allow list of MAC to port 67

}

func DropAll() {

	ipt, err := iptables.New()
	if err != nil {
		log.Fatalf("Error creando instancia de iptables: %v", err)
	}
	err = ipt.AppendUnique("filter", "INPUT", "-p", "udp", "--dport", "67", "-j", "DROP")
	if err != nil {
		log.Fatalf("Error añadiendo regla de iptables: %v", err)
	}
}
