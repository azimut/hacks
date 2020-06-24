package main

import (
	"fmt"

	"github.com/tomsteele/go-nmap"
)

// hostLine returns a bareminimun line foran nmap host
func hostLine(host nmap.Host, start string) string {
	address := " "
	hostname := " "
	if len(host.Addresses) > 0 {
		address = host.Addresses[0].Addr
	}
	if len(host.Hostnames) > 0 {
		hostname = host.Hostnames[0].Name
	}
	return fmt.Sprintf("%s - %s - %s - %s",
		start,
		host.Status.State,
		address,
		hostname)
}

// printNmapFile ...
func printNmap(parsed *nmap.NmapRun) error {
	for _, host := range parsed.Hosts {
		modeline := hostLine(host, parsed.StartStr)
		if len(host.Ports) > 0 {
			for _, port := range host.Ports {
				fmt.Printf("%s - %s - %d - %s - %s\n",
					modeline,
					port.State.State,
					port.PortId,
					port.Service.Name,
					port.Service.Product)
			}
		} else {
			fmt.Println(modeline)
		}
	}
	return nil
}
