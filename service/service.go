package service

import (
	"dhcpfilter/pkg"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/coreos/go-iptables/iptables"
)

// Service run the rules each time the service is started
func ServiceRun() {
	// create a new instance of iptables
	ipt, err := iptables.New()
	if err != nil {
		log.Fatalf("Error error creating instace ipt: %v", err)
	}

	ListMAC, err := pkg.ListMAC()
	if err != nil {
		log.Fatalf("Error Get the list of Mac: %v", err)
	}

	// Drop all the packets that don't match the allow list of MAC to port 67
	pkg.DropAll()

	for _, mac := range ListMAC {
		// Append a rule to the end of a chain in the 'filter' table. dest port udp 67
		err = ipt.AppendUnique("filter", "INPUT", "-p", "udp", "--dport", "67", "-m", "mac", "--mac-source", mac, "-j", "ACCEPT")
		if err != nil {
			log.Fatalf("Error add rule ipt: %v", err)
		}
	}

	fmt.Println("Service dhcpfilter started")

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Notificar al canal cuando se recibe SIGINT o SIGTERM (ctrl + c)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// wait for the signal
		sig := <-sigs

		if sig == syscall.SIGINT {
			fmt.Println("dhcpfilter Killed")
		}

		//Delete all rules from iptables
		ipt, err := iptables.New()
		if err != nil {
			log.Fatalf("Error error creating instace ipt: %v", err)
		}

		err = ipt.Delete("filter", "INPUT", "-p", "udp", "--dport", "67", "-j", "DROP")
		if err != nil {
			log.Fatalf("Error deleting rule ipt: %v", err)
		}
		done <- true
	}()

	// keep the service running
	<-done

}
