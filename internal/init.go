package internal

import "dhcpfilter/pkg"

func Init() {
	pkg.SetPathFile("/var/dhcpfilter/mac_allow_list")

	if !CheckFileMac() {
		CreateTempFileMac()
	}

}
