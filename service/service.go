package service

import (
	"dhcpfilter/pkg"
	"log"

	"github.com/coreos/go-iptables/iptables"
)

// Service run the rules each time the service is started
func ServiceRun() {
	// create a new instance of iptables
	ipt, err := iptables.New()
	if err != nil {
		log.Fatalf("Error creando instancia de iptables: %v", err)
	}

	ListMAC, err := pkg.ListMAC()
	if err != nil {
		log.Fatalf("Error Get the list of Mac: %v", err)
	}

	for _, mac := range ListMAC {
		// Append a rule to the end of a chain in the 'filter' table. dest port udp 67
		err = ipt.AppendUnique("filter", "INPUT", "-p", "udp", "--dport", "67", "-m", "mac", "--mac-source", mac, "-j", "ACCEPT")
		if err != nil {
			log.Fatalf("Error añadiendo regla de iptables: %v", err)
		}
	}

	// Drop all the packets that don't match the allow list of MAC to port 67
	err = ipt.AppendUnique("filter", "INPUT", "-p", "udp", "--dport", "67", "-j", "DROP")
	if err != nil {
		log.Fatalf("Error añadiendo regla de iptables: %v", err)
	}
}
