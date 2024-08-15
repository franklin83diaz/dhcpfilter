package pkg

import (
	"log"

	"github.com/coreos/go-iptables/iptables"
)

// Delete a MAC address rule from iptables
func DelIpt(mac string) {

	ipt, err := iptables.New()
	if err != nil {
		log.Fatalf("Error creando instancia de iptables: %v", err)
	}

	err = ipt.Delete("filter", "INPUT", "-p", "udp", "-m", "mac", "--mac-source", mac, "-j", "ACCEPT")
	if err != nil {
		log.Fatalf("Error eliminando regla de iptables: %v", err)
	}

}

func AddIpt(mac string) {

	ipt, err := iptables.New()
	if err != nil {
		log.Fatalf("Error creando instancia de iptables: %v", err)
	}

	//Append a rule to the end of a chain in the 'filter' table. dest port udp 67
	err = ipt.AppendUnique("filter", "INPUT", "-p", "udp", "-m", "mac", "--mac-source", mac, "-j", "ACCEPT")
	if err != nil {
		log.Fatalf("Error añadiendo regla de iptables: %v", err)
	}

	//Move the rule to drop all to the end of the chain
	////////////////////////////////////////////////////////////////////

	//Delete the rule to drop all the packets that don't match the allow list of MAC to port 67
	err = ipt.Delete("filter", "INPUT", "-p", "udp", "-j", "--dport", "67", "DROP")
	if err != nil {
		log.Fatalf("Error eliminando regla de iptables: %v", err)
	}

	//Create a rule to drop all the packets that don't match the allow list of MAC to port 67
	err = ipt.AppendUnique("filter", "INPUT", "-p", "udp", "--dport", "67", "-j", "DROP")
	if err != nil {
		log.Fatalf("Error añadiendo regla de iptables: %v", err)
	}

	////////////////////////////////////////////////////////////////////

}
