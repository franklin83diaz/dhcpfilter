# DHCP FILTER

DHCP Filter is a service that filters the DHCP service and only assigns IP addresses to the MAC addresses that are allowed.




## Authors

- [@Franklin](https://github.com/franklin83diaz)


## Demo

Checking service is running

`
systemctl status dhcpfilter
`

Change default actions

`
dhcpfilter default close
`

Allow mac address
`
dhcpfilter allow 00:00...
` 
List mac address allow
`
dhcpfilter List
`

remove mac from allowed mac list

`
dhcpfilter delete 00:00..
`