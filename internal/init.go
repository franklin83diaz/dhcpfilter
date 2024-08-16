package internal

import (
	"dhcpfilter/pkg"
	"log"
	"os"
)

func Init() {
	pkg.SetPathFile("/var/dhcpfilter/mac_allow_list")

	if !CheckFileMac() {
		CreateTempFileMac()
	}

	if os.Getenv("DEBUG") == "true" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	} else {
		log.SetFlags(log.LstdFlags)
	}

}
